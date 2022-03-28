package fetchresourceaggregation

import (
	"context"
	"fetch-app/infrastructure/util"
	"fmt"
	"log"
	"strconv"
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

		if !obj.IsValidData() {
			continue
		}

		date, err := util.ForceParseDate(*obj.TglParsed)
		if err != nil {
			log.Println("Error ForceParseDate", err)
		}

		_, week := date.ISOWeek()
		key := *obj.AreaProvinsi + "_" + fmt.Sprint(week)

		price, _ := strconv.Atoi(*obj.Price)
		size, _ := strconv.Atoi(*obj.Size)

		if aggregates[key] == nil {
			aggregates[key] = &AggregateResources{
				AreaProvinsi: *obj.AreaProvinsi,
				Week:         week,
				Price:        []int{price},
				Size:         []int{size},
			}
		} else {
			aggregates[key].Price = append(aggregates[key].Price, price)
			aggregates[key].Size = append(aggregates[key].Size, size)
		}
	}

	for k, v := range aggregates {

		min, max, med, avg := util.CalcMinMaxMedAvg(v.Price)

		aggregates[k].AggregatePrice = Aggregate{
			Min: min,
			Max: max,
			Avg: avg,
			Med: med,
		}

		min, max, med, avg = util.CalcMinMaxMedAvg(v.Size)

		aggregates[k].AggregateSize = Aggregate{
			Min: min,
			Max: max,
			Avg: avg,
			Med: med,
		}

		res.Result = append(res.Result, v)
	}

	return res, nil
}
