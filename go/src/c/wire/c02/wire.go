package main

import "github.com/google/wire"

// InitializeEvent func.
func InitializeEvent() Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}
