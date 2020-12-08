package engine

type CalculationEngine struct {
}

type ContractPart struct {
	ServiceGroups []ServiceGroup
}

type ServiceGroup struct {
	HomeTadigs     []string
	VisitorTadigs  []string
	ChargingModels []ChargingModel
}

type Contract struct {
	Parts []ContractPart
}

//Calculate returns the deal value by inputting the aggregate-usage and a contract
//Go through all the contract parts and charging models
func (c *CalculationEngine) Calculate(aggUsage AggregatedUsage, contract Contract) Result {
	result := Result{IntermediateResults: make([]IntermediateResult, 0)}
	for _, part := range contract.Parts {
		for _, group := range part.ServiceGroups {
			for _, model := range group.ChargingModels {
				h := aggUsage.Aggregate(model.Service, group.HomeTadigs, group.VisitorTadigs)
				v := aggUsage.Aggregate(model.Service, group.VisitorTadigs, group.HomeTadigs)
				intermediateResult := model.Calculate(h, v)
				intermediateResult.Service = model.Service
				intermediateResult.HomeTadigs = group.HomeTadigs
				intermediateResult.VisitorTadigs = group.VisitorTadigs
				result.IntermediateResults = append(result.IntermediateResults, intermediateResult)
			}

		}
	}
	return result
}
