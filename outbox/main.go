package outbox

import (
	"context"
	"fmt"
	"queueing-clean-demo/toolbox/mongo_watcher"
	"queueing-clean-demo/toolbox/rabbitmq"
)

func runOutboxRelay(ctx context.Context, exited chan struct{}) {
	rabbit := rabbitmq.NewClient("root", "admin", "rabbitmq", "5672")
	defer rabbit.Disconnect()

	watcher := mongo_watcher.NewWatcher("root", "admin", "mongodb", "27017", "OPD")
	defer watcher.Disconnect()

	watcherLoop(ctx, watcher, func(data map[string]any) {
		events := extractLatestEvents(data)
		for _, event := range events {
			rabbit.Publish("allEvents", event)
			fmt.Printf("published: %#v\n", event)
		}
	})

	exited <- struct{}{}
	fmt.Println("outbox exited")
}
