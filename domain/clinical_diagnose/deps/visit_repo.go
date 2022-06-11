package _deps

import (
	. "queueing-clean-demo/domain/clinical_diagnose/contract"
)

type IVisitRepo interface {
	FindByIdAndUpdate(id string, update func(visit *VisitRepr) (*VisitRepr, error)) (*VisitRepr, error)
	Create(visit *VisitRepr) (*VisitRepr, error)
}
