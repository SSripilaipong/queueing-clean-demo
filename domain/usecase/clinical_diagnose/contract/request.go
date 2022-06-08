package contract

type SubmitAssessment struct {
	VisitId           string
	NursingAssessment string
	PainScore         int
}

type CreateVisit struct {
	Name   string
	Gender string
	Age    int
}
