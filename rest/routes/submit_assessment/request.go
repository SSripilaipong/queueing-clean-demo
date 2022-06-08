package submit_assessment

import (
	"github.com/gin-gonic/gin"
	"queueing-clean-demo/domain/usecase/clinical_diagnose/contract"
)

type SubmitAssessment struct {
	NursingAssessment string `json:"comment"`
	PainScore         int    `json:"painScore"`
}

func makeRequest(ctx *gin.Context) (contract.SubmitAssessment, error) {
	visitId := ctx.Param("visitId")

	var body SubmitAssessment
	if err := ctx.ShouldBindJSON(&body); err != nil {
		return contract.SubmitAssessment{}, err
	}

	return contract.SubmitAssessment{
		VisitId:           visitId,
		NursingAssessment: body.NursingAssessment,
		PainScore:         body.PainScore,
	}, nil
}
