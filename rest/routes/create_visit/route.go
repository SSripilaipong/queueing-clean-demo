package create_visit

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"queueing-clean-demo/domain/clinical_diagnose"
	"queueing-clean-demo/domain/common"
	d "queueing-clean-demo/rest/deps"
	"queueing-clean-demo/toolbox/endpoint"
)

func Route(ctx *gin.Context, deps *d.RestDeps) endpoint.Response {
	return endpoint.Endpoint(ctx, makeRequest, func(req clinical_diagnose.CreateVisit) endpoint.Response {

		visit, err := deps.ClinicalDiagnoseUsecase.CreateVisit(req)

		switch err.(type) {
		case common.InvalidVisitDataError:
			return endpoint.Response{Code: http.StatusBadRequest, Body: map[string]string{"message": "invalid visit data"}}
		case nil:
			return endpoint.Response{Code: http.StatusCreated, Body: map[string]string{"message": "created", "id": visit.VisitId}}
		}
		panic(err)
	})
}
