package domain

type AggregateRoot struct {
	Events []Event
}

func (a *AggregateRoot) AddEvent(event Event) {
	a.Events = append(a.Events, event)
}
