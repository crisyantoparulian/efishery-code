package restapi

import (
	"fetch-app/infrastructure/config"
	"fetch-app/usecase/displayoneuser"
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
	DisplayOneUser               displayoneuser.Inport
}

// @title Fetch apps
// @version 1.0.0
// @description This API is for fetch & mapping resources

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

//@securityDefinitions.apikey Bearer
//@in header
//@name Authorization

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {

	v1 := r.Router.Group("/api/v1")

	v1.GET("/display/resource-with-price", r.authorized(), r.isRoleAdmin(), r.resourcesWithPriceDisplayHandler(r.FetchResourceWithPriceInport))
	v1.GET("/display/resource-aggregation", r.authorized(), r.isRoleAdmin(), r.resourcesAggregationDisplayHandler(r.FetchResourceAggregation))

	v1.GET("/user", r.authorized(), r.displayOneUserHandler(r.DisplayOneUser))

	r.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
