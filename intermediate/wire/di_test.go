package wire

import "testing"

func TestWire(t *testing.T) {
	event := InitializeEvent("hello word")
	event.Start()
}
