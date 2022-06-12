package worker

import (
	"queueing-clean-demo/domain/clinical_diagnose"
	"queueing-clean-demo/worker/deps"
	"queueing-clean-demo/worker/handler/visit_assessed"
)

func messageRoute(msg message, deps deps.IWorkerDeps) {
	switch msg.Name {
	case "VisitAssessedEvent":
		visit_assessed.Handler(deps, makeEvent[clinical_diagnose.VisitAssessedEvent](msg.Payload))
	}
}
