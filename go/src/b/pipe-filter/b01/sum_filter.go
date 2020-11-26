package b01

import "errors"

// ErrSumFilterWrongFormat var.
var ErrSumFilterWrongFormat = errors.New("Input data should be []int")

// SumFilter struct.
type SumFilter struct{}

// NewSumFilter func.
func NewSumFilter() *SumFilter {
	return &SumFilter{}
}

// Process func.
func (sf *SumFilter) Process(req Request) (Response, error) {
	elems, ok := req.([]int)
	if !ok {
		return nil, ErrSumFilterWrongFormat
	}
	ret := 0
	for _, elem := range elems {
		ret += elem
	}
	return ret, nil
}
