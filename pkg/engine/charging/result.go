/*
 License-Placeholder
*/

package engine

// Result contains the results of the calculation.
type Result struct {
	ContractCommitmentResult []CommitmentResult
	DiscountCommitmentResult []CommitmentResult
	IntermediateResults      []IntermediateResult
}

type CommitmentResult struct {
	Party     string
	DealValue int64
}

// IntermediateResult contains the results per service
type IntermediateResult struct {
	Service       string
	HomeTadigs    []string
	VisitorTadigs []string
	DealValue     int64
}
