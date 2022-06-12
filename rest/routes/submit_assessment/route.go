package submit_assessment

import (
	"github.com/gin-gonic/gin"
	"net/http"
	clinical_diagnose2 "queueing-clean-demo/domain/clinical_diagnose"
	"queueing-clean-demo/domain/common"
	d "queueing-clean-demo/rest/deps"
	"queueing-clean-demo/rest/ext"
)

func Route(ctx *gin.Context, deps *d.RestDeps) ext.Response {
	return ext.Endpoint(ctx, makeRequest, func(req clinical_diagnose2.SubmitAssessment) ext.Response {
		_, err := deps.ClinicalDiagnoseUsecase.SubmitAssessment(req)

		switch err.(type) {
		case common.VisitNotFoundError:
			return ext.Response{Code: http.StatusNotFound, Body: map[string]string{"message": "visit not found"}}
		case clinical_diagnose2.AssessmentAlreadyExistError:
			return ext.Response{Code: http.StatusConflict, Body: map[string]string{"message": "assessment already submitted"}}
		case nil:
			return ext.Response{Code: http.StatusOK, Body: map[string]string{"message": "submitted"}}
		}
		panic(err)
	})
}
