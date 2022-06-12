package endpoint

import (
	"github.com/gin-gonic/gin"
)

func R[D any](route func(ctx *gin.Context, deps D) Response, deps D) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		response := route(ctx, deps)
		ctx.AbortWithStatusJSON(response.Code, response.Body)
	}
}
