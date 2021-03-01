package engine

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

const INF = math.MaxInt64

func TestSumTiers(t *testing.T) {
	r := RatingPlan{[]Tier{
		{From: 0, To: 1000, FixedPrice: 0, LinearPrice: 2},
		{From: 1000, To: 1500, FixedPrice: 100, LinearPrice: 0},
		{From: 1500, To: INF, FixedPrice: 100, LinearPrice: 2},
	}}
	result := r.Calculate(2500, "min")
	assert.Equal(t, int64(4200), result, "Result should add up to 4200, but was %v", result)
}
