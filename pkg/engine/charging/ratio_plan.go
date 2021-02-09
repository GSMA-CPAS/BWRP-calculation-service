package engine

// RatioPlan is the type to hold both a balanced and an unbalanced plan.
type RatioPlan struct {
	BalancedRate   int64
	UnbalancedRate int64
}

// Calculate , calculates the deal value for the ratioplan.
func (rp *RatioPlan) Calculate(h Usage, v Usage) int64 {
	balanced := min(h.Volume, v.Volume)
	unbalanced := max(0, h.Volume-balanced)
	if unbalanced > 0 {
		return rp.BalancedRate*balanced + rp.UnbalancedRate*unbalanced
	}
	return rp.BalancedRate * balanced
}
