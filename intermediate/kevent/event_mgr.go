package kevent

import (
	"context"

	"git.shiyou.kingsoft.com/go/errors"
	"golang.org/x/exp/slog"
)

type FEventListener func(context.Context, *Event) error

type EventManager struct {
	mapListenerList map[int32][]FEventListener // <id, []func>事件以及处理函数
}

var eventMgr = NewEventManager()

func NewEventManager() *EventManager {
	return &EventManager{
		mapListenerList: make(map[int32][]FEventListener),
	}
}

func AddEventListener(eventID int32, listener FEventListener) {
	eventMgr.addListener(eventID, listener)
}

func (e *EventManager) addListener(eventID int32, listener FEventListener) {
	listeners := e.mapListenerList[eventID]

	if listeners == nil {
		listeners = make([]FEventListener, 0, BaseListenerCapacity)
	}

	listeners = append(listeners, listener)
	e.mapListenerList[eventID] = listeners
}

func (e *EventManager) executeListener(ctx context.Context, event *Event, eventID int32) error {
	listenerList, ok := e.mapListenerList[eventID]
	if !ok {
		slog.WarnCtx(ctx, "no  listener for evenID", "eventID", eventID, "event", *event)
		return nil
	}

	for _, listener := range listenerList {
		if err := listener(ctx, event); err != nil {
			return errors.Wrap(err)
		}
	}

	return nil
}

func (e *EventManager) executeEvent(ctx context.Context, event *Event) error {
	slog.InfoCtx(ctx, "dispatchEvent", "event", event)

	return errors.Wrap(e.executeListener(ctx, event, event.EventID))
}
