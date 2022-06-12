package worker

type message struct {
	Name    string         `json:"name"`
	Payload map[string]any `json:"payload"`
}
