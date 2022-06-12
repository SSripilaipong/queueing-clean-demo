package implementation

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	. "queueing-clean-demo/domain/clinical_diagnose"
	"queueing-clean-demo/domain/common"
	"queueing-clean-demo/toolbox/mongodb"
)

type visitRepoInMongo struct {
	repo mongodb.MongoRepo[*VisitRepr]
}

func NewVisitRepoInMongo(collection *mongo.Collection) IVisitRepo {
	return &visitRepoInMongo{
		repo: mongodb.MongoRepo[*VisitRepr]{
			Collection:           collection,
			MaxOptimisticRetries: 20,
		},
	}
}

func (r *visitRepoInMongo) FindByIdAndUpdate(id string, update func(visit *VisitRepr) (*VisitRepr, error)) (*VisitRepr, error) {
	var err error

	var objectId primitive.ObjectID
	if objectId, err = primitive.ObjectIDFromHex(id); err != nil {
		panic(err)
	}

	result, err := r.repo.FindByIdAndUpdate(objectId, update)

	switch err.(type) {
	case mongodb.AggregateNotFoundError:
		return nil, common.VisitNotFoundError{}
	case nil:
		return result, nil
	}
	panic(err)
}

func (r *visitRepoInMongo) Create(visit *VisitRepr) (*VisitRepr, error) {
	var err error

	var objectId primitive.ObjectID
	if objectId, err = primitive.ObjectIDFromHex(visit.Id); err != nil {
		panic(err)
	}

	_, err = r.repo.Create(objectId, visit)

	switch err.(type) {
	case mongodb.DuplicateIdError:
		return nil, common.DuplicateVisitIdError{}
	case nil:
		return visit, nil
	}
	panic(err)
}
