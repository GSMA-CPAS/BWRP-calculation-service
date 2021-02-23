package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllBalanced(t *testing.T) {
	var unit Unit
	br := []Tier{
		{From: 0, To: 1000, FixedPrice: 0, LinearPrice: 2},
		{From: 1000, To: 1500, FixedPrice: 100, LinearPrice: 0},
		{From: 1500, To: INF, FixedPrice: 100, LinearPrice: 2},
	}

	ur := []Tier{
		{From: 0, To: 1000, FixedPrice: 0, LinearPrice: 5},
		{From: 1000, To: 1500, FixedPrice: 100, LinearPrice: 0},
		{From: 1500, To: INF, FixedPrice: 100, LinearPrice: 5},
	}

	ratio := RatioPlan{BalancedRate: br, UnbalancedRate: ur}
	deal := ratio.Calculate(Usage{Volume: 5000, Charge: 5000}, Usage{Volume: 5000, Charge: 5000}, unit)
	assert.Equal(t, float32(9200), deal)
}

func TestBalancedUnbalanced(t *testing.T) {
	var unit Unit
	br := []Tier{
		{From: 0, To: 1000, FixedPrice: 0, LinearPrice: 2},
		{From: 1000, To: 1500, FixedPrice: 100, LinearPrice: 0},
		{From: 1500, To: INF, FixedPrice: 100, LinearPrice: 2},
	}

	ur := []Tier{
		{From: 0, To: 1000, FixedPrice: 0, LinearPrice: 5},
		{From: 1000, To: 1500, FixedPrice: 100, LinearPrice: 0},
		{From: 1500, To: INF, FixedPrice: 100, LinearPrice: 5},
	}

	ratio := RatioPlan{BalancedRate: br, UnbalancedRate: ur}
	deal := ratio.Calculate(Usage{Volume: 5000, Charge: 5000}, Usage{Volume: 3000, Charge: 3000}, unit)
	assert.Equal(t, float32(12900), deal)
}
