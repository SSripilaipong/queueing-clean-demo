package implementation

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	. "queueing-clean-demo/domain/manage_doctor_queue"
	"queueing-clean-demo/toolbox/mongodb"
)

type doctorQueueRepoInMongo struct {
	repo mongodb.MongoRepo[*DoctorQueueRepr]
}

func NewDoctorQueueRepoInMongo(collection *mongo.Collection) IDoctorQueueRepo {
	return &doctorQueueRepoInMongo{
		repo: mongodb.MongoRepo[*DoctorQueueRepr]{
			Collection:           collection,
			MaxOptimisticRetries: 20,
		},
	}
}

func (r *doctorQueueRepoInMongo) FindByDoctorIdAndUpdate(id string, update func(queue *DoctorQueueRepr) (*DoctorQueueRepr, error)) (*DoctorQueueRepr, error) {
	var err error

	var objectId primitive.ObjectID
	if objectId, err = primitive.ObjectIDFromHex(id); err != nil {
		panic(err)
	}

	result, err := r.repo.FindByIdAndUpdate(objectId, update)

	switch err.(type) {
	case mongodb.AggregateNotFoundError:
		return nil, DoctorQueueNotFoundError{}
	case nil:
		return result, nil
	}
	panic(err)
}

func (r *doctorQueueRepoInMongo) FindByDoctorId(id string) (*DoctorQueueRepr, error) {
	var err error

	var objectId primitive.ObjectID
	if objectId, err = primitive.ObjectIDFromHex(id); err != nil {
		return nil, err
	}

	var result *DoctorQueueRepr
	result, err = r.repo.FindOne(bson.D{{"_id", objectId}})

	switch err.(type) {
	case mongodb.AggregateNotFoundError:
		return nil, DoctorQueueNotFoundError{}
	case nil:
		return result, nil
	}
	panic(err)
}

func (r *doctorQueueRepoInMongo) Create(queue *DoctorQueueRepr) (*DoctorQueueRepr, error) {
	var err error

	var objectId primitive.ObjectID
	if objectId, err = primitive.ObjectIDFromHex(queue.DoctorId); err != nil {
		return nil, err
	}

	_, err = r.repo.Create(objectId, queue)

	switch err.(type) {
	case mongodb.DuplicateIdError:
		return nil, DuplicateDoctorQueueIdError{}
	case nil:
		return queue, nil
	}
	panic(err)
}
