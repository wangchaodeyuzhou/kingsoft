package wiretest

import "testing"

// DI 依赖注入
func TestDIWire(t *testing.T) {
	s := InitializeEvent()
	s.Start()
}
