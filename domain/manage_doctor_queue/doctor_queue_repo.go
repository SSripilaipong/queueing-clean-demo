package manage_doctor_queue

type IDoctorQueueRepo interface {
	Create(queue *DoctorQueueRepr) (*DoctorQueueRepr, error)
	FindByDoctorIdAndUpdate(doctorId string, update func(queue *DoctorQueueRepr) (*DoctorQueueRepr, error)) (*DoctorQueueRepr, error)
	FindByDoctorId(doctorId string) (*DoctorQueueRepr, error)
}
