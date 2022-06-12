package internal

import (
	. "queueing-clean-demo/domain/clinical_diagnose"
)

func AssessmentFromSubmitAssessmentRequest(r SubmitAssessment) Assessment {
	return Assessment{
		NursingAssessment: r.NursingAssessment,
		PainScore:         r.PainScore,
	}
}

func VisitResponseFromVisit(visit *Visit) VisitResponse {
	resp := VisitResponse{
		VisitId: visit.Id,
	}

	if a := visit.Assessment; a != nil {
		resp.Assessment = AssessmentResponse{
			NursingAssessment: a.NursingAssessment,
			PainScore:         a.PainScore,
		}
	}
	return resp
}
