package push_visit_to_doctor_queue

import (
	"github.com/gin-gonic/gin"
	"queueing-clean-demo/domain/usecase/manage_doctor_queue/contract"
)

type PushVisit struct {
	VisitId       string
	PatientName   string
	PatientGender string
	PatientAge    int
}

func makeRequest(ctx *gin.Context) (contract.PushVisitToDoctorQueue, error) {
	doctorId := ctx.Param("doctorId")

	var body PushVisit
	if err := ctx.ShouldBindJSON(&body); err != nil {
		return contract.PushVisitToDoctorQueue{}, err
	}

	return contract.PushVisitToDoctorQueue{
		DoctorId:      doctorId,
		VisitId:       body.VisitId,
		PatientName:   body.PatientName,
		PatientGender: body.PatientGender,
		PatientAge:    body.PatientAge,
	}, nil
}
