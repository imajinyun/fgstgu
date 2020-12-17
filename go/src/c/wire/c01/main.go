package main

import "fmt"

// Message type.
type Message string

// Greeter struct.
type Greeter struct {
	Message Message
}

// Event struct.
type Event struct {
	Greeter Greeter
}

// Greet returns a Greeter.
func (g Greeter) Greet() Message {
	return g.Message
}

// Start returns a Event.
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

// NewMessage returns a Message.
func NewMessage() Message {
	return Message("🎉 Hi, VIP!")
}

// NewGreeter returns a Greeter.
func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

// NewEvent returns a Event.
func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

func main() {
	message := NewMessage()
	greeter := NewGreeter(message)
	event := NewEvent(greeter)
	event.Start()
}
