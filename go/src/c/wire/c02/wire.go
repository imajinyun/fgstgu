//+build wireinject

package main

import "github.com/google/wire"

// InitializeEvent returns a Event.
func InitializeEvent() Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}
