package engine

// RatioPlan is the type to hold both a balanced and an unbalanced plan.
type RatioPlan struct {
	Balanced   *RatingPlan
	Unbalanced *RatingPlan
}

// Calculate , calculates the deal value for the ratioplan.
func (rp *RatioPlan) Calculate(h Usage, v Usage) int64 {
	balanced := min(h.Volume, v.Volume)
	unbalanced := max(0, h.Volume-balanced)
	if unbalanced > 0 {
		return rp.Balanced.Calculate(balanced, h.Unit) + rp.Unbalanced.Calculate(unbalanced, h.Unit)
	}
	return rp.Balanced.Calculate(balanced, h.Unit)
}
