package engine

type ChargingModel struct {
	Service    Service
	RatingPlan *RatingPlan
	RatioPlan  *RatioRating
	AccessPlan *RatingPlan
}

type RatioRating struct {
	Balanced   int64
	Unbalanced int64
}

//If Ration plan is not empty the usage is rated according to the balanced/unbalanced model
func (c *ChargingModel) HasRatioPlan() bool {
	return c.RatioPlan != nil
}

const (
	SMS  Service = "SMS"
	MOOC Service = "MOOC"
)

//Calculate returns the intermediate result for a specfic service
func (c *ChargingModel) Calculate(h Usage, v Usage) IntermediateResult {
	var dealValue int64
	if c.HasRatioPlan() {

	} else {
		dealValue = c.RatingPlan.Calculate(h, v)
	}

	return IntermediateResult{Service: c.Service, DealValue: dealValue}
}
