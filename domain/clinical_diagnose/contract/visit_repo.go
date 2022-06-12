package clinical_diagnose

type IVisitRepo interface {
	FindByIdAndUpdate(id string, update func(visit *VisitRepr) (*VisitRepr, error)) (*VisitRepr, error)
	Create(visit *VisitRepr) (*VisitRepr, error)
}
