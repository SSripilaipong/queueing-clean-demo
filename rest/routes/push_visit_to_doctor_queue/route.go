package push_visit_to_doctor_queue

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"queueing-clean-demo/domain/usecase/manage_doctor_queue/contract"
	d "queueing-clean-demo/rest/deps"
	"queueing-clean-demo/rest/ext"
)

func Route(ctx *gin.Context, deps *d.RestDeps) ext.Response {
	return ext.Endpoint(ctx, makeRequest, func(req contract.PushVisitToDoctorQueue) ext.Response {
		queue, err := deps.ManageDoctorQueueUsecase.PushVisit(req)

		switch err.(type) {
		case contract.DoctorQueueNotFoundError:
			return ext.Response{Code: http.StatusNotFound, Body: gin.H{"message": "doctor queue not found"}}
		case contract.VisitAlreadyExistsError:
			return ext.Response{Code: http.StatusConflict, Body: gin.H{"message": "visit already exists in the doctor queue"}}
		case nil:
			return ext.Response{Code: http.StatusOK, Body: queue}
		}
		panic(err)
	})
}
