package state

import (
	"context"
	"google.golang.org/grpc/stats"
)

type ServerHandler struct {
	BaseState
	CallBack ClosedCallBack
}

func (h *ServerHandler) HandleConn(ctx context.Context, s stats.ConnStats) {
	iInfo := ctx.Value(connCtxKey{})

	stat, ok := iInfo.(*stats.ConnTagInfo)
	if !ok {
		return
	}

	if _, ok := s.(*stats.ConnEnd); ok {
		if h.CallBack != nil {
			h.CallBack(ctx, stat.RemoteAddr.String())
		}
	}
}
