package main

import "fmt"

// MyEvent struct.
type MyEvent struct{}

func (e *MyEvent) devEvent(param interface{}) {
	fmt.Println("Dev event:", param)
}

func proEvent(param interface{}) {
	fmt.Println("Pro event:", param)
}

var eventByNames = make(map[string][]func(interface{}))

func main() {
	e := new(MyEvent)
	RegisterEvent("MyEvents", e.devEvent)
	RegisterEvent("MyEvents", proEvent)
	CallEvent("MyEvents", "The event is fire!")
}

// RegisterEvent func.
func RegisterEvent(name string, callback func(interface{})) {
	list := eventByNames[name]
	list = append(list, callback)
	eventByNames[name] = list
}

// CallEvent func.
func CallEvent(name string, param interface{}) {
	list := eventByNames[name]
	for _, callback := range list {
		callback(param)
	}
}
