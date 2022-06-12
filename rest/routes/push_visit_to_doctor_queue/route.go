package push_visit_to_doctor_queue

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"queueing-clean-demo/domain/manage_doctor_queue"
	d "queueing-clean-demo/rest/deps"
	"queueing-clean-demo/toolbox/endpoint"
)

func Route(ctx *gin.Context, deps *d.RestDeps) endpoint.Response {
	return endpoint.Endpoint(ctx, makeRequest, func(req manage_doctor_queue.PushVisitToDoctorQueue) endpoint.Response {
		queue, err := deps.ManageDoctorQueueUsecase.PushVisit(req)

		switch err.(type) {
		case manage_doctor_queue.DoctorQueueNotFoundError:
			return endpoint.Response{Code: http.StatusNotFound, Body: gin.H{"message": "doctor queue not found"}}
		case manage_doctor_queue.VisitAlreadyExistsError:
			return endpoint.Response{Code: http.StatusConflict, Body: gin.H{"message": "visit already exists in the doctor queue"}}
		case nil:
			return endpoint.Response{Code: http.StatusOK, Body: queue}
		}
		panic(err)
	})
}
