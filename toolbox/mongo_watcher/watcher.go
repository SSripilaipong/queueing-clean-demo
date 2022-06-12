package mongo_watcher

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"queueing-clean-demo/toolbox/mongodb"
)

type MongoWatcher struct {
	dbConn *mongodb.Connection
	stream *mongo.ChangeStream
}

func NewWatcher(username string, password string, host string, port string, dbName string) *MongoWatcher {
	dbConn, err := mongodb.CreateConnection(context.Background(), username, password, host, port)
	if err != nil {
		panic(err)
	}
	db := dbConn.Client.Database(dbName)

	stream, err := db.Watch(context.Background(), mongo.Pipeline{}, options.ChangeStream().SetFullDocument(options.UpdateLookup))
	if err != nil {
		panic(err)
	}

	return &MongoWatcher{
		dbConn: dbConn,
		stream: stream,
	}
}

func (w *MongoWatcher) Stream() *mongo.ChangeStream {
	return w.stream
}

func (w *MongoWatcher) Next(ctx context.Context) bool {
	return w.stream.Next(ctx)
}

func (w *MongoWatcher) Get() map[string]any {
	data := make(map[string]any)
	if err := w.stream.Decode(&data); err != nil {
		panic(err)
	}
	return data
}

func (w *MongoWatcher) Disconnect() {
	w.dbConn.Disconnect(context.Background())
}
