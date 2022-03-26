package fetchresourcewithprice

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
	Resources []*ResourceWithPrice
}

type ResourceWithPrice struct {
	Uuid      *string `json:"uuid"`
	Komoditas *string `json:"komoditas"`
	Area      *Area   `json:"area"`
	Size      *string `json:"size"`
	Price     *Price  `json:"price"`
	TglParsed *string `json:"tgl_parsed"`
	Timestamp *string `json:"timestamp"`
}

type Area struct {
	Provinsi *string `json:"provinsi"`
	Kota     *string `json:"kota"`
}

type Price struct {
	Idr *string `json:"idr"`
	Usd *string `json:"usd"`
}
