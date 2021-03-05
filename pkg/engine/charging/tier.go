/*
 License-Placeholder
*/

package engine

//Tier defines a calculation function with a threshold
type Tier struct {
	From        float64
	To          float64
	FixedPrice  float64
	LinearPrice float64
	Unit        Unit
}

//Calculate performs the rating formula and also takes care of the threshold.
//when the volume is outside of the range of the threshold the result is 0, this way
//a simpel sum over all tiers can be made.
func (t *Tier) Calculate(volume float64, unit Unit) float64 {
	deltaThreshold := t.To - t.From
	adjustedVolume := max(0, min((volume-t.From), deltaThreshold))
	if adjustedVolume <= 0 {
		return 0
	}
	rate := t.FixedPrice + t.LinearPrice*adjustedVolume
	return rate
}

func max(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}

func min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}
