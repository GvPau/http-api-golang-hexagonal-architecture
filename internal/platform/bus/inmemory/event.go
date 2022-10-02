package inmemory

import (
	"api/kit/event"
	"context"
)

type EventBus struct {
	handlers map[event.Type][]event.Handler
}

func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make(map[event.Type][]event.Handler),
	}
}

func (b *EventBus) Publish(ctx context.Context, events []event.Event) error {
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

func (b *EventBus) Subscribe(evType event.Type, handler event.Handler) {
	subscribersForType, ok := b.handlers[evType]
	if !ok {
		b.handlers[evType] = []event.Handler{handler}
	}

	subscribersForType = append(subscribersForType, handler)
}
