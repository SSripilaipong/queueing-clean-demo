package deps

import (
	"context"
)

type IOutboxDeps interface {
	Broker() IPublisher
	Stream() IStream
	Destroy()
}

type IPublisher interface {
	Publish(key string, event any)
}

type IStream interface {
	Next(ctx context.Context) bool
	Get() map[string]any
}
