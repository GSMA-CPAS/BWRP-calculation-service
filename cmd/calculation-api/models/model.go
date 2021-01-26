package models

// Result contains the results of the calculation.
type Result struct {
	IntermediateResults []IntermediateResult `json:"intermediateResults"`
}

// IntermediateResult contains the results per service
type IntermediateResult struct {
	Service       Service  `json:"service"`
	HomeTadigs    []string `json:"homeTadigs"`
	VisitorTadigs []string `json:"visitorTadigs"`
	DealValue     int64    `json:"dealValue"`
}

// Service is a string
type Service string
