package kevent

import (
	"context"
	"golang.org/x/exp/slog"
)

type ContextKey string

const (
	CtxEventDataKey ContextKey = "__event_data"
)

type Plugin struct{}

func (p *Plugin) ProcessStart(ctx context.Context) (context.Context, error) {
	e := EventBuffer{}
	e.InitEventBuffer(ctx)
	ctx = context.WithValue(ctx, CtxEventDataKey, &e)

	return ctx, nil
}

func (p *Plugin) ProcessEnd(ctx context.Context) {
	e := getEventBuffer(ctx)
	if e == nil {
		slog.ErrorCtx(ctx, "getEventBuffer err")

		return
	}

	e.ResetEventBuffer()
}

func (p *Plugin) PostCallSuccess(ctx context.Context) {
	e := getEventBuffer(ctx)
	if e == nil {
		slog.ErrorCtx(ctx, "getEventBuffer err")

		return
	}

	if err := e.ReallyDispatchAllEvent(ctx); err != nil {
		slog.ErrorCtx(ctx, "ReallyDispatchAllEvent ", "err", err.Error())
	}
}
