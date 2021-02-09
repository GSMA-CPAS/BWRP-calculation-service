package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllBalanced(t *testing.T) {
	var bal int64 = 10
	var unbal int64 = 20

	ratio := RatioPlan{BalancedRate: bal, UnbalancedRate: unbal}
	deal := ratio.Calculate(Usage{Volume: 1000, Charge: 1000}, Usage{Volume: 1000, Charge: 1000})
	assert.Equal(t, int64(10000), deal)
}

func TestBalancedUnbalanced(t *testing.T) {
	var bal int64 = 10
	var unbal int64 = 20

	ratio := RatioPlan{BalancedRate: bal, UnbalancedRate: unbal}
	deal := ratio.Calculate(Usage{Volume: 1500, Charge: 1000}, Usage{Volume: 1000, Charge: 1000})
	assert.Equal(t, int64(20000), deal)
}
