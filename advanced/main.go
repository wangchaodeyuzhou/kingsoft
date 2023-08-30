package main

import (
	"go.kingsoft.com/advanced/consul"
	"log/slog"
)

func main() {
	slog.Info("llll wo shi log debug", "time", 1)
	consul.FmtP()
}
