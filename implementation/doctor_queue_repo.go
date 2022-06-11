package implementation

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"queueing-clean-demo/base"
	. "queueing-clean-demo/domain/manage_doctor_queue/contract"
)

type DoctorQueueRepoInMongo struct {
	Collection *mongo.Collection
}

func (r *DoctorQueueRepoInMongo) FindByDoctorIdAndUpdate(id string, update func(queue *DoctorQueueRepr) (*DoctorQueueRepr, error)) (*DoctorQueueRepr, error) {
	var queue *DoctorQueueRepr
	var err error

	if queue, err = base.OptimisticLockingRetry(20, func() (*DoctorQueueRepr, error) {
		if queue, err = r.FindByDoctorId(id); err != nil {
			return nil, err
		}

		if queue, err = update(queue); err != nil {
			return nil, err
		}

		if queue, err = r.Save(queue); err != nil {
			return nil, err
		}

		return queue, err
	}); err != nil {
		return nil, err
	}

	return queue, nil
}

func (r *DoctorQueueRepoInMongo) FindByDoctorId(id string) (*DoctorQueueRepr, error) {
	var err error

	var objectId primitive.ObjectID
	if objectId, err = primitive.ObjectIDFromHex(id); err != nil {
		return nil, err
	}
	filter := bson.D{{"_id", objectId}}

	var result *mongo.SingleResult
	if result = r.Collection.FindOne(context.Background(), filter); result.Err() == mongo.ErrNoDocuments {
		return nil, DoctorQueueNotFoundError{}
	}

	queue := &DoctorQueueRepr{}
	if err = DecodeDocument(result, queue); err != nil {
		return nil, err
	}

	return queue, nil
}

func (r *DoctorQueueRepoInMongo) Save(queue *DoctorQueueRepr) (*DoctorQueueRepr, error) {
	var err error

	var objectId primitive.ObjectID
	if objectId, err = primitive.ObjectIDFromHex(queue.DoctorId); err != nil {
		return nil, err
	}
	filter := bson.D{{"_id", objectId}, {"_version", queue.GetVersion()}}

	var document map[string]any
	if document, err = MakeDocument(queue.DoctorId, queue); err != nil {
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

	queue.IncreaseVersion()
	return queue, nil
}

func (r *DoctorQueueRepoInMongo) Create(queue *DoctorQueueRepr) (*DoctorQueueRepr, error) {
	var err error
	var document map[string]any

	if document, err = MakeDocument(queue.DoctorId, queue); err != nil {
		return nil, err
	}

	switch _, err = r.Collection.InsertOne(context.Background(), &document); {
	case err == nil:
		break
	case IsDuplicateKeyError(err):
		return nil, DuplicateDoctorQueueIdError{}
	default:
		return nil, err
	}
	return queue, nil
}
