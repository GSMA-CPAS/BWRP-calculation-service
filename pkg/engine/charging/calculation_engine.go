package engine

import (
	"math"
)

//Uncondtional, ContractRevenue and DiscountRevenue are valid contract conditions
const (
	Unconditional = iota
	ContractRevenue
	DiscountRevenue
)

// CalculationEngine is the type for the engine.
type CalculationEngine struct {
}

// ContractPart , defines a part of the contract containing the services and chargingmodels.
type ContractPart struct {
	Party         string
	Condition     Condition
	ServiceGroups []ServiceGroup
}

// Condition contains the discount condition
type Condition struct {
	Type           int
	Value          float64
	Currency       string
	IncludingTaxes bool
}

// ServiceGroup is a grouping between tadigs and chargingmodels.
type ServiceGroup struct {
	HomeTadigs     []string
	VisitorTadigs  []string
	ChargingModels []ChargingModel
}

// Contract is the overarching contract type.
type Contract struct {
	Parts []ContractPart
}

//Calculate returns the deal value by inputting the aggregate-usage and a contract
//Go through all the contract parts and charging models
//Take Contract Revenue and Deal Revenue commitment conditions into account. If condition is not met then no deal
func (c *CalculationEngine) Calculate(aggUsage AggregatedUsage, contract Contract) Result {
	result := Result{
		IntermediateResults: make([]IntermediateResult, 0),
	}
	for _, part := range contract.Parts {
		var inCommitmentValue float64 = 0
		var aggregatedChargeHome float64 = 0
		var shortFromCommitment float64 = 0
		var sumOfAllIncludedDeal float64 = 0
		partIntermediateResults := make([]IntermediateResult, 0)
		for _, group := range part.ServiceGroups {
			for _, model := range group.ChargingModels {
				var intermediateResult IntermediateResult
				h := aggUsage.Aggregate(model.Service, group.HomeTadigs, group.VisitorTadigs)
				if part.Condition.Type == ContractRevenue {
					aggregatedChargeHome += h.Charge
					if part.Condition.IncludingTaxes {
						aggregatedChargeHome += h.Tax
					}
				}
				v := aggUsage.Aggregate(model.Service, group.VisitorTadigs, group.HomeTadigs)
				intermediateResult = model.Calculate(h, v)
				intermediateResult.Volume = h.Volume
				intermediateResult.Service = model.Service
				intermediateResult.HomeTadigs = group.HomeTadigs
				intermediateResult.VisitorTadigs = group.VisitorTadigs
				intermediateResult.Direction = h.Direction

				if model.IncludedInCommitment {
					inCommitmentValue += intermediateResult.DealValue
					intermediateResult.IsIncluded = true
				}
				partIntermediateResults = append(partIntermediateResults, intermediateResult)
			}
			if part.Condition.Type == DiscountRevenue {
				shortFromCommitment = part.Condition.Value - inCommitmentValue
				sumOfAllIncludedDeal = inCommitmentValue
			}
			if part.Condition.Type == ContractRevenue {
				shortFromCommitment = (part.Condition.Value - aggregatedChargeHome)
				sumOfAllIncludedDeal = aggregatedChargeHome
			}
			//Ratio Based Short of Commitment
			updateSoC(partIntermediateResults, shortFromCommitment, sumOfAllIncludedDeal)
		}
		result.IntermediateResults = append(result.IntermediateResults, partIntermediateResults...)
	}
	return result
}

func updateSoC(partIntermediateResults []IntermediateResult, shortage float64, sumOfAllIncludedDeal float64) []IntermediateResult {
	// Looping over all services to update shortage based on deal value ratio
	for i := range partIntermediateResults {
		if !partIntermediateResults[i].IsIncluded {
			continue
		}
		partIntermediateResults[i].ShortOfCommitment = fixed(shortage*(partIntermediateResults[i].DealValue/sumOfAllIncludedDeal), 4)
	}
	return partIntermediateResults
}

//Functions for roundings off the shortage
func fixed(i float64, points int) float64 {
	r := math.Pow(10, float64(points))
	return float64(roundOff(i*r)) / r
}

func roundOff(i float64) int {
	return int(i + math.Copysign(0.2, i))
}
