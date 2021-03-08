package models

import (
	"strconv"

	engine "github.com/GSMA-CPAS/BWRP-calculation-service/pkg/engine/charging"
)

// ConvertToEngineAggregatedUsage adds all the usage and creates an aggregatedUsage
// This is a map from a specific service,hometadig,visitortadig combination
func ConvertToEngineAggregatedUsage(usages []UsageData) engine.AggregatedUsage {
	var aggregatedUsage engine.AggregatedUsage = make(engine.AggregatedUsage)
	for _, usage := range usages {
		key := engine.ServiceTadig{Service: usage.Service, HomeTadig: usage.HomeTadig, VisitorTadig: usage.VisitorTadig}
		item, ok := aggregatedUsage[key]
		v, _ := strconv.ParseFloat(usage.Volume, 64)
		c, _ := strconv.ParseFloat(usage.Charge, 64)
		t, _ := strconv.ParseFloat(usage.Tax, 64)
		if !ok {
			aggregatedUsage[key] = engine.Usage{Volume: v, Unit: engine.Unit(usage.Unit), Charge: c, Tax: t}
		} else {
			aggregatedUsage[key] = engine.Usage{Volume: item.Volume + v, Charge: item.Charge + c, Tax: item.Tax + t}
		}
	}
	return aggregatedUsage
}
