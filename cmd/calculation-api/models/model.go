package models

const (
	Unconditional   string = "Unconditional"
	ContractRevenue string = "Contract Revenue Commitment"
	DealRevenue     string = "Deal Revenue Commitment"
)

// Result contains the results of the calculation.
type Result struct {
	IntermediateResults []IntermediateResult `json:"intermediateResults"`
}

// IntermediateResult contains the results per service
type IntermediateResult struct {
	Service       string   `json:"service"`
	HomeTadigs    []string `json:"homeTadigs"`
	VisitorTadigs []string `json:"visitorTadigs"`
	DealValue     string   `json:"dealValue"`
}

//CalculateRequest contains the Usage data and discount models in the API request body
type CalculateRequest struct {
	Usage          []UsageData              `json:"usage"`
	DiscountModels map[string]DiscountModel `json:"discounts"`
}

//Usage contains usageData records
type Usage struct {
	Usage []UsageData `json:"usage"`
}

//UsageData contains usageData
type UsageData struct {
	Volume       string `json:"usage"`
	Unit         string `json:"units"`
	Charge       string `json:"charges"`
	Tax          string `json:"taxes"`
	Service      string `json:"service"`
	HomeTadig    string `json:"homeTadig"`
	VisitorTadig string `json:"visitorTadig"`
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
	CommitmentsValue string `json:"value"`
	Currency         string `json:"currency"`
	IncludingTaxes   bool   `json:"includingTaxes"`
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
	Unit       int        `json:"unit"`
	RatingPlan RatingPlan `json:"ratingPlan"`
}

type RatingPlan struct {
	Kind           string `json:"kind"`
	Rate           Rate   `json:"rate"`
	BalancedRate   Rate   `json:"balancedRate"`
	UnbalancedRate Rate   `json:"unbalancedRate"`
}

type Rate struct {
	Thresholds  []Tier `json:"thresholds"`
	FixedPrice  string `json:"fixedPrice"`
	LinearPrice string `json:"linearPrice"`
}

type Tier struct {
	Start       string `json:"start"`
	FixedPrice  string `json:"fixedPrice"`
	LinearPrice string `json:"linearPrice"`
}
