package outbox

import (
	"context"
	"queueing-clean-demo/toolbox/mongo_watcher"
)

func watcherLoop(ctx context.Context, watcher *mongo_watcher.MongoWatcher, handle func(data map[string]any)) {
	running := true
	for running && watcher.Next(ctx) {
		data := watcher.Get()

		handle(data)

		select {
		case <-ctx.Done():
			running = false
		default:
		}
	}
}
