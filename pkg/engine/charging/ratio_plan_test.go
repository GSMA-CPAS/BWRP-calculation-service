package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllBalanced(t *testing.T) {
	bal := RatingPlan{[]Tier{
		{From: 0, To: INF, FixedPrice: 0, LinearPrice: 2},
	}}

	unbal := RatingPlan{[]Tier{
		{From: 0, To: INF, FixedPrice: 0, LinearPrice: 4},
	}}

	ratio := RatioPlan{Balanced: &bal, Unbalanced: &unbal}
	deal := ratio.Calculate(Usage{Volume: 1000, Charge: 1000}, Usage{Volume: 1000, Charge: 1000})
	assert.Equal(t, int64(2000), deal)
}

func TestBalancedUnbalanced(t *testing.T) {
	bal := RatingPlan{[]Tier{
		{From: 0, To: INF, FixedPrice: 0, LinearPrice: 2},
	}}

	unbal := RatingPlan{[]Tier{
		{From: 0, To: INF, FixedPrice: 0, LinearPrice: 4},
	}}

	ratio := RatioPlan{Balanced: &bal, Unbalanced: &unbal}
	deal := ratio.Calculate(Usage{Volume: 1500, Charge: 1000}, Usage{Volume: 1000, Charge: 1000})
	assert.Equal(t, int64(4000), deal)
}
