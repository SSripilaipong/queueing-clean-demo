package implementation

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"queueing-clean-demo/base"
	v "queueing-clean-demo/domain/clinical_diagnose"
	"queueing-clean-demo/domain/common/contract"
)

type VisitRepoInMongo struct {
	Collection *mongo.Collection
}

func (r *VisitRepoInMongo) FindByIdAndUpdate(id string, update func(visit *v.Visit) (*v.Visit, error)) (*v.Visit, error) {
	var visit *v.Visit
	var err error

	if visit, err = base.OptimisticLockingRetry(20, func() (*v.Visit, error) {
		if visit, err = r.FindById(id); err != nil {
			return nil, err
		}

		if visit, err = update(visit); err != nil {
			return nil, err
		}

		if visit, err = r.Save(visit); err != nil {
			return nil, err
		}

		return visit, err
	}); err != nil {
		return nil, err
	}

	return visit, nil
}

func (r *VisitRepoInMongo) FindById(id string) (*v.Visit, error) {
	var err error

	var objectId primitive.ObjectID
	if objectId, err = primitive.ObjectIDFromHex(id); err != nil {
		return nil, err
	}
	filter := bson.D{{"_id", objectId}}

	var result *mongo.SingleResult
	if result = r.Collection.FindOne(context.Background(), filter); result.Err() == mongo.ErrNoDocuments {
		return nil, common.VisitNotFoundError{}
	}

	visit := &v.Visit{}
	if err = DecodeDocument(result, visit); err != nil {
		return nil, err
	}

	return visit, nil
}

func (r *VisitRepoInMongo) Save(visit *v.Visit) (*v.Visit, error) {
	var err error

	var objectId primitive.ObjectID
	if objectId, err = primitive.ObjectIDFromHex(visit.Id); err != nil {
		return nil, err
	}
	filter := bson.D{{"_id", objectId}, {"_version", visit.GetVersion()}}

	var document map[string]any
	if document, err = makeDocument(visit.Id, visit.GetVersion()+1, visit, err); err != nil {
		return nil, err
	}

	var result *mongo.UpdateResult
	switch result, err = r.Collection.UpdateOne(context.Background(), filter, bson.D{{"$set", &document}}); {
	case result.ModifiedCount == 0:
		return nil, base.OptimisticLockFailedError{}
	case err == nil:
		break
	default:
		return nil, err
	}

	visit.IncreaseVersion()
	return visit, nil
}

func (r *VisitRepoInMongo) Create(visit *v.Visit) (*v.Visit, error) {
	var err error
	var document map[string]any

	if document, err = makeDocument(visit.Id, visit.GetVersion(), visit, err); err != nil {
		return nil, err
	}

	switch _, err = r.Collection.InsertOne(context.Background(), &document); {
	case err == nil:
		break
	case IsDuplicateKeyError(err):
		return nil, common.DuplicateVisitIdError{}
	default:
		return nil, err
	}
	return visit, nil
}
