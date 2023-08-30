package state

import (
	"context"
	"google.golang.org/grpc/stats"
	"log/slog"
)

type ClientHandler struct {
	BaseState
	CallBack ClosedCallBack
}

func (h *ClientHandler) HandleConn(ctx context.Context, s stats.ConnStats) {
	iInfo := ctx.Value(connCtxKey{})

	stat, ok := iInfo.(*stats.ConnTagInfo)
	if !ok {
		return
	}

	if _, ok := s.(*stats.ConnEnd); ok {
		slog.Info("handle client conn closed", "addr", stat.RemoteAddr.String())

		if h.CallBack != nil {
			h.CallBack(ctx, stat.RemoteAddr.String())
		} else {
			slog.Info("close call back is nil")
		}
	}

	if _, ok := s.(*stats.ConnBegin); ok {
		slog.Info("handle client conn ", "addr", stat.RemoteAddr.String())
	}

}
