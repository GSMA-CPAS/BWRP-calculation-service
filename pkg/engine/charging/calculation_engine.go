package engine

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
		Total:               make([]Deal, 0),
		IntermediateResults: make([]IntermediateResult, 0),
	}

	for _, part := range contract.Parts {
		var inCommitmentValue, aggregatedChargeHome, aggregatedDealValue float64
		partIntermediateResults := make([]IntermediateResult, 0)
		homeTadigs := make([]string, 0)
		visitorTadigs := make([]string, 0)
		for _, group := range part.ServiceGroups {
			homeTadigs = append(homeTadigs, group.HomeTadigs...)
			visitorTadigs = append(visitorTadigs, group.VisitorTadigs...)
			for _, model := range group.ChargingModels {
				var intermediateResult IntermediateResult
				h := aggUsage.Aggregate(model.Service, group.HomeTadigs, group.VisitorTadigs)
				if part.Condition.Type == ContractRevenue {
					aggregatedChargeHome += h.Charge
					if part.Condition.IncludingTaxes == true {
						aggregatedChargeHome += h.Tax
					}
				}
				v := aggUsage.Aggregate(model.Service, group.VisitorTadigs, group.HomeTadigs)
				intermediateResult = model.Calculate(h, v)
				intermediateResult.Volume = h.Volume
				intermediateResult.Service = model.Service
				intermediateResult.HomeTadigs = group.HomeTadigs
				intermediateResult.VisitorTadigs = group.VisitorTadigs
				aggregatedDealValue += intermediateResult.DealValue
				if len(h.Direction) > 0 {
					intermediateResult.Direction = h.Direction
					partIntermediateResults = append(partIntermediateResults, intermediateResult)
				}
				if model.IncludedInCommitment {
					inCommitmentValue += intermediateResult.DealValue
				}
			}
		}
		result.IntermediateResults = append(result.IntermediateResults, partIntermediateResults...)
		if part.Condition.Type == DiscountRevenue && inCommitmentValue < part.Condition.Value {
			result.Total = append(result.Total, Deal{HomeTadigs: toUniqueTadigs(homeTadigs), VisitorTadigs: toUniqueTadigs(visitorTadigs), CommitmentReached: false, Value: part.Condition.Value})
		} else if part.Condition.Type == ContractRevenue && aggregatedChargeHome < part.Condition.Value {
			result.Total = append(result.Total, Deal{HomeTadigs: toUniqueTadigs(homeTadigs), VisitorTadigs: toUniqueTadigs(visitorTadigs), CommitmentReached: false, Value: part.Condition.Value})
		} else {
			result.Total = append(result.Total, Deal{HomeTadigs: toUniqueTadigs(homeTadigs), VisitorTadigs: toUniqueTadigs(visitorTadigs), CommitmentReached: true, Value: aggregatedDealValue})
		}

	}
	return result
}

func toUniqueTadigs(tadigs []string) []string {
	keys := make(map[string]bool)
	list := make([]string, 0)
	for _, tadig := range tadigs {
		if _, value := keys[tadig]; !value {
			keys[tadig] = true
			list = append(list, tadig)
		}
	}
	return list
}
