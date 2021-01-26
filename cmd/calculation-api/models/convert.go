package models

import engine "github.com/GSMA-CPAS/BWRP-calculation-service/pkg/engine/charging"

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
			Service:   Service(item.Service),
		}
		copy(ir[i].HomeTadigs, item.HomeTadigs)
		copy(ir[i].VisitorTadigs, item.VisitorTadigs)
	}
	return ir
}
