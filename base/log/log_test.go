package log

import "testing"

func TestAdd(t *testing.T) {
	Trace.Println("i have something to say")
	Info.Println("Special Information")
	Warning.Println("There is something you need to known about")
	Error.Println("Something has failed")
}
