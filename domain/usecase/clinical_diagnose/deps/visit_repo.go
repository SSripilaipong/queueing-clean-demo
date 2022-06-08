package deps

import v "queueing-clean-demo/domain"

type IVisitRepo interface {
	FindByIdAndUpdate(id string, update func(visit *v.Visit) (*v.Visit, error)) (*v.Visit, error)
	Create(visit *v.Visit) (*v.Visit, error)
}
