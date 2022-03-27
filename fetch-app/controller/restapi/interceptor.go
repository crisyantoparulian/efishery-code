package restapi

import (
	"fetch-app/application/apperror"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

type userInfo struct {
	Name      string
	Phone     string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// authorized is an interceptor
func (r *Controller) authorized() gin.HandlerFunc {

	return func(c *gin.Context) {

		authorizationHeader := c.Request.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			c.AbortWithStatusJSON(http.StatusForbidden, NewErrorResponse(apperror.MissingJwtTokenError, ""))
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, apperror.SigningMethodJwtInvalidError
			} else if method != jwt.SigningMethodHS256 {
				return nil, apperror.SigningMethodJwtInvalidError
			}

			return []byte(r.Config.Jwt.Secret), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, NewErrorResponse(err, ""))
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusForbidden, NewErrorResponse(err, ""))
			return
		}

		c.Set("userInfo", claims)
		c.Next()
	}
}

// authorized is an interceptor
func (r *Controller) isRoleAdmin() gin.HandlerFunc {

	return func(c *gin.Context) {
		userInfo, ok := c.MustGet("userInfo").(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, NewErrorResponse(apperror.UserClaimNotFoundError, ""))
			return
		}

		if role, ok := userInfo["role"]; ok {
			if role == "admin" {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, NewErrorResponse(apperror.UnauthorizeRoleAccesError, ""))
		return
	}
}
