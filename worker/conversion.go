package worker

import "encoding/json"

func makeEvent[T any](payload map[string]any) T {
	var err error

	var b []byte
	if b, err = json.Marshal(payload); err != nil {
		panic(err)
	}

	var event T
	if err = json.Unmarshal(b, &event); err != nil {
		panic(err)
	}
	return event
}
