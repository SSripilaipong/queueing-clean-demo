package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection struct {
	Client *mongo.Client
}

func (c *Connection) Disconnect(ctx context.Context) {
	if err := c.Client.Disconnect(ctx); err != nil {
		panic(err)
	}
}

func CreateConnection(ctx context.Context, username, password, host, port string) (*Connection, error) {
	var client *mongo.Client
	var err error

	uri := fmt.Sprintf(UriTemplate, username, password, host, port)
	if client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri)); err != nil {
		return nil, err
	}
	return &Connection{Client: client}, nil
}

const UriTemplate = "mongodb://%s:%s@%s:%s/?maxPoolSize=20&w=majority"
