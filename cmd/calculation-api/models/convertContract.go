package models

import (
	engine "github.com/GSMA-CPAS/BWRP-calculation-service/pkg/engine/charging"
)

// ConvertToEngineContract converts the discountmodels from the API request to the Calculation Engine Contract model
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
	if isRate(ratingPlan) {
		chargingModels = append(chargingModels, engine.ChargingModel{
			Service:              service.Service,
			IncludedInCommitment: service.IncludedInCommitment,
			RatingPlan:           toEngineRatingPlan(ratingPlan.Rate, (ratingPlan.Kind == RatingPlanKindBackToFirst)),
		})
	}
	if isRatio(ratingPlan.BalancedRate, ratingPlan.UnbalancedRate) {
		chargingModels = append(chargingModels, engine.ChargingModel{
			Service:              service.Service,
			IncludedInCommitment: service.IncludedInCommitment,
			RatioPlan:            toEngineRatioPlan(ratingPlan.BalancedRate, ratingPlan.UnbalancedRate),
		})
	}
	return chargingModels
}

func toEngineRatingPlan(rate Rate, isBackToFirst bool) *engine.RatingPlan {
	var engineTiers []engine.Tier
	if len(rate.Thresholds) > 0 {
		engineTiers = toEngineTiers(rate.Thresholds)
	} else {
		engineTiers = []engine.Tier{engine.Tier{
			From:        0,
			To:          INF,
			FixedPrice:  rate.FixedPrice,
			LinearPrice: rate.LinearPrice,
		}}
	}
	return &engine.RatingPlan{IsBackToFirst: isBackToFirst, Tiers: engineTiers}
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

//Rate struct is always set to zero initially,
//hence checking against ratioBased. All non-
//ratio based are said to be rate based.
func isRate(ratingPlan RatingPlan) bool {
	return (len(ratingPlan.BalancedRate.Thresholds) == 0 || len(ratingPlan.UnbalancedRate.Thresholds) == 0)
}

//Keeping this redundant function for
//future models support
func isRatio(balanced Rate, unbalanced Rate) bool {
	return (len(balanced.Thresholds) > 0 || len(unbalanced.Thresholds) > 0)
}

func toEngineCondition(condition Condition) engine.Condition {
	switch c := condition.SelectedConditionName; c {
	case ContractRevenue:
		return engine.Condition{
			Type:           engine.ContractRevenue,
			Value:          condition.SelectedCondition.CommitmentsValue,
			IncludingTaxes: condition.SelectedCondition.IncludingTaxes,
		}
	case DealRevenue:
		return engine.Condition{
			Type:           engine.DiscountRevenue,
			Value:          condition.SelectedCondition.CommitmentsValue,
			IncludingTaxes: condition.SelectedCondition.IncludingTaxes,
		}
	default:
		return engine.Condition{Type: engine.Unconditional}
	}
}

func toEngineTiers(tiers []Tier) []engine.Tier {
	var engineTiers = make([]engine.Tier, len(tiers))
	if len(tiers) > 0 {
		to := INF
		for i := len(tiers) - 1; i >= 0; i-- {
			// engineFixedPrice, _ := strconv.ParseFloat(tiers[i].FixedPrice, 64)
			// engineLinearPrice, _ := strconv.ParseFloat(tiers[i].LinearPrice, 64)
			// engineFrom, _ := strconv.ParseFloat(tiers[i].Start, 64)
			engineTiers[i] = engine.Tier{
				FixedPrice:  tiers[i].FixedPrice,
				LinearPrice: tiers[i].LinearPrice,
				From:        tiers[i].Start,
				To:          to,
			}
			to = tiers[i].Start
		}
	}
	return engineTiers
}
