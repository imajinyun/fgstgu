package b08

import (
	"fmt"
	"strings"
)

const (
	invalidType               = "%s is an invalid type name"
	typeFail                  = "%s in %s must be of type %s"
	typeFailWithData          = "%s in %s must be of type %s: %q"
	typeFailWithError         = "%s in %s must be of type %s, because: %s"
	requiredFail              = "%s in %s is required"
	readOnlyFail              = "%s in %s is readOnly"
	tooLongMessage            = "%s in %s should be at most %d chars long"
	tooShortMessage           = "%s in %s should be at least %d chars long"
	patternFail               = "%s in %s should match '%s'"
	enumFail                  = "%s in %s should be one of %v"
	multipleOfFail            = "%s in %s should be a multiple of %v"
	maxIncFail                = "%s in %s should be less than or equal to %v"
	maxExcFail                = "%s in %s should be less than %v"
	minIncFail                = "%s in %s should be greater than or equal to %v"
	minExcFail                = "%s in %s should be greater than %v"
	uniqueFail                = "%s in %s shouldn't contain duplicates"
	maxItemsFail              = "%s in %s should have at most %d items"
	minItemsFail              = "%s in %s should have at least %d items"
	typeFailNoIn              = "%s must be of type %s"
	typeFailWithDataNoIn      = "%s must be of type %s: %q"
	typeFailWithErrorNoIn     = "%s must be of type %s, because: %s"
	requiredFailNoIn          = "%s is required"
	readOnlyFailNoIn          = "%s is readOnly"
	tooLongMessageNoIn        = "%s should be at most %d chars long"
	tooShortMessageNoIn       = "%s should be at least %d chars long"
	patternFailNoIn           = "%s should match '%s'"
	enumFailNoIn              = "%s should be one of %v"
	multipleOfFailNoIn        = "%s should be a multiple of %v"
	maxIncFailNoIn            = "%s should be less than or equal to %v"
	maxExcFailNoIn            = "%s should be less than %v"
	minIncFailNoIn            = "%s should be greater than or equal to %v"
	minExcFailNoIn            = "%s should be greater than %v"
	uniqueFailNoIn            = "%s shouldn't contain duplicates"
	maxItemsFailNoIn          = "%s should have at most %d items"
	minItemsFailNoIn          = "%s should have at least %d items"
	noAdditionalItems         = "%s in %s can't have additional items"
	noAdditionalItemsNoIn     = "%s can't have additional items"
	tooFewProperties          = "%s in %s should have at least %d properties"
	tooFewPropertiesNoIn      = "%s should have at least %d properties"
	tooManyProperties         = "%s in %s should have at most %d properties"
	tooManyPropertiesNoIn     = "%s should have at most %d properties"
	unallowedProperty         = "%s.%s in %s is a forbidden property"
	unallowedPropertyNoIn     = "%s.%s is a forbidden property"
	failedAllPatternProps     = "%s.%s in %s failed all pattern properties"
	failedAllPatternPropsNoIn = "%s.%s failed all pattern properties"
	multipleOfMustBePositive  = "factor MultipleOf declared for %s must be positive: %v"
)

// All code responses can be used to differentiate errors for different handling
// by the consuming program
const (
	// CompositeErrorCode remains 422 for backwards-compatibility
	// and to separate it from validation errors with cause
	CompositeErrorCode = 422
	// InvalidTypeCode is used for any subclass of invalid types
	InvalidTypeCode = 600 + iota
	RequiredFailCode
	TooLongFailCode
	TooShortFailCode
	PatternFailCode
	EnumFailCode
	MultipleOfFailCode
	MaxFailCode
	MinFailCode
	UniqueFailCode
	MaxItemsFailCode
	MinItemsFailCode
	NoAdditionalItemsCode
	TooFewPropertiesCode
	TooManyPropertiesCode
	UnallowedPropertyCode
	FailedAllPatternPropsCode
	MultipleOfMustBePositiveCode
	ReadOnlyFailCode
)

// CompositeError is an error that groups several errors together.
type CompositeError struct {
	Errors  []error
	code    int32
	message string
}

// Error is return this error message.
func (c *CompositeError) Error() string {
	if len(c.Errors) > 0 {
		msgs := []string{c.message + ":"}
		for _, e := range c.Errors {
			msgs = append(msgs, e.Error())
		}
		return strings.Join(msgs, "\n")
	}
	return c.message
}

// Code is return this error code.
func (c *CompositeError) Code() int32 {
	return c.code
}

// CompositeValidationError an error to wrap a bunch of other errors.
func CompositeValidationError(errors ...error) *CompositeError {
	return &CompositeError{
		code:    CompositeErrorCode,
		Errors:  append([]error{}, errors...),
		message: "validation failure list",
	}
}

// InvalidTypeName an error for when the type is invalid.
func InvalidTypeName(typeName string) *Validation {
	return &Validation{
		code: InvalidTypeCode,
		Value: typeName,
		message: fmt.Sprintf(invalidType, typeName),
	}
}
