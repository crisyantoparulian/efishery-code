package registry

import (
	"fetch-app/application"
	"fetch-app/controller"
	"fetch-app/controller/restapi"
	"fetch-app/gateway/prod"
	"fetch-app/infrastructure/config"
	"fetch-app/infrastructure/server"
	"fetch-app/usecase/fetchresourcewithprice"
)

type fetchapp struct {
	*server.GinHTTPHandler
	controller.Controller
}

func NewFetchapp() func() application.RegistryContract {
	return func() application.RegistryContract {

		cfg, err := config.ReadConfig()
		if err != nil {
			panic(err)
		}

		httpHandler := server.NewGinHTTPHandlerDefault(cfg.Host)

		datasource := prod.NewGateway(cfg)

		return &fetchapp{
			GinHTTPHandler: &httpHandler,
			Controller: &restapi.Controller{
				Router:                       httpHandler.Router,
				FetchResourceWithPriceInport: fetchresourcewithprice.NewUsecase(datasource),
			},
		}

	}
}
