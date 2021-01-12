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
	Service       Service
	HomeTadigs    []string
	VisitorTadigs []string
	DealValue     int64
}

// Service is a string
type Service string
