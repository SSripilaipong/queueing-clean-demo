package manage_doctor_queue

type PushVisitToDoctorQueue struct {
	DoctorId      string
	VisitId       string
	PatientName   string
	PatientGender string
	PatientAge    int
}

type CallVisitFromDoctorQueue struct {
	DoctorId string
	VisitId  string
}

type CreateDoctorQueue struct {
	DoctorId string
}

type CompleteDiagnosis struct {
	DoctorId string
}

type CheckVisits struct {
	DoctorId string
}
