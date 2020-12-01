package engine

//Usage specifies the usage for a single service
type Usage struct {
	Volume int64
	Unit   Unit
	Charge int64
	Tax    int64
	Tadigs []string
}

//Unit defines the calculation units
type Unit int

type ServiceTadig struct {
	Service Service
	Tadig   string
}

type AggregatedUsage map[ServiceTadig]Usage

// TODO take account of units and currecy?
//Aggregate gets all the usage for every service / tadig combination and then aggregates them
func (a AggregatedUsage) Aggregate(service Service, tadigs []string) (bool, Usage) {
	var volume, charge, tax int64
	for _, tadig := range tadigs {
		key := ServiceTadig{service, tadig}
		usage := a[key]
		volume = volume + usage.Volume
		charge = charge + usage.Charge
		tax = charge + usage.Tax

	}
	return true, Usage{volume, 0, charge, tax, tadigs}
}
