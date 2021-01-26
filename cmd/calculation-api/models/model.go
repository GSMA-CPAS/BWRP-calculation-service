package models

// Result contains the results of the calculation.
type Result struct {
	IntermediateResults []IntermediateResult `json:"intermediateResults"`
}

// IntermediateResult contains the results per service
type IntermediateResult struct {
	Service       string   `json:"service"`
	HomeTadigs    []string `json:"homeTadigs"`
	VisitorTadigs []string `json:"visitorTadigs"`
	DealValue     int64    `json:"dealValue"`
}

// Usage contains a usage record
type Usage struct {
	Volume       int64
	Unit         int
	Charge       int64
	Tax          int64
	Service      string
	HomeTadig    string
	VisitorTadig string
}
