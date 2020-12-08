package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAggregateService(t *testing.T) {
	agg := AggregatedUsage{ServiceTadig{"SMS", "HOR1", "HOR3"}: Usage{Volume: 1000},
		ServiceTadig{"SMS", "HOR2", "HOR3"}: Usage{Volume: 1300},
		ServiceTadig{"SMS", "HOR3", "HOR1"}: Usage{Volume: 1500},
	}
	ok, result := agg.Aggregate("SMS", []string{"HOR1", "HOR2"}, []string{"HOR3"})
	if !ok {
		t.Errorf("Result should be aggregated")
	}

	if result.Volume != 2300 {
		t.Errorf("Usage should add up to 2300")
	}

	assert.Equal(t, result.HomeTadigs, []string{"HOR1", "HOR2"})
}
