package engine

// ChargingModel defines the ratingplans for a service.
type ChargingModel struct {
	Service              string
	IncludedInCommitment bool
	RatingPlan           *RatingPlan
	RatioPlan            *RatioPlan
	AccessPlan           *RatingPlan
}

//HasRatioPlan checks if the rating should be done using the balanced/unbalanced model
func (c *ChargingModel) HasRatioPlan() bool {
	return c.RatioPlan != nil
}

//Calculate returns the intermediate result for a specfic service
func (c *ChargingModel) Calculate(h Usage, v Usage) IntermediateResult {
	var dealValue float64
	if c.HasRatioPlan() {
		dealValue = c.RatioPlan.Calculate(h, v, v.Unit)
	} else {
		dealValue = c.RatingPlan.Calculate(h.Volume, v.Unit)
	}

	return IntermediateResult{Service: c.Service, DealValue: dealValue}
}
