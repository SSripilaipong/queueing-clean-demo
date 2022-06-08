package domain

import (
	"queueing-clean-demo/domain/usecase/clinical_diagnose/contract"
	contract2 "queueing-clean-demo/domain/usecase/manage_doctor_queue/contract"
)

func AssessmentFromSubmitAssessmentRequest(r contract.SubmitAssessment) Assessment {
	return Assessment{
		NursingAssessment: r.NursingAssessment,
		PainScore:         r.PainScore,
	}
}

func VisitResponseFromVisit(visit *Visit) contract.Visit {
	resp := contract.Visit{
		VisitId: visit.Id,
	}

	if a := visit.Assessment; a != nil {
		resp.Assessment = contract2.AssessmentResponse{
			NursingAssessment: a.NursingAssessment,
			PainScore:         a.PainScore,
		}
	}
	return resp
}
