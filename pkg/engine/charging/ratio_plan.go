package engine

// RatioPlan is the type to hold both a balanced and an unbalanced plan.
type RatioPlan struct {
	BalancedRate   []Tier
	UnbalancedRate []Tier
}

// Calculate , calculates the deal value for the ratioplan.
func (rp *RatioPlan) Calculate(h Usage, v Usage, unit Unit) float64 {
	balancedVolume := min(h.Volume, v.Volume)
	unbalancedVolume := max(0, h.Volume-balancedVolume)
	result := calculate(rp.BalancedRate, balancedVolume, unit)
	if unbalancedVolume > 0 {
		result += calculate(rp.UnbalancedRate, unbalancedVolume, unit)
	}
	return result
}

//Calculate sums the result of all tiers over the input usage.
func calculate(tiers []Tier, vol float64, unit Unit) float64 {
	sum := float64(0)
	for _, tier := range tiers {
		sum += tier.Calculate(vol, unit)
	}
	return sum
}
