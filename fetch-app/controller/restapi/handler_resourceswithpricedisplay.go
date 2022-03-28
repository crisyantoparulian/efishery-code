package restapi

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"fetch-app/infrastructure/log"
	"fetch-app/infrastructure/util"
	"fetch-app/usecase/fetchresourcewithprice"
)

// resourcesWithPriceDisplayHandler godoc
// @Summary Resources With Price USD
// @Tags Resources
// @Description This api is for display resources with price USD
// @Accept  json
// @Produce  json
// @Security Bearer
// @Success 200 {object} Response{success=bool,errorCode=string,errorMessage=string,traceId=string,data=fetchresourcewithprice.InportResponse.Resources}
// @Router /api/v1/display/resource-with-price [get]
func (r *Controller) resourcesWithPriceDisplayHandler(inputPort fetchresourcewithprice.Inport) gin.HandlerFunc {

	return func(c *gin.Context) {

		traceID := util.GenerateID()

		ctx := log.Context(c.Request.Context(), traceID)

		var req fetchresourcewithprice.InportRequest

		res, err := inputPort.Execute(ctx, req)
		if err != nil {
			log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, NewErrorResponse(err, traceID))
			return
		}

		// log.Info(ctx, util.MustJSON(res))
		c.JSON(http.StatusOK, NewSuccessResponse(res.Resources, traceID))

	}
}
