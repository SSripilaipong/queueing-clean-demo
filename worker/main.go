package worker

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"queueing-clean-demo/domain/clinical_diagnose"
	"queueing-clean-demo/worker/deps"
	"queueing-clean-demo/worker/handler/visit_assessed"
)

func RunWorker(ctx context.Context) {

	workerLoop(ctx, func(deps *worker_deps.Deps, delivery amqp.Delivery) {
		var msg message
		if err := json.Unmarshal(delivery.Body, &msg); err != nil {
			panic(err)
		}

		switch msg.Name {
		case "VisitAssessedEvent":
			visit_assessed.Handler(deps, makeEvent[clinical_diagnose.VisitAssessedEvent](msg.Payload))
		}
	})
}
