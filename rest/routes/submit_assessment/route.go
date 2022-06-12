package submit_assessment

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"queueing-clean-demo/domain/clinical_diagnose"
	"queueing-clean-demo/domain/common"
	d "queueing-clean-demo/rest/deps"
	"queueing-clean-demo/toolbox/endpoint"
)

func Route(ctx *gin.Context, deps d.IRestDeps) endpoint.Response {
	return endpoint.Endpoint(ctx, makeRequest, func(req clinical_diagnose.SubmitAssessment) endpoint.Response {
		_, err := deps.ClinicalDiagnose().SubmitAssessment(req)

		switch err.(type) {
		case common.VisitNotFoundError:
			return endpoint.Response{Code: http.StatusNotFound, Body: map[string]string{"message": "visit not found"}}
		case clinical_diagnose.AssessmentAlreadyExistError:
			return endpoint.Response{Code: http.StatusConflict, Body: map[string]string{"message": "assessment already submitted"}}
		case nil:
			return endpoint.Response{Code: http.StatusOK, Body: map[string]string{"message": "submitted"}}
		}
		panic(err)
	})
}
