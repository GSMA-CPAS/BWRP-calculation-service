package engine

//Usage specifies the usage for a single service
type Usage struct {
	Volume int64
	Unit   Unit
	Tadigs []string
}

//Unit defines the calculation units
type Unit int

type ServiceTadig struct {
	Service Service
	Tadig   string
}

type AggregatedUsage map[ServiceTadig]Usage

// TODO take account of units
//Aggregate gets all the usage for every service / tadig combination and then aggregates them
func (a AggregatedUsage) Aggregate(service Service, tadigs []string) (bool, Usage) {
	volume := int64(0)
	for _, tadig := range tadigs {
		key := ServiceTadig{service, tadig}
		usage := a[key]
		volume = volume + usage.Volume
	}
	return true, Usage{volume, 0, tadigs}
}
