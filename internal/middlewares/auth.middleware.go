package middlewares

import (
	"github.com/VanThen60hz/GoShop/pkg/response"
	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrResponse(ctx, response.ErrInvalidToken, "")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
