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
			ID:        toID(item.Service),
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
		key := engine.ServiceTadig{Service: toServiceID(usage.ID), HomeTadig: usage.HomeTadig, VisitorTadig: usage.VisitorTadig}
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
				if service.Rate != nil {
					chargeModels = append(chargeModels, engine.ChargingModel{Service: toServiceID(service.ID), RatingPlan: toRatingPlan(service)})
				}
			}
			serviceGroups = append(serviceGroups, engine.ServiceGroup{HomeTadigs: serviceGroup.HomeTadigs, VisitorTadigs: serviceGroup.VisitorTadigs, ChargingModels: chargeModels})
		}
		parts = append(parts, engine.ContractPart{ServiceGroups: serviceGroups})
	}
	return engine.Contract{Parts: parts}
}

func toRatingPlan(service Service) *engine.RatingPlan {
	var tiers = make([]engine.Tier, 0)
	var from int64 = 0
	for _, tier := range service.Rate {
		tiers = append(tiers, engine.Tier{FixedPrice: tier.FixedPrice, LinearPrice: tier.LinearPrice, From: from, To: tier.Threshold})
		from = tier.Threshold
	}
	return &engine.RatingPlan{Tiers: tiers}
}

func toRatioPlan(service Service) *engine.RatioPlan {
	return &engine.RatioPlan{Balanced: nil, Unbalanced: nil}
}

func toServiceID(ID string) engine.Service {
	if ID == SMS {
		return engine.SMS
	}
	return engine.MOOC
}

func toID(service engine.Service) string {
	if service == engine.SMS {
		return SMS
	}
	return MOOC
}
