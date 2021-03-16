/*
 License-Placeholder
*/

package engine

// Result contains the results of the calculation.
type Result struct {
	IntermediateResults []IntermediateResult
}

// IntermediateResult contains the results per service
type IntermediateResult struct {
	Service       string
	HomeTadigs    []string
	VisitorTadigs []string
	DealValue     float64
	Direction     string
}
