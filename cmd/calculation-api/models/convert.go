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
			Service:   string(item.Service),
		}
		copy(ir[i].HomeTadigs, item.HomeTadigs)
		copy(ir[i].VisitorTadigs, item.VisitorTadigs)
	}
	return ir
}

// ConvertToEngineAggregatedUsage adds all the usage and creates an aggregatedUsage
// This is a map from a specific service,hometadig,visitortadig combination
func ConvertToEngineAggregatedUsage(usages []Usage) engine.AggregatedUsage {
	var aggregatedUsage engine.AggregatedUsage = make(engine.AggregatedUsage)
	for _, usage := range usages {
		key := engine.ServiceTadig{Service: engine.Service(usage.Service), HomeTadig: usage.HomeTadig, VisitorTadig: usage.VisitorTadig}
		item, ok := aggregatedUsage[key]
		if !ok {
			aggregatedUsage[key] = engine.Usage{Volume: usage.Volume, Unit: engine.Unit(usage.Unit), Charge: usage.Charge, Tax: usage.Tax}
		} else {
			aggregatedUsage[key] = engine.Usage{Volume: item.Volume + usage.Volume, Charge: item.Charge + usage.Charge, Tax: usage.Tax + usage.Tax}
		}
	}
	return aggregatedUsage
}
