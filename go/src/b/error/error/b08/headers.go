package b08

import (
	"fmt"
	"net/http"
)

type Validation struct {
	code    int32
	Name    string
	In      string
	Value   interface{}
	message string
	Values  []interface{}
}

// Error is return Validation error message.
func (v *Validation) Error() string {
	return v.message
}

// Code is return Validation error code.
func (v *Validation) Code() int32 {
	return v.code
}

// ValidateName produces an error message name for an aliased property.
func (v *Validation) ValidateName(name string) *Validation {
	if v.Name == "" && name != "" {
		v.Name = name
		v.message = name + v.message
	}
	return v
}

const (
	contentTypeFail    = `unsupported media type %q, only %v are allowed`
	responseFormatFail = `unsupported media type requested, only %v are available`
)

// InvalidContentType error for an invalid content type.
func InvalidContentType(value string, allowed []string) *Validation {
	values := make([]interface{}, 0, len(allowed))
	for _, v := range allowed {
		values = append(values, v)
	}
	return &Validation{
		code:    http.StatusUnsupportedMediaType,
		Name:    "Content-Type",
		In:      "header",
		Value:   value,
		Values:  values,
		message: fmt.Sprintf(contentTypeFail, value, allowed),
	}
}

// InvalidResponseFormat error for an unacceptable response format request.
func InvalidResponseFormat(value string, allowed []string) *Validation {
	values := make([]interface{}, 0, len(allowed))
	for _, v := range allowed {
		values = append(values, v)
	}
	return &Validation{
		code:    http.StatusNotAcceptable,
		Name:    "Accept",
		In:      "header",
		Value:   value,
		Values:  values,
		message: fmt.Sprintf(responseFormatFail, allowed),
	}
}
