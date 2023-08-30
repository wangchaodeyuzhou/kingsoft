package state

import (
	"context"

	"google.golang.org/grpc/stats"
)

type connCtxKey struct{}

type ClosedCallBack func(context.Context, string)

type BaseState struct{}

func (h *BaseState) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	// TODO implement me
	panic("implement me")
}

func (h *BaseState) HandleRPC(ctx context.Context, rpcStats stats.RPCStats) {
	// TODO implement me
	panic("implement me")
}

func (h *BaseState) TagConn(ctx context.Context, info *stats.ConnTagInfo) context.Context {
	return context.WithValue(ctx, connCtxKey{}, info)
}
