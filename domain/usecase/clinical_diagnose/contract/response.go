package contract

import (
	"queueing-clean-demo/domain/usecase/manage_doctor_queue/contract"
)

type Visit struct {
	VisitId    string
	Assessment contract.AssessmentResponse
}
