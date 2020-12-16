package main

import (
	"fmt"
)

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

// Greet method for Greeter.
func (g Greeter) Greet() Message {
	return g.Message
}

// Start method for Event.
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

// NewMessage returns Message.
func NewMessage() Message {
	return Message("🎉 Hi, VIP")
}

// NewGreeter returns Greeter.
func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

// NewEvent returns Event.
func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

// Using wire generate.
func main() {
	e := InitializeEvent()
	e.Start()
}
