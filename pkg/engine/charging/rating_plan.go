package engine

//RatingPlan defines the tiers that go into a rating plan.
type RatingPlan struct {
	Tiers []Tier
}

//Usage specifies the usage for a single service
type Usage struct {
	Volume int64
	Unit   Unit
}

//Calculate sums the result of all tiers over the input usage.
func (r *RatingPlan) Calculate(usage Usage) int64 {
	sum := int64(0)
	for _, tier := range r.Tiers {
		sum += tier.Calculate(usage.Volume, usage.Unit)
	}
	return sum
}
