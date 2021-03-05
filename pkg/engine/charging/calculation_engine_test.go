package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var model1 = Contract{
	Parts: []ContractPart{
		{
			"HOME",
			Condition{Type: 0},
			[]ServiceGroup{{
				[]string{"HOR1"},
				[]string{"HOR2"},
				[]ChargingModel{
					{
						Service: "SMS",
						RatingPlan: &RatingPlan{
							false,
							[]Tier{
								{From: 0, To: 1500, LinearPrice: 5},
							},
						},
					},
					{
						Service: "MOOC",
						RatingPlan: &RatingPlan{
							false,
							[]Tier{
								{From: 0, To: 1500, FixedPrice: 1500},
							},
						},
					},
				},
			}}},
		{
			"VISITOR",
			Condition{Type: 0},
			[]ServiceGroup{{
				[]string{"HOR2"},
				[]string{"HOR1"},
				[]ChargingModel{
					{
						Service: "SMS",
						RatingPlan: &RatingPlan{
							false,
							[]Tier{
								{From: 0, To: 1500, FixedPrice: 5000},
							},
						},
					},
					{
						Service: "MOOC",
						RatingPlan: &RatingPlan{
							false,
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
	ServiceTadig{"SMS", "HOR1", "HOR2"}:  Usage{Volume: 3000, Charge: 1007},
	ServiceTadig{"SMS", "HOR2", "HOR1"}:  Usage{Volume: 2000, Charge: 1005},
	ServiceTadig{"MOOC", "HOR1", "HOR2"}: Usage{Volume: 5000, Charge: 3003},
}

var result1 = Result{
	IntermediateResults: []IntermediateResult{
		{Service: "SMS", HomeTadigs: []string{"HOR1"}, VisitorTadigs: []string{"HOR2"}, DealValue: 7500},
		{Service: "MOOC", HomeTadigs: []string{"HOR1"}, VisitorTadigs: []string{"HOR2"}, DealValue: 1500},
		{Service: "SMS", HomeTadigs: []string{"HOR2"}, VisitorTadigs: []string{"HOR1"}, DealValue: 5000},
		{Service: "MOOC", HomeTadigs: []string{"HOR2"}, VisitorTadigs: []string{"HOR1"}, DealValue: 0},
	}}

func TestScenario1(t *testing.T) {
	engine := CalculationEngine{}
	result := engine.Calculate(usage1, model1)
	assert.Equal(t, result1, result)
}
