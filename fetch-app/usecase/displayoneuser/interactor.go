package displayoneuser

import (
	"context"
)

//go:generate mockery --name Outport -output mocks/

type displayoneuserInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &displayoneuserInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *displayoneuserInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	res.User = req.Claims

	return res, nil
}
