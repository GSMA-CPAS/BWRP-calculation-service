package engine

type CalculationEngine struct {
}

type ServiceTadig struct {
	Service string
	Tadig   string
}
type AggregatedUsage map[ServiceTadig]Usage

func (a *AggregatedUsage) GetForService(Service) (bool, Usage) {
	return true, Usage{}
}

type ContractPart struct {
	ChargingModels []ChargingModel
}

type Contract struct {
	Parts []ContractPart
}

type Result struct {
	IntermediateResults []IntermediateResult
}

//Calculate returns the deal value by inputting the aggregate-usage and a contract
func (c *CalculationEngine) Calculate(usage AggregatedUsage, contract Contract) Result {
	result := Result{IntermediateResults: make([]IntermediateResult, 0)}
	for _, part := range contract.Parts {
		for _, model := range part.ChargingModels {
			if ok, usage := usage.GetForService(model.Service); ok {
				intermediateResult := model.Calculate(usage)
				result.IntermediateResults = append(result.IntermediateResults, intermediateResult)
			}

		}
	}
	return result
}
