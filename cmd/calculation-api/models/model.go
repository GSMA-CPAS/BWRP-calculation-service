package models

import "math"

const (
	Unconditional             string  = "Unconditional"
	ContractRevenue           string  = "Contract Revenue Commitment"
	DealRevenue               string  = "Deal Revenue Commitment"
	RatingPlanKindBackToFirst string  = "Threshold - back to first"
	INF                       float64 = math.MaxFloat64
)

// Result contains the results of the calculation.
type Result struct {
	Header              Header               `json:"header"`
	Total               []Deal               `json:"total"`
	IntermediateResults []IntermediateResult `json:"intermediateResults"`
}

// Deal contains the details of the deal.
type Deal struct {
	HomeTadigs        []string `json:"homeTadigs"`
	VisitorTadigs     []string `json:"visitorTadigs"`
	CommitmentReached bool     `json:"commitmentReached"`
	Value             float64  `json:"value"`
}

// Header contains the calculation versioning information
type Header struct {
	Version string `json:"version"`
	Hash    string `json:"md5hash"`
}

// IntermediateResult contains the results per service
type IntermediateResult struct {
	Service       string   `json:"service"`
	HomeTadigs    []string `json:"homeTadigs"`
	VisitorTadigs []string `json:"visitorTadigs"`
	DealValue     float64  `json:"dealValue"`
	Usage         float64  `json:"usage"`
	Type          string   `json:"type"`
}

//CalculateRequest contains the Usage data and discount models in the API request body
type CalculateRequest struct {
	Usage          Usage                    `json:"usage"`
	DiscountModels map[string]DiscountModel `json:"discounts"`
}

//Usage contains usageData records
type Usage struct {
	Inbound  []UsageData `json:"inbound"`
	Outbound []UsageData `json:"outbound"`
}

//UsageData contains usageData
type UsageData struct {
	Volume       float64 `json:"usage"`
	Unit         string  `json:"units"`
	Charge       float64 `json:"charges"`
	Tax          float64 `json:"taxes"`
	Service      string  `json:"service"`
	HomeTadig    string  `json:"homeTadig"`
	VisitorTadig string  `json:"visitorTadig"`
	Direction    string  `json:"direction"`
}

//DiscountModel contains a discount agreement
type DiscountModel struct {
	Condition     Condition      `json:"condition"`
	ServiceGroups []ServiceGroup `json:"serviceGroups"`
}

//Condition contains the discount condition
type Condition struct {
	SelectedConditionName string            `json:"kind"`
	SelectedCondition     SelectedCondition `json:"commitment"`
}

//SelectedCondition contains the parameters for the condition
type SelectedCondition struct {
	CommitmentsValue float64 `json:"value"`
	Currency         string  `json:"currency"`
	IncludingTaxes   bool    `json:"includingTaxes"`
}

//ServiceGroup contains the sergvice group data
type ServiceGroup struct {
	HomeTadigs    []string  `json:"homeTadigs"`
	VisitorTadigs []string  `json:"visitorTadigs"`
	Services      []Service `json:"services"`
}

//Service contains a chosenservice
type Service struct {
	Service              string   `json:"service"`
	IncludedInCommitment bool     `json:"includedInCommitment"`
	UsagePricing         *Pricing `json:"usagePricing"`
	AccessPricing        *Pricing `json:"accessPricing"`
}

type Pricing struct {
	Unit       string     `json:"unit"`
	RatingPlan RatingPlan `json:"ratingPlan"`
}

type RatingPlan struct {
	Kind           string `json:"kind"`
	Rate           Rate   `json:"rate"`
	BalancedRate   Rate   `json:"balancedRate"`
	UnbalancedRate Rate   `json:"unbalancedRate"`
}

type Rate struct {
	Thresholds  []Tier  `json:"thresholds"`
	FixedPrice  float64 `json:"fixedPrice"`
	LinearPrice float64 `json:"linearPrice"`
}

type Tier struct {
	Start       float64 `json:"start"`
	FixedPrice  float64 `json:"fixedPrice"`
	LinearPrice float64 `json:"linearPrice"`
}
