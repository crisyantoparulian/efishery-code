package prod

import (
	"fetch-app/infrastructure/config"
	"time"

	"github.com/go-resty/resty/v2"
)

type gateway struct {
	*EfisheryResourcesImpl
	*CurrencyImpl
}

// NewGateway ...
func NewGateway(cfg *config.Config) *gateway {
	return &gateway{
		EfisheryResourcesImpl: &EfisheryResourcesImpl{Config: cfg},
		CurrencyImpl:          &CurrencyImpl{Config: cfg},
	}
}

func newReqGeneralResty() *resty.Request {
	client := resty.New()

	req := client.SetTimeout(10*time.Second).R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	return req
}
