package push_visit_to_doctor_queue

import (
	"github.com/gin-gonic/gin"
	"net/http"
	manage_doctor_queue2 "queueing-clean-demo/domain/manage_doctor_queue"
	d "queueing-clean-demo/rest/deps"
	"queueing-clean-demo/rest/ext"
)

func Route(ctx *gin.Context, deps *d.RestDeps) ext.Response {
	return ext.Endpoint(ctx, makeRequest, func(req manage_doctor_queue2.PushVisitToDoctorQueue) ext.Response {
		queue, err := deps.ManageDoctorQueueUsecase.PushVisit(req)

		switch err.(type) {
		case manage_doctor_queue2.DoctorQueueNotFoundError:
			return ext.Response{Code: http.StatusNotFound, Body: gin.H{"message": "doctor queue not found"}}
		case manage_doctor_queue2.VisitAlreadyExistsError:
			return ext.Response{Code: http.StatusConflict, Body: gin.H{"message": "visit already exists in the doctor queue"}}
		case nil:
			return ext.Response{Code: http.StatusOK, Body: queue}
		}
		panic(err)
	})
}
