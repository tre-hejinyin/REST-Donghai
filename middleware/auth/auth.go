package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"server/middleware/constants"
	"server/middleware/response"

	"server/middleware/cache"
	"server/middleware/jwt"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  "Authorization  error",
			})
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				response.Body{
					Code: http.StatusBadRequest,
					Msg:  "Bearer  error",
				})
			return
		}
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				response.Body{
					Code: http.StatusBadRequest,
					Msg:  "token  error",
				})
			return
		}
		// from redis
		token, err := cache.HashGet(constants.CACHE_ACCOUNT_GROUP+mc.Email, constants.HASH_TOKEN_KEY)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				response.Body{
					Code: http.StatusInternalServerError,
					Msg:  err.Error(),
				})
			return
		}
		if token != parts[1] {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				response.Body{
					Code: http.StatusBadRequest,
					Msg:  "invalid token",
				})
			return
		}
		c.Set(constants.EMAIL, mc.Email)
		c.Next()
	}
}
