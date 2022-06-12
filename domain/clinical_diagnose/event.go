package clinical_diagnose

type VisitAssessedEvent struct {
	VisitId string `json:"VisitId"`
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	Age     int    `json:"age"`
}
