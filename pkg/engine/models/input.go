/*
 License-Placeholder
*/

package engine

type AggregatedUsage struct{}

type Usage struct{}

type CalculateCharging interface {
	Charge() int64
}

type Contract struct {
	ChargingModels ChargingModels
}

type ChargingModels []CalculateCharging
