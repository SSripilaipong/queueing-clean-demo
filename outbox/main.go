package outbox

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"queueing-clean-demo/toolbox/mongodb"
)

func relayLoop(ctx context.Context, ch *amqp.Channel, stream *mongo.ChangeStream) {
	running := true
	for running && stream.Next(ctx) {
		var data map[string]any
		if err := stream.Decode(&data); err != nil {
			panic(err)
		}
		fmt.Printf("Outbox: %v\n", data)

		events := extractLatestEvents(data)
		publishEvents(ch, events)

		select {
		case <-ctx.Done():
			running = false
		default:
		}
	}
}

func extractLatestEvents(data map[string]any) []any {
	var ok bool

	var document map[string]any
	if document, ok = data["fullDocument"].(map[string]any); !ok {
		var desc map[string]any
		if desc, ok = data["updateDescription"].(map[string]any); !ok {
			panic(fmt.Errorf("unexpected error"))
		}
		if document, ok = desc["updatedFields"].(map[string]any); !ok {
			panic(fmt.Errorf("no document in change stream"))
		}
	}

	var events []any
	fmt.Printf("%#v\n", document)
	if events, ok = document["_latestEvents"].(primitive.A); !ok {
		panic(fmt.Errorf("no latestEvents in document"))
	}

	return events
}

func RunOutboxRelay(ctx context.Context) {
	rbConn, rbCh := makeChannel()
	defer rbConn.Close()
	defer rbCh.Close()

	dbConn, err := mongodb.CreateConnection(context.Background(), "root", "admin", "mongodb", "27017")
	if err != nil {
		panic(err)
	}
	defer dbConn.Disconnect(context.Background())
	db := dbConn.Client.Database("OPD")

	stream, err := db.Watch(context.Background(), mongo.Pipeline{}, options.ChangeStream().SetFullDocument(options.UpdateLookup))
	if err != nil {
		panic(err)
	}

	relayLoop(ctx, rbCh, stream)
}

func publishEvents(ch *amqp.Channel, array []any) {
	var err error

	for _, event := range array {
		var j []byte
		if j, err = json.Marshal(event); err != nil {
			panic(err)
		}

		if err = ch.Publish("", "allEvents", false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        j,
		}); err != nil {
			panic(err)
		}
		fmt.Println("published ", j)
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
		defer conn.Close()
		panic(err)
	}
	return conn, ch
}
