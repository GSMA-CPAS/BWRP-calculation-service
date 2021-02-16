package engine

import "fmt"

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
	Value          int64
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
func (c *CalculationEngine) Calculate(aggUsage AggregatedUsage, contract Contract) Result {
	fmt.Printf("%+v\n", contract)
	result := Result{
		ContractCommitmentResult: make([]CommitmentResult, 0),
		DiscountCommitmentResult: make([]CommitmentResult, 0),
		IntermediateResults:      make([]IntermediateResult, 0),
	}
	for _, part := range contract.Parts {
		if part.Condition.Type == ContractRevenue {
			result.ContractCommitmentResult = append(result.ContractCommitmentResult, calcContractDeal(aggUsage, part))
		} else {
			var inCommitmentValue int64 = 0
			for _, group := range part.ServiceGroups {
				for _, model := range group.ChargingModels {
					h := aggUsage.Aggregate(model.Service, group.HomeTadigs, group.VisitorTadigs)
					v := aggUsage.Aggregate(model.Service, group.VisitorTadigs, group.HomeTadigs)
					intermediateResult := model.Calculate(h, v)
					intermediateResult.Service = model.Service
					intermediateResult.HomeTadigs = group.HomeTadigs
					intermediateResult.VisitorTadigs = group.VisitorTadigs
					result.IntermediateResults = append(result.IntermediateResults, intermediateResult)
					if model.IncludedInCommitment {
						inCommitmentValue += intermediateResult.DealValue
					}
				}
			}
			if part.Condition.Type == DiscountRevenue {
				result.DiscountCommitmentResult = append(result.DiscountCommitmentResult,
					CommitmentResult{Party: part.Party, DealValue: calcDiscountDeal(inCommitmentValue, part.Condition.Value)})
			}
		}
	}
	return result
}

func calcContractDeal(aggUsage AggregatedUsage, part ContractPart) CommitmentResult {
	var aggregatedChargeHome int64
	for _, group := range part.ServiceGroups {
		for _, model := range group.ChargingModels {
			h := aggUsage.Aggregate(model.Service, group.HomeTadigs, group.VisitorTadigs)
			aggregatedChargeHome += h.Charge
		}
	}
	result := max(part.Condition.Value, aggregatedChargeHome)
	return CommitmentResult{Party: part.Party, DealValue: result}
}

func calcDiscountDeal(includedValue int64, commitValue int64) int64 {
	return max(includedValue, commitValue)
}
