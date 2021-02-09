package models

import (
	engine "github.com/GSMA-CPAS/BWRP-calculation-service/pkg/engine/charging"
)

// ConvertFromEngineResult convert from the engine type to the api type
func ConvertFromEngineResult(result engine.Result) Result {
	return Result{
		IntermediateResults: ConvertFromEngineIntermediateResults(result.IntermediateResults),
	}
}

// ConvertFromEngineIntermediateResults convert from the engine type to the api type
func ConvertFromEngineIntermediateResults(intermediateResults []engine.IntermediateResult) []IntermediateResult {
	ir := make([]IntermediateResult, len(intermediateResults))
	for i, item := range intermediateResults {
		ir[i] = IntermediateResult{
			DealValue: item.DealValue,
			Service:   item.Service,
		}
		ir[i].HomeTadigs = item.HomeTadigs
		ir[i].VisitorTadigs = item.VisitorTadigs
	}
	return ir
}

// ConvertToEngineAggregatedUsage adds all the usage and creates an aggregatedUsage
// This is a map from a specific service,hometadig,visitortadig combination
func ConvertToEngineAggregatedUsage(usages []UsageData) engine.AggregatedUsage {
	var aggregatedUsage engine.AggregatedUsage = make(engine.AggregatedUsage)
	for _, usage := range usages {
		key := engine.ServiceTadig{Service: usage.Service, HomeTadig: usage.HomeTadig, VisitorTadig: usage.VisitorTadig}
		item, ok := aggregatedUsage[key]
		if !ok {
			aggregatedUsage[key] = engine.Usage{Volume: usage.Volume, Unit: engine.Unit(usage.Unit), Charge: usage.Charge, Tax: usage.Tax}
		} else {
			aggregatedUsage[key] = engine.Usage{Volume: item.Volume + usage.Volume, Charge: item.Charge + usage.Charge, Tax: usage.Tax + usage.Tax}
		}
	}
	return aggregatedUsage
}

// ConvertToEngineContract ...
func ConvertToEngineContract(contract []DiscountModel) engine.Contract {
	var parts = make([]engine.ContractPart, 0)
	for _, discount := range contract {
		var serviceGroups = make([]engine.ServiceGroup, 0)
		for _, serviceGroup := range discount.ServiceGroups {
			var chargeModels = make([]engine.ChargingModel, 0)
			for _, service := range serviceGroup.Services {
				chargeModels = append(chargeModels, *toEngineChargingModel(service))
			}
			serviceGroups = append(serviceGroups, engine.ServiceGroup{HomeTadigs: serviceGroup.HomeTadigs, VisitorTadigs: serviceGroup.VisitorTadigs, ChargingModels: chargeModels})
		}
		parts = append(parts, engine.ContractPart{ServiceGroups: serviceGroups})
	}
	return engine.Contract{Parts: parts}
}

func toEngineChargingModel(service Service) *engine.ChargingModel {
	var ratingPlan RatingPlan = getRatingPlanForPricing(service)
	if isRate(ratingPlan.Rate) {
		return &engine.ChargingModel{Service: service.Service, RatingPlan: toEngineRatingPlan(ratingPlan.Rate)}
	} else if isRatio(ratingPlan.BalancedRate, ratingPlan.UnbalancedRate) {
		return &engine.ChargingModel{Service: service.Service, RatioPlan: toEngineRatioPlan(ratingPlan.BalancedRate.Value, ratingPlan.UnbalancedRate.Value)}
	}
	return nil
}

func toEngineRatingPlan(rate Rate) *engine.RatingPlan {
	var engineTiers = make([]engine.Tier, 0)

	if len(rate.Thresholds) > 0 {
		var to int64 = 0
		for i := len(rate.Thresholds) - 1; i >= 0; i-- {
			engineTiers = append(engineTiers, engine.Tier{
				FixedPrice:  rate.Thresholds[i].FixedPrice,
				LinearPrice: rate.Thresholds[i].LinearPrice,
				From:        rate.Thresholds[i].Start,
				To:          to,
			})
			to = rate.Thresholds[i].Start
		}
	} else {
		engineTiers = append(engineTiers, engine.Tier{FixedPrice: int64(rate.FixedPrice), LinearPrice: int64(rate.LinearPrice), From: 0, To: 0})
	}
	return &engine.RatingPlan{Tiers: engineTiers}
}

func toEngineRatioPlan(balancedRate int64, unbalancedRate int64) *engine.RatioPlan {
	return &engine.RatioPlan{BalancedRate: balancedRate, UnbalancedRate: unbalancedRate}
}

func getRatingPlanForPricing(service Service) RatingPlan {
	if service.UsagePricing != nil {
		return service.UsagePricing.RatingPlan
	} else {
		return service.AccessPricing.RatingPlan
	}
}

func isRate(rate Rate) bool {
	return (len(rate.Thresholds) > 0 || rate.FixedPrice > 0 || rate.LinearPrice > 0)
}

func isRatio(balanced BalancedRate, unbalanced BalancedRate) bool {
	return (balanced.Value > 0 || unbalanced.Value > 0)
}
