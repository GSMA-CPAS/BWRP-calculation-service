package engine

type ChargingModel struct {
	Service        Service
	RatingPlan     *RatingPlan
	UnbalancedPlan *RatingPlan
	AccessPlan     *RatingPlan
}

func (c *ChargingModel) HasRatioPlan() bool {
	return c.UnbalancedPlan != nil
}

const (
	SMS  Service = "SMS"
	MOOC Service = "MOOC"
)

//Calculate returns the intermediate result for a specfic service
func (c *ChargingModel) Calculate(usage Usage) IntermediateResult {
	dealValue := c.RatingPlan.Calculate(usage)
	return IntermediateResult{Service: c.Service, DealValue: dealValue}
}
