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
		Deal:                make(map[string]bool, len(contract.Parts)),
		IntermediateResults: make([]IntermediateResult, 0),
	}

	for _, part := range contract.Parts {
		var inCommitmentValue float64 = 0
		var aggregatedChargeHome float64 = 0
		partIntermediateResults := make([]IntermediateResult, 0)
		for _, group := range part.ServiceGroups {
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
				intermediateResult.Service = model.Service
				intermediateResult.HomeTadigs = group.HomeTadigs
				intermediateResult.VisitorTadigs = group.VisitorTadigs
				if len(h.Direction) > 0 {
					intermediateResult.Direction = h.Direction
					partIntermediateResults = append(partIntermediateResults, intermediateResult)
				}
				if model.IncludedInCommitment {
					inCommitmentValue += intermediateResult.DealValue
				}
			}
		}
		if part.Condition.Type == DiscountRevenue && inCommitmentValue < part.Condition.Value {
			result.Deal[part.Party] = false
			// partIntermediateResults = []IntermediateResult{}
		} else if part.Condition.Type == ContractRevenue && aggregatedChargeHome < part.Condition.Value {
			result.Deal[part.Party] = false
			// partIntermediateResults = []IntermediateResult{}
		} else {
			result.Deal[part.Party] = true
		}
		result.IntermediateResults = append(result.IntermediateResults, partIntermediateResults...)
	}
	return result
}
