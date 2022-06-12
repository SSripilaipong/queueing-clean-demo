package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"queueing-clean-demo/base"
)

type MongoRepo[R base.IAggregateRepr] struct {
	Collection           *mongo.Collection
	MaxOptimisticRetries int
}

func (r *MongoRepo[R]) FindByIdAndUpdate(id primitive.ObjectID, update func(repr R) (R, error)) (repr R, err error) {
	repr, err = base.OptimisticLockingRetry(r.MaxOptimisticRetries, func() (R, error) {
		if repr, err = r.FindById(id); err != nil {
			return repr, err
		}

		if repr, err = update(repr); err != nil {
			return repr, err
		}

		if repr, err = r.Save(id, repr); err != nil {
			return repr, err
		}

		return repr, nil
	})

	return repr, err
}

func (r *MongoRepo[R]) FindOne(filter bson.D) (repr R, err error) {
	var result *mongo.SingleResult
	if result = r.Collection.FindOne(context.Background(), filter); result.Err() == mongo.ErrNoDocuments {
		return repr, AggregateNotFoundError{}
	}

	ptr := new(R)
	err = DecodeDocument(result, ptr)
	return *ptr, err
}

func (r *MongoRepo[R]) Save(id primitive.ObjectID, repr R) (R, error) {
	var err error

	filter := bson.D{{"_id", id}, {"_version", repr.GetVersion()}}

	var document map[string]any
	if document, err = MakeDocument(id, repr); err != nil {
		return repr, err
	}

	var result *mongo.UpdateResult
	switch result, err = r.Collection.UpdateOne(context.Background(), filter, bson.D{{"$set", &document}}); {
	case result.ModifiedCount == 0:
		return repr, base.OptimisticLockFailedError{}
	case err == nil:
		break
	default:
		return repr, err
	}

	repr.IncreaseVersion()
	return repr, nil
}

func (r *MongoRepo[R]) Create(id primitive.ObjectID, repr R) (R, error) {
	var err error
	var document map[string]any

	if document, err = MakeDocument(id, repr); err != nil {
		return repr, err
	}

	switch _, err = r.Collection.InsertOne(context.Background(), &document); {
	case err == nil:
		break
	case IsDuplicateKeyError(err):
		return repr, DuplicateIdError{}
	default:
		return repr, err
	}
	return repr, nil
}

func (r *MongoRepo[R]) FindById(id primitive.ObjectID) (R, error) {
	return r.FindOne(bson.D{{"_id", id}})
}
