package connection

import (
	"context"
	"queueing-clean-demo/toolbox/mongodb"
)

func MakeMongoDbConnection() *mongodb.Connection {
	connection, err := mongodb.CreateConnection(context.Background(), "root", "admin", "mongodb", "27017")
	if err != nil {
		panic(err)
	}
	return connection
}
