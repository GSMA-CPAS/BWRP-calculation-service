package models

const (
	MOOC string = "service-0"
	SMS  string = "service-1"
)

// Result contains the results of the calculation.
type Result struct {
	IntermediateResults []IntermediateResult `json:"intermediateResults"`
}

// IntermediateResult contains the results per service
type IntermediateResult struct {
	ID            string   `json:"id"`
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
	ID           string `json:"id"`
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
	Services      []Service `json:"chosenServices"`
}

//Service contains a chosenservice
type Service struct {
	ID             string `json:"id"`
	Rate           []Tier `json:"rate"`
	BalancedRate   int64  `json:"balancedRate"`
	UnbalancedRate int64  `json:"unbalancedRate"`
}

//Tier contains a rate tier
type Tier struct {
	Threshold   int64 `json:"threshold"`
	FixedPrice  int64 `json:"fixedPrice"`
	LinearPrice int64 `json:"linearPrice"`
}
