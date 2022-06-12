package outbox

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	if events, ok = document["_latestEvents"].(primitive.A); !ok {
		panic(fmt.Errorf("no latestEvents in document"))
	}

	return events
}
