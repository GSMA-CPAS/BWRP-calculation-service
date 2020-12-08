package engine

//Usage specifies the usage for a single service
type Usage struct {
	Volume        int64
	Unit          Unit
	Charge        int64
	Tax           int64
	HomeTadigs    []string
	VisitorTadigs []string
}

//Unit defines the calculation units
type Unit int

type ServiceTadig struct {
	Service      Service
	HomeTadig    string
	VisitorTadig string
}

type AggregatedUsage map[ServiceTadig]Usage

// TODO take account of units and currency? (Not in MVP1)
//Aggregate gets all the usage for every service / tadig combination and then aggregates them
func (a AggregatedUsage) Aggregate(service Service, htadigs []string, vtadigs []string) (bool, Usage) {
	var volume, charge, tax int64
	for _, h := range htadigs {
		for _, v := range vtadigs {
			key := ServiceTadig{service, h, v}
			usage := a[key]
			volume = volume + usage.Volume
			charge = charge + usage.Charge
			tax = charge + usage.Tax
		}
	}
	return true, Usage{volume, 0, charge, tax, htadigs, vtadigs}
}
