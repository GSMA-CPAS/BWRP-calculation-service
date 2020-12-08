/*
 License-Placeholder
*/

package engine

type Result struct {
	IntermediateResults []IntermediateResult
}

type IntermediateResult struct {
	Service       Service
	HomeTadigs    []string
	VisitorTadigs []string
	DealValue     int64
}

type Service string
