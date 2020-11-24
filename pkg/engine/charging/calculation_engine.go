package engine

type CalculationEngine struct {
}

type ContractPart struct {
	Tadigs         []string //TODO services can override this
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
		for _, model := range part.ChargingModels {
			if ok, usage := aggUsage.Aggregate(model.Service, part.Tadigs); ok {
				intermediateResult := model.Calculate(usage)
				intermediateResult.Service = model.Service
				intermediateResult.Tadigs = part.Tadigs
				result.IntermediateResults = append(result.IntermediateResults, intermediateResult)
			}

		}
	}
	return result
}
