package _deps

import (
	. "queueing-clean-demo/domain/clinical_diagnose"
)

type IVisitRepo interface {
	FindByIdAndUpdate(id string, update func(visit *Visit) (*Visit, error)) (*Visit, error)
	Create(visit *Visit) (*Visit, error)
}
