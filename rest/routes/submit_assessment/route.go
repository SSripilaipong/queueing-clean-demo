package submit_assessment

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"queueing-clean-demo/domain/clinical_diagnose/contract"
	"queueing-clean-demo/domain/common/contract"
	d "queueing-clean-demo/rest/deps"
	"queueing-clean-demo/rest/ext"
)

func Route(ctx *gin.Context, deps *d.RestDeps) ext.Response {
	return ext.Endpoint(ctx, makeRequest, func(req clinical_diagnose.SubmitAssessment) ext.Response {
		_, err := deps.ClinicalDiagnoseUsecase.SubmitAssessment(req)

		switch err.(type) {
		case common.VisitNotFoundError:
			return ext.Response{Code: http.StatusNotFound, Body: map[string]string{"message": "visit not found"}}
		case clinical_diagnose.AssessmentAlreadyExistError:
			return ext.Response{Code: http.StatusConflict, Body: map[string]string{"message": "assessment already submitted"}}
		case nil:
			return ext.Response{Code: http.StatusOK, Body: map[string]string{"message": "submitted"}}
		}
		panic(err)
	})
}
