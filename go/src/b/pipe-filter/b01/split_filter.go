package b01

import (
	"errors"
	"strings"
)

// ErrSplitFilterWrongFormat var.
var ErrSplitFilterWrongFormat = errors.New("Input data should be string")

// SplitFilter struct.
type SplitFilter struct {
	delimiter string
}

// NewSplitFilter func.
func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter}
}

// Process func.
func (sf *SplitFilter) Process(req Request) (Response, error) {
	str, ok := req.(string)
	if !ok {
		return nil, ErrSplitFilterWrongFormat
	}
	parts := strings.Split(str, sf.delimiter)
	return parts, nil
}
