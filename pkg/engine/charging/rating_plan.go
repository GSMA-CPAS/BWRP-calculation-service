package engine

//RatingPlan defines the tiers that go into a rating plan.
type RatingPlan struct {
	Tiers []Tier
}

//Calculate sums the result of all tiers over the input usage.
func (r *RatingPlan) Calculate(vol float64, unit Unit) float64 {
	sum := float64(0)
	for _, tier := range r.Tiers {
		sum += tier.Calculate(vol, unit)
	}
	return sum
}
