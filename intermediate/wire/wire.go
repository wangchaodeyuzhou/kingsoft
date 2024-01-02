//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package wire

import "github.com/google/wire"

func InitializeEvent(msg string) Event {
	wire.Build(NewMessage, NewGreeter, NewEvent)
	return Event{}
}
