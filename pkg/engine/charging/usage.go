package engine

//Usage specifies the usage for a single service
type Usage struct {
	Volume        float64
	Unit          Unit
	Charge        float64
	Tax           float64
	HomeTadigs    []string
	VisitorTadigs []string
}

//Unit defines the calculation units
type Unit string

type ServiceTadig struct {
	Service      string
	HomeTadig    string
	VisitorTadig string
}

type AggregatedUsage map[ServiceTadig]Usage

// TODO take account of units and currency? (Not in MVP1)
//Aggregate gets all the usage for every service / tadig combination and then aggregates them
func (a AggregatedUsage) Aggregate(service string, htadigs []string, vtadigs []string) Usage {
	var volume, charge, tax float64
	for _, h := range htadigs {
		for _, v := range vtadigs {
			key := ServiceTadig{service, h, v}
			usage := a[key]
			volume = volume + usage.Volume
			charge = charge + usage.Charge
			tax = charge + usage.Tax
		}
	}
	return Usage{volume, "", charge, tax, htadigs, vtadigs}
}
