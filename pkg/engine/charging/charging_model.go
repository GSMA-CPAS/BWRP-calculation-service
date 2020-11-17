package engine

type ChargingModel struct {
	Service    Service
	RatingPlan RatingPlan
	AccessPlan RatingPlan
}

type IntermediateResult struct {
	Service   Service
	DealValue int64
}

type Service string

const (
	SMS  Service = "SMS"
	MOOC Service = "MOOC"
)

//Calculate returns the intermediate result for a specfic service
func (c *ChargingModel) Calculate(usage Usage) IntermediateResult {
	dealValue := c.RatingPlan.Calculate(usage)
	return IntermediateResult{Service: c.Service, DealValue: dealValue}
}
