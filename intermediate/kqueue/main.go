package main

import (
	"git.kingsoft.go/intermediate/kqueue/conf"
	"git.kingsoft.go/intermediate/kqueue/manager"
	"git.kingsoft.go/intermediate/kqueue/server"
	"git.kingsoft.go/intermediate/kqueue/util"
	"github.com/gorilla/websocket"
	"golang.org/x/exp/slog"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func initConfig() {
	_, err := conf.LoadConfig()
	if err != nil {
		slog.Error("load configs fail", "err", err)
		return
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	os.Exit(realMain())
}

func realMain() (exitCode int) {
	initConfig()

	util.InitIDGenerator(uint16(1))
	manager.Mgr = manager.NewWorkerManager()
	go manager.Mgr.Run()

	up := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	wsServer := &server.WebSocketServer{
		Up:      up,
		Clients: make(map[*websocket.Conn]struct{}),
		M:       manager.Mgr,
	}

	wsServer.Start()
	return 0
}
