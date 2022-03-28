package restapi

import (
	"fetch-app/infrastructure/config"
	"fetch-app/usecase/fetchresourceaggregation"
	"fetch-app/usecase/fetchresourcewithprice"

	_ "fetch-app/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type Controller struct {
	Config                       *config.Config
	Router                       gin.IRouter
	FetchResourceWithPriceInport fetchresourcewithprice.Inport
	FetchResourceAggregation     fetchresourceaggregation.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
	v1 := r.Router.Group("/v1")

	v1.GET("/display/with-price", r.authorized(), r.isRoleAdmin(), r.resourcesWithPriceDisplayHandler(r.FetchResourceWithPriceInport))
	v1.GET("/display/aggregation", r.authorized(), r.isRoleAdmin(), r.resourcesAggregationDisplayHandler(r.FetchResourceAggregation))

	r.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
