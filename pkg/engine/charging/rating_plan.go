package engine

//RatingPlan defines the tiers that go into a rating plan.
type RatingPlan struct {
	Tiers []Tier
}

//Calculate sums the result of all tiers over the input usage.
func (r *RatingPlan) Calculate(h Usage, v Usage) int64 {
	sum := int64(0)
	for _, tier := range r.Tiers {
		sum += tier.Calculate(h.Volume, h.Unit)
	}
	return sum
}
