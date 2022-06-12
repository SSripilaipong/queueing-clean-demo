package internal

import (
	clinical_diagnose2 "queueing-clean-demo/domain/clinical_diagnose"
)

func AssessmentFromSubmitAssessmentRequest(r clinical_diagnose2.SubmitAssessment) Assessment {
	return Assessment{
		NursingAssessment: r.NursingAssessment,
		PainScore:         r.PainScore,
	}
}

func VisitResponseFromVisit(visit *Visit) clinical_diagnose2.VisitResponse {
	resp := clinical_diagnose2.VisitResponse{
		VisitId: visit.Id,
	}

	if a := visit.Assessment; a != nil {
		resp.Assessment = clinical_diagnose2.AssessmentResponse{
			NursingAssessment: a.NursingAssessment,
			PainScore:         a.PainScore,
		}
	}
	return resp
}
