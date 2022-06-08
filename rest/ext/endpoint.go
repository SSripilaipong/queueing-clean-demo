package ext

import "github.com/gin-gonic/gin"

func Endpoint[R any](ctx *gin.Context, convertRequest func(*gin.Context) (R, error), handle func(R) Response) Response {
	var req R
	var err error
	if req, err = convertRequest(ctx); err != nil {
		return Response{422, map[string]string{"message": "unable to parse body"}}
	}

	return handle(req)
}
