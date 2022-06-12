package app

import (
	"context"
	"queueing-clean-demo/toolbox/mongodb"
)

func makeMongoDbConnection() *mongodb.Connection {
	connection, err := mongodb.CreateConnection(context.Background(), "root", "admin", "mongodb", "27017")
	if err != nil {
		panic(err)
	}
	return connection
}
