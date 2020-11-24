package engine

var model1 = Contract{
	Parts: []ContractPart{
		{
			[]string{"HOR1"},
			[]ChargingModel{
				{
					Service: SMS,
					RatingPlan: RatingPlan{
						[]Tier{
							{From: 0, To: 1500, LinearPrice: 5},
						},
					},
				},
			},
		},
		{
			[]string{"HOR2"},
			[]ChargingModel{
				{
					Service: SMS,
					RatingPlan: RatingPlan{
						[]Tier{
							{From: 0, To: 1500, FixedPrice: 5000},
						},
					},
				},
			},
		},
	},
}

var usage1 = AggregatedUsage{
	ServiceTadig{SMS, "HOR1"}: Usage{Volume: 3000},
	ServiceTadig{SMS, "HOR2"}: Usage{Volume: 3000},
}
