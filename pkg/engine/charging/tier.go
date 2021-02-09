/*
 License-Placeholder
*/

package engine

//Tier defines a calculation function with a threshold
type Tier struct {
	From        int64
	To          int64
	FixedPrice  int64
	LinearPrice int64
	LinearUnit  Unit
}

//Calculate performs the rating formula and also takes care of the threshold.
//when the volume is outside of the range of the threshold the result is 0, this way
//a simpel sum over all tiers can be made.
func (t *Tier) Calculate(volume int64, unit Unit) int64 {
	var adjustedVolume int64
	deltaThreshold := t.To - t.From

	if deltaThreshold > 0 {
		adjustedVolume = max(0, min((volume-t.From), deltaThreshold))
	} else {
		adjustedVolume = max(0, (volume - t.From))
	}

	if adjustedVolume == 0 {
		return 0
	}
	rate := t.FixedPrice + t.LinearPrice*adjustedVolume
	return rate
}

func max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}
