package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAggregateService(t *testing.T) {
	agg := AggregatedUsage{ServiceTadig{"SMS", "HOR1"}: Usage{Volume: 1000},
		ServiceTadig{"SMS", "HOR2"}: Usage{Volume: 1300},
		ServiceTadig{"SMS", "HOR3"}: Usage{Volume: 1500},
	}
	ok, result := agg.Aggregate("SMS", []string{"HOR1", "HOR2"})
	if !ok {
		t.Errorf("Result should be aggregated")
	}

	if result.Volume != 2300 {
		t.Errorf("Usage should add up to 2300")
	}

	assert.Equal(t, result.Tadigs, []string{"HOR1", "HOR2"})
}
