package create_doctor_queue

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"queueing-clean-demo/domain/usecase/manage_doctor_queue/contract"
	d "queueing-clean-demo/rest/deps"
	"queueing-clean-demo/rest/ext"
)

func Route(ctx *gin.Context, deps *d.RestDeps) ext.Response {
	return ext.Endpoint(ctx, makeRequest, func(req contract.CreateDoctorQueue) ext.Response {

		_, err := deps.ManageDoctorQueueUsecase.CreateDoctorQueue(req)

		switch err.(type) {
		case contract.DuplicateDoctorQueueIdError:
			return ext.Response{Code: http.StatusConflict, Body: map[string]string{"message": "doctor queue already exists"}}
		case nil:
			return ext.Response{Code: http.StatusCreated, Body: map[string]string{"message": "created"}}
		}
		panic(err)
	})
}
