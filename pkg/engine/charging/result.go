/*
 License-Placeholder
*/

package engine

// Result contains the results of the calculation.
type Result struct {
	Total               []Deal
	IntermediateResults []IntermediateResult
}

type Deal struct {
	HomeTadigs        []string
	VisitorTadigs     []string
	CommitmentReached bool
	Value             float64
}

// IntermediateResult contains the results per service
type IntermediateResult struct {
	Service       string
	HomeTadigs    []string
	VisitorTadigs []string
	DealValue     float64
	Direction     string
}
