package fetchresourceaggregation

import (
	"context"
	"fetch-app/infrastructure/util"
	"log"
	"sort"
)

//go:generate mockery --name Outport -output mocks/

type fetchresourceaggregationInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &fetchresourceaggregationInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *fetchresourceaggregationInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// code your usecase definition here ...

	resourcesObjs, err := r.outport.FindResources(ctx)
	if err != nil {
		return nil, err
	}

	aggregates := map[string]*AggregateResources{}
	for _, obj := range resourcesObjs {
		if obj.AreaProvinsi != nil && obj.TglParsed != nil {
			date, err := util.ForceParseDate(*obj.TglParsed)
			if err != nil {
				log.Println("Error", err)
				return res, nil
			}
			_, week := date.ISOWeek()

			if aggregates[*obj.AreaProvinsi] == nil {
				aggregates[*obj.AreaProvinsi] = &AggregateResources{
					AreaProvinsi: *obj.AreaProvinsi,
					Weeks:        []int{week},
				}
			} else {
				aggregates[*obj.AreaProvinsi].Weeks = append(aggregates[*obj.AreaProvinsi].Weeks, week)
			}
		}
	}

	for k, v := range aggregates {
		sort.Ints(v.Weeks)

		min, max := 0, 0
		lenWeeks := len(v.Weeks)
		if lenWeeks > 0 {
			min = v.Weeks[0]
			max = v.Weeks[lenWeeks-1]
		}

		aggregates[k].Min = min
		aggregates[k].Max = max
		aggregates[k].Med = util.CalcMedian(v.Weeks)
		aggregates[k].Avg = util.CalcAvg(v.Weeks)

		res.Result = append(res.Result, v)
	}

	return res, nil
}
