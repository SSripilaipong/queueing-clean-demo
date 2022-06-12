package submit_assessment

import (
	"github.com/gin-gonic/gin"
	"queueing-clean-demo/domain/clinical_diagnose"
)

type SubmitAssessment struct {
	NursingAssessment string `json:"comment"`
	PainScore         int    `json:"painScore"`
}

func makeRequest(ctx *gin.Context) (clinical_diagnose.SubmitAssessment, error) {
	visitId := ctx.Param("visitId")

	var body SubmitAssessment
	if err := ctx.ShouldBindJSON(&body); err != nil {
		return clinical_diagnose.SubmitAssessment{}, err
	}

	return clinical_diagnose.SubmitAssessment{
		VisitId:           visitId,
		NursingAssessment: body.NursingAssessment,
		PainScore:         body.PainScore,
	}, nil
}
