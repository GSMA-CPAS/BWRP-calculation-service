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
	DealValue     float32  `json:"dealValue"`
}

//CalculateRequest contains the Usage data and discount models in the API request body
type CalculateRequest struct {
	Usage          []UsageData              `json:"usage"`
	DiscountModels map[string]DiscountModel `json:"discountModels"`
}

//Usage contains usageData records
type Usage struct {
	Usage []UsageData `json:"usage"`
}

//UsageData contains usageData
type UsageData struct {
	Volume       float32 `json:"volume"`
	Unit         int     `json:"unit"`
	Charge       float32 `json:"charge"`
	Tax          float32 `json:"tax"`
	Service      string  `json:"service"`
	HomeTadig    string  `json:"homeTadig"`
	VisitorTadig string  `json:"visitorTadig"`
}

//DiscountModel contains a discount agreement
type DiscountModel struct {
	Condition     Condition      `json:"condition"`
	ServiceGroups []ServiceGroup `json:"serviceGroups"`
}

//Condition contains the discount condition
type Condition struct {
	SelectedConditionName string            `json:"selectedConditionName"`
	SelectedCondition     SelectedCondition `json:"selectedCondition"`
}

//SelectedCondition contains the parameters for the condition
type SelectedCondition struct {
	CommitmentsValue int64  `json:"commitmentValue"`
	Currency         string `json:"currency"`
	IncludingTaxes   bool   `json:"includingTaxes"`
}

//ServiceGroup contains the sergvice group data
type ServiceGroup struct {
	HomeTadigs    Codes     `json:"homeTadigs"`
	VisitorTadigs Codes     `json:"visitorTadigs"`
	Services      []Service `json:"chosenServices"`
}

type Codes struct {
	Codes []string `json:"codes"`
}

//Service contains a chosenservice
type Service struct {
	ID                   string `json:"id"`
	Name                 string `json:"name"`
	IncludedInCommitment bool   `json:"includedInCommitment"`
	Unit                 string `json:"unit"`
	Rate                 []Tier `json:"rate"`
	AccessPricingUnit    string `json:"accessPricingUnit"`
	AccessPricingRate    []Tier `json:"accessPricingRate"`
	BalancedRate         []Tier `json:"balancedRate"`
	UnbalancedRate       []Tier `json:"unbalancedRate"`
}

// type BalancedRate struct {
// 	Unit  int64 `json:"unit"`
// 	Value int64 `json:"value"`
// }

type Rate struct {
	Thresholds  []Tier
	FixedPrice  int64 `json:"fixedPrice"`
	LinearPrice int64 `json:"linearPrice"`
}

type Tier struct {
	Threshold   int64 `json:"threshold"`
	FixedPrice  int64 `json:"fixedPrice"`
	LinearPrice int64 `json:"linearPrice"`
}
