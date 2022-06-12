package create_visit

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"queueing-clean-demo/domain/clinical_diagnose"
	"queueing-clean-demo/domain/common/contract"
	d "queueing-clean-demo/rest/deps"
	"queueing-clean-demo/rest/ext"
)

func Route(ctx *gin.Context, deps *d.RestDeps) ext.Response {
	return ext.Endpoint(ctx, makeRequest, func(req clinical_diagnose.CreateVisit) ext.Response {

		visit, err := deps.ClinicalDiagnoseUsecase.CreateVisit(req)

		switch err.(type) {
		case common.InvalidVisitDataError:
			return ext.Response{Code: http.StatusBadRequest, Body: map[string]string{"message": "invalid visit data"}}
		case nil:
			return ext.Response{Code: http.StatusCreated, Body: map[string]string{"message": "created", "id": visit.VisitId}}
		}
		panic(err)
	})
}
