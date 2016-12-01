package coordinator

import (
	"time"
)

type EventAggregator struct {
	listeners map[string][]func(EventData)
}

func NewEventAggregator() *EventAggregator {
	ea := EventAggregator{
		listeners: make(map[string][]func(EventData)),
	}

	return &ea
}

func (ea *EventAggregator) AddListener(name string, f func(EventData)) {
	ea.listeners[name] = append(ea.listeners[name], f)
}

func (ea *EventAggregator) PublishEvent(name string, eventdata EventData) {
	if ea.listeners[name] != nil {
		for _, r := range ea.listeners[name] {
			r(eventdata)
		}
	}
}

type EventData struct {
	Name      string
	value     float64
	timestamp time.Time
}