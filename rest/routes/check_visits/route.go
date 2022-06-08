package check_visits

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"queueing-clean-demo/domain/usecase/manage_doctor_queue/contract"
	d "queueing-clean-demo/rest/deps"
	"queueing-clean-demo/rest/ext"
)

func Route(ctx *gin.Context, deps *d.RestDeps) ext.Response {
	return ext.Endpoint(ctx, makeRequest, func(req contract.CheckVisits) ext.Response {

		queue, err := deps.ManageDoctorQueueUsecase.CheckVisits(req)

		switch err.(type) {
		case contract.DoctorQueueNotFoundError:
			return ext.Response{Code: http.StatusNotFound, Body: gin.H{"message": "doctor queue not found"}}
		case nil:
			return ext.Response{Code: http.StatusOK, Body: queue}
		}
		panic(err)
	})
}
