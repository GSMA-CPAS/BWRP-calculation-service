/*
 License-Placeholder
*/

package engine

// Result contains the results of the calculation.
type Result struct {
	IntermediateResults []IntermediateResult
}

type CommitmentResult struct {
	Party     string
	DealValue float32
}

// IntermediateResult contains the results per service
type IntermediateResult struct {
	Service       string
	HomeTadigs    []string
	VisitorTadigs []string
	DealValue     float32
}
