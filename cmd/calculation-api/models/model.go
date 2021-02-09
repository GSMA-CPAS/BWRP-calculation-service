package models

const (
	MOOC string = "service-0"
	SMS  string = "service-1"
)

//FixedPrice defines the fixedPrice
type FixedPrice int64

//LinearPrice defines the linearPrice
type LinearPrice int64

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

type CalculateRequest struct {
	Usage     []UsageData     `json:"usage"`
	Discounts []DiscountModel `json:"discounts"`
}

//Usage contains usageData records
type Usage struct {
	Usage []UsageData `json:"usage"`
}

//UsageData contains usageData
type UsageData struct {
	Volume       int64  `json:"volume"`
	Unit         int    `json:"unit"`
	Charge       int64  `json:"charge"`
	Tax          int64  `json:"tax"`
	Service      string `json:"service"`
	HomeTadig    string `json:"homeTadig"`
	VisitorTadig string `json:"visitorTadig"`
}

//Contract contains contract discount models (user and partner)
type Contract struct {
	Discounts []DiscountModel `json:"discounts"`
}

//DiscountModels contain the discount agreement
type DiscountModel struct {
	Condition     Condition      `json:"condition"`
	ServiceGroups []ServiceGroup `json:"serviceGroups"`
}

//Condition contains the discount condition
type Condition struct {
	SelectedConditionName string `json:"selectedConditionName"`
	SelectedCondition     string `json:"selectedCondition"`
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
	Rate           Rate         `json:"rate"`
	BalancedRate   BalancedRate `json:"balancedRate"`
	UnbalancedRate BalancedRate `json:"unbalancedRate"`
}

type BalancedRate struct {
	Unit  int64 `json:"unit"`
	Value int64 `json:"value"`
}

type Rate struct {
	Thresholds  []Tier `json:"thresholds"`
	FixedPrice  int64  `json:"fixedPrice"`
	LinearPrice int64  `json:"linearPrice"`
}

type Tier struct {
	Start       int64 `json:"start"`
	FixedPrice  int64 `json:"fixedPrice"`
	LinearPrice int64 `json:"linearPrice"`
}
