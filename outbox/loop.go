package outbox

import (
	"context"
	"queueing-clean-demo/outbox/deps"
)

func watcherLoop(ctx context.Context, watcher deps.IStream, handle func(data map[string]any)) {
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
