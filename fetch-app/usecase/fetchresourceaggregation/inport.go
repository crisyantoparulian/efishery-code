package fetchresourceaggregation

import (
	"context"
)

// Inport of Usecase
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase
type InportRequest struct {
}

// InportResponse is response payload after running the usecase
type InportResponse struct {
	Result []*AggregateResources
}

type AggregateResources struct {
	AreaProvinsi string  `json:"area_provinsi"`
	Weeks        []int   `json:"weeks"`
	Min          int     `json:"min"`
	Max          int     `json:"max"`
	Avg          float64 `json:"avg"`
	Med          float64 `json:"med"`
}
