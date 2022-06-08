package submit_assessment

import (
	"github.com/gin-gonic/gin"
	"net/http"
	e "queueing-clean-demo/domain/contract"
	"queueing-clean-demo/domain/usecase/clinical_diagnose/contract"
	d "queueing-clean-demo/rest/deps"
	"queueing-clean-demo/rest/ext"
)

func Route(ctx *gin.Context, deps *d.RestDeps) ext.Response {
	return ext.Endpoint(ctx, makeRequest, func(req contract.SubmitAssessment) ext.Response {
		_, err := deps.ClinicalDiagnoseUsecase.SubmitAssessment(req)

		switch err.(type) {
		case e.VisitNotFoundError:
			return ext.Response{Code: http.StatusNotFound, Body: map[string]string{"message": "visit not found"}}
		case contract.AssessmentAlreadyExistError:
			return ext.Response{Code: http.StatusConflict, Body: map[string]string{"message": "assessment already submitted"}}
		case nil:
			return ext.Response{Code: http.StatusOK, Body: map[string]string{"message": "submitted"}}
		}
		panic(err)
	})
}
