package fetchresourceaggregation

import (
	"context"
	"fmt"
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

	resourcesObjs, err := r.outport.FindResources(ctx, "resourcesID")
	if err != nil {
		return nil, err
	}

	for _, obj := range resourcesObjs {
		fmt.Printf("%v\n", obj)
	}

	return res, nil
}
