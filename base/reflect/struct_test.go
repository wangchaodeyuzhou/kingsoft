package reflect

import "testing"

func TestCallGmMethod(t *testing.T) {
	CallGmMethod("Say")
	CallGmMethod("Hello")
}
