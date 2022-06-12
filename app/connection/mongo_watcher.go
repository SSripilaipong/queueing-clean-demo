package connection

import "queueing-clean-demo/toolbox/mongo_watcher"

func MakeMongoWatcher() *mongo_watcher.MongoWatcher {
	return mongo_watcher.NewWatcher("root", "admin", "mongodb", "27017", "OPD")
}
