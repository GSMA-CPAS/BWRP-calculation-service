package engine

import (
	"encoding/json"
	"testing"
)

var model1 = Contract{
	Parts: []ContractPart{
		{
			[]ServiceGroup{{
				[]string{"HOR1"},
				[]string{"HOR2"},
				[]ChargingModel{
					{
						Service: SMS,
						RatingPlan: &RatingPlan{
							[]Tier{
								{From: 0, To: 1500, LinearPrice: 5},
							},
						},
					},
					{
						Service: MOOC,
						RatingPlan: &RatingPlan{
							[]Tier{
								{From: 0, To: 1500, FixedPrice: 1500},
							},
						},
					},
				},
			}}},
		{
			[]ServiceGroup{{
				[]string{"HOR2"},
				[]string{"HOR1"},
				[]ChargingModel{
					{
						Service: SMS,
						RatingPlan: &RatingPlan{
							[]Tier{
								{From: 0, To: 1500, FixedPrice: 5000},
							},
						},
					},
					{
						Service: MOOC,
						RatingPlan: &RatingPlan{
							[]Tier{
								{From: 0, To: 1500, FixedPrice: 3500},
							},
						},
					},
				},
			}}},
	},
}

var usage1 = AggregatedUsage{
	ServiceTadig{SMS, "HOR1", "HOR2"}:  Usage{Volume: 3000, Charge: 1007},
	ServiceTadig{SMS, "HOR2", "HOR1"}:  Usage{Volume: 2000, Charge: 1005},
	ServiceTadig{MOOC, "HOR1", "HOR2"}: Usage{Volume: 5000, Charge: 3003},
}

func TestScenario1(t *testing.T) {
	engine := CalculationEngine{}
	result := engine.Calculate(usage1, model1)
	js, _ := json.Marshal(result)
	t.Log(string(js))
}
