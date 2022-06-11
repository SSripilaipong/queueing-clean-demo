package check_visits

import (
	"github.com/gin-gonic/gin"
	"queueing-clean-demo/domain/manage_doctor_queue/contract"
)

func makeRequest(ctx *gin.Context) (manage_doctor_queue.manage_doctor_queue, error) {
	doctorId := ctx.Param("doctorId")

	return manage_doctor_queue.CheckVisits{
		DoctorId: doctorId,
	}, nil
}
