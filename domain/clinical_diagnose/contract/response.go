package clinical_diagnose

type VisitResponse struct {
	VisitId    string
	Assessment AssessmentResponse
}

type AssessmentResponse struct {
	NursingAssessment string
	PainScore         int
}
