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
func ConvertToEngineContract(contract map[string]DiscountModel) engine.Contract {
	var parts = make([]engine.ContractPart, 0)
	for k, discount := range contract {
		var serviceGroups = make([]engine.ServiceGroup, 0)
		for _, serviceGroup := range discount.ServiceGroups {
			var chargeModels = make([]engine.ChargingModel, 0)
			for _, service := range serviceGroup.Services {
				cm := toEngineChargingModels(service)
				if len(cm) > 0 {
					chargeModels = append(chargeModels, cm...)
				}
			}
			serviceGroups = append(serviceGroups, engine.ServiceGroup{HomeTadigs: serviceGroup.HomeTadigs, VisitorTadigs: serviceGroup.VisitorTadigs, ChargingModels: chargeModels})
		}
		parts = append(parts, engine.ContractPart{Party: k, Condition: toEngineCondition(discount.Condition), ServiceGroups: serviceGroups})
	}
	return engine.Contract{Parts: parts}
}

func toEngineChargingModels(service Service) []engine.ChargingModel {
	var ratingPlan RatingPlan = getRatingPlanForPricing(service)
	chargingModels := make([]engine.ChargingModel, 0)
	if isRate(ratingPlan.Rate) {
		chargingModels = append(chargingModels, engine.ChargingModel{
			Service:              service.Service,
			IncludedInCommitment: service.IncludedInCommitment,
			RatingPlan:           toEngineRatingPlan(ratingPlan.Rate),
		})
	}
	// if isRate(service.AccessPricingRate) {
	// 	chargingModels = append(chargingModels, engine.ChargingModel{
	// 		Service:              service.Name,
	// 		IncludedInCommitment: service.IncludedInCommitment,
	// 		RatingPlan:           toEngineRatingPlan(service.AccessPricingRate),
	// 	})
	// }

	if isRatio(ratingPlan.BalancedRate, ratingPlan.UnbalancedRate) {
		chargingModels = append(chargingModels, engine.ChargingModel{
			Service:              service.Service,
			IncludedInCommitment: service.IncludedInCommitment,
			RatioPlan:            toEngineRatioPlan(ratingPlan.BalancedRate, ratingPlan.UnbalancedRate),
		})
	}
	return chargingModels
}

func toEngineRatingPlan(rate Rate) *engine.RatingPlan {
	return &engine.RatingPlan{Tiers: toEngineTiers(rate.Thresholds)}
}

func toEngineRatioPlan(balancedRate Rate, unbalancedRate Rate) *engine.RatioPlan {
	return &engine.RatioPlan{BalancedRate: toEngineTiers(balancedRate.Thresholds), UnbalancedRate: toEngineTiers(unbalancedRate.Thresholds)}
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

func isRatio(balanced Rate, unbalanced Rate) bool {
	return (len(balanced.Thresholds) > 0 || len(unbalanced.Thresholds) > 0)
}

func toEngineCondition(condition Condition) engine.Condition {
	switch c := condition.SelectedConditionName; c {
	case ContractRevenue:
		return engine.Condition{Type: engine.ContractRevenue, Value: condition.SelectedCondition.CommitmentsValue}
	case DealRevenue:
		return engine.Condition{Type: engine.DiscountRevenue, Value: condition.SelectedCondition.CommitmentsValue}
	default:
		return engine.Condition{Type: engine.Unconditional}
	}
}

func toEngineTiers(tiers []Tier) []engine.Tier {
	var engineTiers = make([]engine.Tier, 0)

	if len(tiers) > 0 {
		var to int64 = 0
		for i := len(tiers) - 1; i >= 0; i-- {
			engineTiers = append(engineTiers, engine.Tier{
				FixedPrice:  tiers[i].FixedPrice,
				LinearPrice: tiers[i].LinearPrice,
				From:        tiers[i].Start,
				To:          to,
			})
			to = tiers[i].Start
		}
	}
	return engineTiers
}
