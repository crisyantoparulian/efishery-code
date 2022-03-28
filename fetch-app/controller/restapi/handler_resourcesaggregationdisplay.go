package restapi

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"fetch-app/infrastructure/log"
	"fetch-app/infrastructure/util"
	"fetch-app/usecase/fetchresourceaggregation"
)

// resourcesAggregationDisplayHandler godoc
// @Summary Resources Aggregation
// @Tags Resources
// @Description This api is for display aggregation of resources
// @Accept  json
// @Produce  json
// @Security Bearer
// @Success 200 {object} Response{success=bool,errorCode=string,errorMessage=string,traceId=string,data=fetchresourceaggregation.InportResponse.Result}
// @Router /api/v1/display/resource-aggregation [get]
func (r *Controller) resourcesAggregationDisplayHandler(inputPort fetchresourceaggregation.Inport) gin.HandlerFunc {

	return func(c *gin.Context) {

		traceID := util.GenerateID()

		ctx := log.Context(c.Request.Context(), traceID)

		var req fetchresourceaggregation.InportRequest

		res, err := inputPort.Execute(ctx, req)
		if err != nil {
			log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, NewErrorResponse(err, traceID))
			return
		}

		// log.Info(ctx, util.MustJSON(res))
		c.JSON(http.StatusOK, NewSuccessResponse(res.Result, traceID))

	}
}
