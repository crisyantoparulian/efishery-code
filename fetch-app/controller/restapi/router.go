package restapi

import (
	"fetch-app/usecase/fetchresourcewithprice"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Router                       gin.IRouter
	FetchResourceWithPriceInport fetchresourcewithprice.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
	v1 := r.Router.Group("/v1")

	v1.GET("/display-with-price", r.authorized(), r.resourcesWithPriceDisplayHandler(r.FetchResourceWithPriceInport))
}
