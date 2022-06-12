package create_doctor_queue

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"queueing-clean-demo/domain/manage_doctor_queue"
	d "queueing-clean-demo/rest/deps"
	"queueing-clean-demo/toolbox/endpoint"
)

func Route(ctx *gin.Context, deps *d.RestDeps) endpoint.Response {
	return endpoint.Endpoint(ctx, makeRequest, func(req manage_doctor_queue.CreateDoctorQueue) endpoint.Response {

		_, err := deps.ManageDoctorQueueUsecase.CreateDoctorQueue(req)

		switch err.(type) {
		case manage_doctor_queue.DuplicateDoctorQueueIdError:
			return endpoint.Response{Code: http.StatusConflict, Body: map[string]string{"message": "doctor queue already exists"}}
		case nil:
			return endpoint.Response{Code: http.StatusCreated, Body: map[string]string{"message": "created"}}
		}
		panic(err)
	})
}
