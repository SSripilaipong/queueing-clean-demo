package ext

import (
	"github.com/gin-gonic/gin"
	"queueing-clean-demo/rest/deps"
)

func R(route func(ctx *gin.Context, deps *deps.RestDeps) Response, deps *deps.RestDeps) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		response := route(ctx, deps)
		ctx.AbortWithStatusJSON(response.Code, response.Body)
	}
}
