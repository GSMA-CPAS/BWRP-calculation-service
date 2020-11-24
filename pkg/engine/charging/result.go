/*
 License-Placeholder
*/

package engine

type Result struct {
	IntermediateResults []IntermediateResult
}

type IntermediateResult struct {
	Service   Service
	Tadigs    []string
	DealValue int64
}

type Service string
