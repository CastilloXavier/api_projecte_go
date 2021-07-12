package inmemory

import (
	event2 "api_project/API/kit/event"
	"context"
)

// EventBus is an in-memory implementation of the event.Bus.
type EventBus struct {
	handlers map[event2.Type][]event2.Handler
}

// NewEventBus initializes a new EventBus.
func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make(map[event2.Type][]event2.Handler),
	}
}

// Publish implements the event.Bus interface.
func (b *EventBus) Publish(ctx context.Context, events []event2.Event) error {
	for _, evt := range events {
		handlers, ok := b.handlers[evt.Type()]
		if !ok {
			return nil
		}

		for _, handler := range handlers {
			handler.Handle(ctx, evt)
		}
	}

	return nil
}

// Subscribe implements the event.Bus interface.
func (b *EventBus) Subscribe(evtType event2.Type, handler event2.Handler) {
	subscribersForType, ok := b.handlers[evtType]
	if !ok {
		b.handlers[evtType] = []event2.Handler{handler}
	}

	subscribersForType = append(subscribersForType, handler)
}