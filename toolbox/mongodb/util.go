package mongodb

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"queueing-clean-demo/base"
	"reflect"
)

func DecodeDocument(result *mongo.SingleResult, ptr any) (err error) {
	var ok bool

	var doc map[string]any
	if err = result.Decode(&doc); err != nil {
		return err
	}

	var obj map[string]any
	if obj, ok = doc["payload"].(map[string]any); !ok {
		return fmt.Errorf("unexpected error")
	}

	obj["_aggregate"] = map[string]any{
		"_version":      doc["_version"],
		"_latestEvents": doc["_latestEvents"],
	}

	var j []byte
	if j, err = json.Marshal(obj); err != nil {
		return err
	}

	if err = json.Unmarshal(j, ptr); err != nil {
		return err
	}
	return nil
}

func MakeDocument(id primitive.ObjectID, aggregate base.IAggregateRepr) (doc map[string]any, err error) {
	doc = make(map[string]any)
	var payload map[string]any
	payload, err = structToMap(aggregate)
	delete(payload, "_aggregate")

	doc["_id"] = id
	doc["payload"] = payload
	doc["_version"] = aggregate.GetVersion()
	doc["_latestEvents"], err = makeEvents(aggregate.GetEvents())
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func makeEvents(events []base.IEvent) ([]map[string]any, error) {
	result := make([]map[string]any, len(events))

	for i := 0; i < len(events); i++ {
		e := events[i]
		if m, err := structToMap(e); err != nil {
			return nil, err
		} else {
			name := reflect.TypeOf(e).Name()
			result[i] = map[string]any{"name": name, "payload": m}
		}
	}
	return result, nil
}

func structToMap(object any) (map[string]any, error) {
	var j []byte
	var err error

	j, err = json.Marshal(object)
	if err != nil {
		return nil, err
	}
	var m map[string]any
	err = json.Unmarshal(j, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func IsDuplicateKeyError(err error) bool {
	// handles SERVER-7164 and SERVER-11493
	for ; err != nil; err = errors.Unwrap(err) {
		if e, ok := err.(mongo.ServerError); ok {
			return e.HasErrorCode(11000) || e.HasErrorCode(11001) || e.HasErrorCode(12582) ||
				e.HasErrorCodeWithMessage(16460, " E11000 ")
		}
	}
	return false
}
