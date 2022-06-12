package check_visits

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"queueing-clean-demo/domain/manage_doctor_queue"
	d "queueing-clean-demo/rest/deps"
	"queueing-clean-demo/toolbox/endpoint"
)

func Route(ctx *gin.Context, deps *d.RestDeps) endpoint.Response {
	return endpoint.Endpoint(ctx, makeRequest, func(req manage_doctor_queue.CheckVisits) endpoint.Response {

		queue, err := deps.ManageDoctorQueueUsecase.CheckVisits(req)

		switch err.(type) {
		case manage_doctor_queue.DoctorQueueNotFoundError:
			return endpoint.Response{Code: http.StatusNotFound, Body: gin.H{"message": "doctor queue not found"}}
		case nil:
			return endpoint.Response{Code: http.StatusOK, Body: queue}
		}
		panic(err)
	})
}
