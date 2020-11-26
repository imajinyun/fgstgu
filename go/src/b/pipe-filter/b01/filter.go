package b01

// Request is the input of the filter.
type Request interface{}

// Response is the output of the filter.
type Response interface{}

// Filter interface is the definition of the data processing components.
type Filter interface {
	Process(request Request) (Response, error)
}
