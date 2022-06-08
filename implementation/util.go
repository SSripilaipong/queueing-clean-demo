package implementation

import (
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"queueing-clean-demo/base"
	"reflect"
)

func DecodeDocument[T base.IAggregate](result *mongo.SingleResult, object T) error {
	var err error
	var m map[string]any
	if err = result.Decode(&m); err != nil {
		return err
	}

	var j []byte
	payload := m["payload"]
	if j, err = json.Marshal(payload); err != nil {
		return err
	}

	if err = json.Unmarshal(j, &object); err != nil {
		return err
	}

	object.SetVersion(int(m["_version"].(int32)))
	return nil
}

func makeDocument(id string, version int, aggregate base.IAggregate, err error) (map[string]any, error) {
	document := make(map[string]any)

	if document["payload"], err = structToMap(aggregate); err != nil {
		return nil, err
	}
	if document["_id"], err = primitive.ObjectIDFromHex(id); err != nil {
		return nil, err
	}
	if document["_latestEvents"], err = makeEvents(aggregate.GetEvents()); err != nil {
		return nil, err
	}
	document["_version"] = version
	return document, nil
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
