package base

type IAggregate interface {
	GetEvents() []IEvent
	AppendEvent(event IEvent)
	GetVersion() int
	IncreaseVersion()
	SetVersion(int)
}

type Aggregate struct {
	version int
	events  events
}

func (a *Aggregate) GetEvents() []IEvent {
	return a.events.getArray()
}

func (a *Aggregate) AppendEvent(event IEvent) {
	a.events.append(event)
}

func (a *Aggregate) GetVersion() int {
	return a.version
}

func (a *Aggregate) SetVersion(version int) {
	a.version = version
}

func (a *Aggregate) IncreaseVersion() {
	a.version++
}

type events struct {
	array []IEvent
}

func (e *events) append(event IEvent) {
	e.array = append(e.array, event)
}

func (e *events) getArray() []IEvent {
	result := make([]IEvent, len(e.array))
	for i, e := range e.array {
		result[i] = e
	}
	return result
}
