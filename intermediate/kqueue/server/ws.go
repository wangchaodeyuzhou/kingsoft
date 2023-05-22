package server

import (
	"encoding/json"
	"git.kingsoft.go/intermediate/kqueue/api"
	"git.kingsoft.go/intermediate/kqueue/manager"
	"git.kingsoft.go/intermediate/kqueue/request"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"golang.org/x/exp/slog"
	"time"
)

type WebSocketServer struct {
	Up      websocket.Upgrader
	Clients map[*websocket.Conn]struct{}
	M       *manager.Manager
}

func (ws *WebSocketServer) handleWebSocket(c *gin.Context) {
	conn, err := ws.Up.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		slog.Error("Failed to upgrade WebSocket connection:", "err", err)
		return
	}
	defer conn.Close()

	ws.Clients[conn] = struct{}{}
	defer delete(ws.Clients, conn)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		mData := request.ConvertManagerData(ws.M)
		data, err := json.Marshal(mData)
		if err != nil {
			slog.Error("Failed to marshal queue data:", "err", err)
			continue
		}

		if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
			slog.Error("Failed to send queue data:", "err", err)
			break
		}
	}
}

func (ws *WebSocketServer) Start() {
	router := gin.Default()
	router.GET("/queueInfo", func(c *gin.Context) {
		ws.handleWebSocket(c)
	})

	// 提交任务
	router.POST("/commit", api.CommitTaskToQueue)

	if err := router.Run(":13000"); err != nil {
		slog.Error("websocket start fail", "err", err)
		return
	}
}
