package b01

import (
	"errors"
	"strconv"
)

// ErrToIntFilterWrongFormat var.
var ErrToIntFilterWrongFormat = errors.New("Input data should be []string")

// ToIntFilter struct.
type ToIntFilter struct{}

// NewToIntFilter func.
func NewToIntFilter() *ToIntFilter {
	return &ToIntFilter{}
}

// Process func.
func (tif *ToIntFilter) Process(req Request) (Response, error) {
	parts, ok := req.([]string)
	if !ok {
		return nil, ErrToIntFilterWrongFormat
	}
	ret := []int{}
	for _, part := range parts {
		s, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		ret = append(ret, s)
	}
	return ret, nil
}
