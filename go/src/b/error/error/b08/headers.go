package b08

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
