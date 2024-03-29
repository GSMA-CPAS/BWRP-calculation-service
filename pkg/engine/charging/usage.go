package engine

//Usage specifies the usage for a single service
type Usage struct {
	Volume        float64
	Unit          Unit
	Charge        float64
	Tax           float64
	HomeTadigs    []string
	VisitorTadigs []string
	Direction     string
}

//Unit defines the calculation units
type Unit string

type ServiceTadig struct {
	Service      string
	HomeTadig    string
	VisitorTadig string
}

type AggregatedUsage map[ServiceTadig]Usage

//Aggregate gets all the usage for every service / tadig combination and then aggregates them
func (a AggregatedUsage) Aggregate(service string, htadigs []string, vtadigs []string) Usage {
	var volume, charge, tax float64
	var direction string
	for _, h := range htadigs {
		for _, v := range vtadigs {
			key := ServiceTadig{service, h, v}
			usage := a[key]
			volume = volume + usage.Volume
			charge = charge + usage.Charge
			tax = tax + usage.Tax
			if len(usage.Direction) > 1 {
				direction = usage.Direction
			}

		}
	}
	return Usage{volume, "", charge, tax, htadigs, vtadigs, direction}
}
