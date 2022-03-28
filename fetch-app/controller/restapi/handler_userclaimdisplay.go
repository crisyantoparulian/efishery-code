package restapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"

	"fetch-app/application/apperror"
	"fetch-app/infrastructure/log"
	"fetch-app/infrastructure/util"
	"fetch-app/usecase/displayoneuser"
)

// userClaimDisplayHandler godoc
// @Summary Display user claim
// @Tags Users
// @Description This api is for display resources with price USD
// @Accept  json
// @Produce  json
// @Security Bearer
// @Success 200 {object} Response{success=bool,errorCode=string,errorMessage=string,traceId=string,data=displayoneuser.InportResponse.User}
// @Router /api/v1/user [get]
func (r *Controller) displayOneUserHandler(inputPort displayoneuser.Inport) gin.HandlerFunc {

	return func(c *gin.Context) {

		traceID := util.GenerateID()

		ctx := log.Context(c.Request.Context(), traceID)

		var req displayoneuser.InportRequest

		userInfo, ok := c.MustGet("userInfo").(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, NewErrorResponse(apperror.UserClaimNotFoundError, ""))
			return
		}

		req.Claims = userInfo
		res, err := inputPort.Execute(ctx, req)
		if err != nil {
			log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, NewErrorResponse(err, traceID))
			return
		}

		// log.Info(ctx, util.MustJSON(res))
		c.JSON(http.StatusOK, NewSuccessResponse(res.User, traceID))

	}
}
