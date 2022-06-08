package check_visits

import (
	"github.com/gin-gonic/gin"
	"queueing-clean-demo/domain/usecase/manage_doctor_queue/contract"
)

func makeRequest(ctx *gin.Context) (contract.CheckVisits, error) {
	doctorId := ctx.Param("doctorId")

	return contract.CheckVisits{
		DoctorId: doctorId,
	}, nil
}
