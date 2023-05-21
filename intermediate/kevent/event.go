package kevent

import (
	"context"
	"fmt"
	"git.shiyou.kingsoft.com/go/log"
	"sync"

	"golang.org/x/exp/slog"
)

const (
	EventDeep = 5
)

const (
	BaseListenerCapacity = 8
	BaseEventBufferSize  = 8
)

type EventListener func(ctx context.Context)

type Event struct {
	EventRoleID string // role id
	EventID     int32  // id
	EventIDList string // ids
	EventData   []any
}

var eventDummy Event

// 事件池
var poolEvent = sync.Pool{
	New: func() any {
		return new(Event)
	},
}

func NewEvent() *Event {
	event, ok := poolEvent.Get().(*Event)
	if !ok {
		slog.Error("new event err")
	}

	*event = eventDummy
	return event
}

type EventBuffer struct {
	allEvent  [][]*Event
	eventDeep int // 事件 deep
}

func (e *EventBuffer) InitEventBuffer(ctx context.Context) {
	if e.allEvent == nil {
		e.allEvent = make([][]*Event, 0, EventDeep)
	}

	for idx := 0; idx < EventDeep; idx++ {
		if e.allEvent[idx] != nil {
			continue
		}

		e.allEvent[idx] = make([]*Event, 0, BaseEventBufferSize)
	}
}

func (e *EventBuffer) ResetEventBuffer() {
	for idx, eventList := range e.allEvent {
		if eventList == nil {
			break
		}

		for _, event := range eventList {
			poolEvent.Put(event)
		}

		e.allEvent[idx] = eventList[:0]
	}

	e.eventDeep = 0
}

func (e *EventBuffer) dispatchEvent(ctx context.Context, roleID string, eventID int32, data ...any) {
	slog.InfoCtx(ctx, "DispatchEvent", "roleID", roleID, "eventID", eventID, "param", data)

	if e.eventDeep >= EventDeep {
		slog.WarnCtx(ctx, "can not dispatch", "roleID", roleID, "eventID", eventID, "param", data, "deep", e.eventDeep)

		return
	}

	event := NewEvent()
	event.EventRoleID = roleID
	event.EventID = eventID
	event.EventData = data

	events := e.allEvent[e.eventDeep]
	events = append(events, event)
	e.allEvent[e.eventDeep] = events
}

func (e *EventBuffer) ReallyDispatchAllEvent(ctx context.Context) error {
	for _, events := range e.allEvent {
		e.eventDeep++

		// 分发 event
		for _, event := range events {
			if err := eventMgr.executeEvent(ctx, event); err != nil {
				log.Errorx(ctx, "ReallyDispatchAllEvent", "err", err.Error(), "event", fmt.Sprintf("%+v", event))
			}
		}
	}

	return nil
}

func ExecuteEvent(ctx context.Context, roleID string, eventID int32, data ...any) {
	event := NewEvent()
	event.EventRoleID = roleID
	event.EventID = eventID
	event.EventData = data

	if err := eventMgr.executeEvent(ctx, event); err != nil {
		slog.ErrorCtx(ctx, "ExecuteEvent", "err", err.Error(), "event", fmt.Sprintf("%+v", event))
	}
}

func DispatchEvent(ctx context.Context, roleID string, eventID int32, data ...any) {
	e := getEventBuffer(ctx)
	if e == nil {
		slog.ErrorCtx(ctx, "getEventBuffer err")
		return
	}

	e.dispatchEvent(ctx, roleID, eventID, data)
}

func getEventBuffer(ctx context.Context) *EventBuffer {
	value := ctx.Value(CtxEventDataKey)
	if value == nil {
		slog.ErrorCtx(ctx, "value is Nil")
		return nil
	}

	e, ok := value.(*EventBuffer)
	if !ok {
		slog.ErrorCtx(ctx, "Value EventBuffer error")

		return nil
	}

	return e
}
