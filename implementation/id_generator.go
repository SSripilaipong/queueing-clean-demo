package implementation

import "go.mongodb.org/mongo-driver/bson/primitive"

type IdGenerator struct {
}

func (IdGenerator) GetId() string {
	return primitive.NewObjectID().Hex()
}
