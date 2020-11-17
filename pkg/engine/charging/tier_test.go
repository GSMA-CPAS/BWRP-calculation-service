package engine

import (
	"testing"
)

func TestMatchThresholdOver(t *testing.T) {
	tier := Tier{From: 0, To: 500, FixedPrice: 0, LinearPrice: 2}
	result := tier.Calculate(1000, 0)

	if result != 1000 {
		t.Errorf("Result should be 1000, 2*500")
	}
}

func TestMatchThresholdBelow(t *testing.T) {
	tier := Tier{From: 0, To: 500, FixedPrice: 0, LinearPrice: 2}
	result := tier.Calculate(400, 0)

	if result != 800 {
		t.Errorf("Result should be 800, 2*400")
	}
}

func TestOutsideThreshold(t *testing.T) {
	tier := Tier{From: 1500, To: 2000, FixedPrice: 0, LinearPrice: 2}
	result := tier.Calculate(1000, 0)

	if result != 0 {
		t.Errorf("Result should be 0, because 1000 < 1500")
	}
}

func TestInfThreshold(t *testing.T) {
	tier := Tier{From: 0, To: INF, FixedPrice: 0, LinearPrice: 2}
	result := tier.Calculate(3000, 0)

	if result != 6000 {
		t.Errorf("Result should be 6000, 2*3000")
	}
}

func TestIncludingFP(t *testing.T) {
	tier := Tier{From: 0, To: 1500, FixedPrice: 100, LinearPrice: 2}
	result := tier.Calculate(1000, 0)

	if result != 2100 {
		t.Errorf("Result should be 6000, 100+2*1000")
	}
}

func TestBoundaryTo(t *testing.T) {
	tier := Tier{From: 0, To: 500, FixedPrice: 0, LinearPrice: 2}
	result := tier.Calculate(500, 0)

	if result != 1000 {
		t.Errorf("Result should be 1000, 2*500")
	}
}

func TestBoundaryFrom(t *testing.T) {
	tier := Tier{From: 500, To: 1000, FixedPrice: 0, LinearPrice: 2}
	result := tier.Calculate(1000, 0)

	if result != 1000 {
		t.Errorf("Result should be 1000, 2*500")
	}
}
