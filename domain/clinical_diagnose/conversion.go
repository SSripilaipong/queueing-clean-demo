package _clinical_diagnose

import (
	"queueing-clean-demo/domain/clinical_diagnose/contract"
)

func AssessmentFromSubmitAssessmentRequest(r clinical_diagnose.SubmitAssessment) Assessment {
	return Assessment{
		NursingAssessment: r.NursingAssessment,
		PainScore:         r.PainScore,
	}
}

func VisitResponseFromVisit(visit *Visit) clinical_diagnose.VisitResponse {
	resp := clinical_diagnose.VisitResponse{
		VisitId: visit.Id,
	}

	if a := visit.Assessment; a != nil {
		resp.Assessment = clinical_diagnose.AssessmentResponse{
			NursingAssessment: a.NursingAssessment,
			PainScore:         a.PainScore,
		}
	}
	return resp
}
