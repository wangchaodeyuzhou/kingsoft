package wire

import "fmt"

type Message struct {
	msg string
}

type Greeter struct {
	Message Message
}

type Event struct {
	Greeter Greeter
}

func NewEvent(g Greeter) Event {
	return Event{
		Greeter: g,
	}
}

func NewMessage(msg string) Message {
	return Message{
		msg: msg,
	}
}

func NewGreeter(m Message) Greeter {
	return Greeter{
		Message: m,
	}
}

func (e Event) Start() {
	msg := e.Greeter.GetGreetMessage()
	fmt.Println(msg)
}

func (e Greeter) GetGreetMessage() Message {
	return e.Message
}
