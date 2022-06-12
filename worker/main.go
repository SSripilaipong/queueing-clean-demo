package worker

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"queueing-clean-demo/domain/clinical_diagnose"
	"queueing-clean-demo/domain/manage_doctor_queue"
	"queueing-clean-demo/toolbox/mongodb"
)

func RunWorker(ctx context.Context) {

	runWorkerLoop(ctx, func(deps *Deps, delivery amqp.Delivery) {
		var message Message
		if err := json.Unmarshal(delivery.Body, &message); err != nil {
			panic(err)
		}

		switch message.Name {
		case "VisitAssessedEvent":
			var e clinical_diagnose.VisitAssessedEvent
			mapToStruct(message.Payload, &e)
			handleVisitAssessedEvent(deps, e)
		}
	})

}

func mapToStruct(payload map[string]any, e *clinical_diagnose.VisitAssessedEvent) {
	var err error

	var b []byte
	if b, err = json.Marshal(payload); err != nil {
		panic(err)
	}

	if err = json.Unmarshal(b, e); err != nil {
		panic(err)
	}
}

func handleVisitAssessedEvent(deps *Deps, e clinical_diagnose.VisitAssessedEvent) {
	var err error
	switch _, err = deps.ManageDoctorQueueUsecase.PushVisit(manage_doctor_queue.PushVisitToDoctorQueue{
		DoctorId:      "629c93cae6509bc3a7b1aaf7",
		VisitId:       e.VisitId,
		PatientName:   e.Name,
		PatientGender: e.Gender,
		PatientAge:    e.Age,
	}); err {
	case manage_doctor_queue.VisitAlreadyExistsError{}:
		return
	case nil:
		return
	}
	panic(err)
}

type Message struct {
	Name    string         `json:"name"`
	Payload map[string]any `json:"payload"`
}

func runWorkerLoop(ctx context.Context, handle func(*Deps, amqp.Delivery)) {
	var err error

	var mgConn *mongodb.Connection
	if mgConn, err = mongodb.CreateConnection(ctx, "root", "admin", "mongodb", "27017"); err != nil {
		panic(err)
	}
	defer mgConn.Disconnect(ctx)

	deps := createDeps(mgConn.Client.Database("OPD"))

	rbConn, rbCh := makeChannel()
	defer func(rbConn *amqp.Connection) {
		_ = rbConn.Close()
	}(rbConn)
	defer func(rbCh *amqp.Channel) {
		_ = rbCh.Close()
	}(rbCh)

	var delivery <-chan amqp.Delivery
	if delivery, err = rbCh.Consume(
		"allEvents",
		"",
		false,
		false,
		false,
		false,
		nil,
	); err != nil {
		panic(err)
	}

	running := true
	for running {
		select {
		case message := <-delivery:
			handle(deps, message)
			if err := message.Ack(false); err != nil {
				panic(err)
			}

		case <-ctx.Done():
			running = false
		}
	}
}
func makeChannel() (*amqp.Connection, *amqp.Channel) {
	var err error

	var conn *amqp.Connection
	if conn, err = amqp.Dial("amqp://root:admin@rabbitmq:5672"); err != nil {
		panic(err)
	}

	var ch *amqp.Channel
	if ch, err = conn.Channel(); err != nil {
		defer func(conn *amqp.Connection) {
			_ = conn.Close()
		}(conn)
		panic(err)
	}
	return conn, ch
}
