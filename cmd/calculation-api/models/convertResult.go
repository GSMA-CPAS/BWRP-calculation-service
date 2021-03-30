package models

import (
	engine "github.com/GSMA-CPAS/BWRP-calculation-service/pkg/engine/charging"
)

// ConvertFromEngineResult convert from the engine type to the api type
func ConvertFromEngineResult(result engine.Result) Result {
	return Result{
		Total:               ConvertFromEngineTotal(result.Total),
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
			Usage:     item.Volume,
		}
		ir[i].HomeTadigs = item.HomeTadigs
		ir[i].VisitorTadigs = item.VisitorTadigs
		ir[i].Type = item.Direction
	}
	return ir
}

// ConvertFromEngineTotal convert the total result form the engine to the API model
func ConvertFromEngineTotal(total []engine.Deal) []Deal {
	result := make([]Deal, 0)
	for _, v := range total {
		result = append(result, Deal{HomeTadigs: v.HomeTadigs, VisitorTadigs: v.VisitorTadigs, CommitmentReached: v.CommitmentReached, Value: v.Value})
	}
	return result
}
