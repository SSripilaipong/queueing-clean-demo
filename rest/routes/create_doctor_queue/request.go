package create_doctor_queue

import (
	"github.com/gin-gonic/gin"
	"queueing-clean-demo/domain/usecase/manage_doctor_queue/contract"
)

func makeRequest(ctx *gin.Context) (contract.CreateDoctorQueue, error) {
	doctorId := ctx.Param("doctorId")

	return contract.CreateDoctorQueue{
		DoctorId: doctorId,
	}, nil
}
