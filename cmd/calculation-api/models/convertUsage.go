package models

import (
	engine "github.com/GSMA-CPAS/BWRP-calculation-service/pkg/engine/charging"
)

const (
	inbound  = "inbound"
	outbound = "outbound"
)

// ConvertToEngineAggregatedUsage adds all the usage and creates an aggregatedUsage
// This is a map from a specific service,hometadig,visitortadig combination
func ConvertToEngineAggregatedUsage(usage Usage) engine.AggregatedUsage {
	var aggregatedUsage engine.AggregatedUsage = make(engine.AggregatedUsage)

	aggregatedUsageData := addDirectionToUsageData(usage.Inbound, inbound)
	aggregatedOutbound := addDirectionToUsageData(usage.Outbound, outbound)
	aggregatedUsageData = append(aggregatedUsageData, aggregatedOutbound...)

	for _, usage := range aggregatedUsageData {
		key := engine.ServiceTadig{Service: usage.Service, HomeTadig: usage.HomeTadig, VisitorTadig: usage.VisitorTadig}
		item, ok := aggregatedUsage[key]
		// v, _ := strconv.ParseFloat(usage.Volume, 64)
		// c, _ := strconv.ParseFloat(usage.Charge, 64)
		// t, _ := strconv.ParseFloat(usage.Tax, 64)
		if !ok {
			aggregatedUsage[key] = engine.Usage{
				Volume:    usage.Volume,
				Unit:      engine.Unit(usage.Unit),
				Charge:    usage.Charge,
				Tax:       usage.Tax,
				Direction: usage.Direction,
			}
		} else {
			aggregatedUsage[key] = engine.Usage{
				Volume:    item.Volume + usage.Volume,
				Charge:    item.Charge + usage.Charge,
				Tax:       item.Tax + usage.Tax,
				Direction: usage.Direction,
			}
		}
	}
	return aggregatedUsage
}

func addDirectionToUsageData(usages []UsageData, direction string) []UsageData {
	var aggregatedUsageData = make([]UsageData, 0)
	for _, usageData := range usages {
		usageData.Direction = direction
		aggregatedUsageData = append(aggregatedUsageData, usageData)
	}
	return aggregatedUsageData
}
