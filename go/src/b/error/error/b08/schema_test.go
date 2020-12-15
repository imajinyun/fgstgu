package b08

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchemaErrors(t *testing.T) {
	err := InvalidType("confirmed", "query", "boolean", nil)
	assert.Error(t, err)
	assert.EqualValues(t, InvalidTypeCode, err.Code())
	assert.Equal(t, "confirmed in query must be of type boolean", err.Error())

	err = InvalidType("confirmed", "", "boolean", nil)
	assert.Error(t, err)
	assert.EqualValues(t, InvalidTypeCode, err.Code())
	assert.Equal(t, "confirmed must be of type boolean", err.Error())

	err = InvalidType("confirmed", "query", "boolean", "hello")
	assert.Error(t, err)
	assert.EqualValues(t, InvalidTypeCode, err.Code())
	assert.Equal(t, "confirmed in query must be of type boolean: \"hello\"", err.Error())

	err = InvalidType("confirmed", "query", "boolean", errors.New("hello"))
	assert.Error(t, err)
	assert.EqualValues(t, InvalidTypeCode, err.Code())
	assert.Equal(t, "confirmed in query must be of type boolean, because: hello", err.Error())

	err = InvalidType("confirmed", "", "boolean", "hello")
	assert.Error(t, err)
	assert.EqualValues(t, InvalidTypeCode, err.Code())
	assert.EqualValues(t, "confirmed must be of type boolean: \"hello\"", err.Error())

	err = InvalidType("confirmed", "", "boolean", errors.New("hello"))
	assert.Error(t, err)
	assert.EqualValues(t, InvalidTypeCode, err.Code())
	assert.Equal(t, "confirmed must be of type boolean, because: hello", err.Error())

	err = DuplicateItems("uniques", "query")
	assert.Error(t, err)
	assert.EqualValues(t, UniqueFailCode, err.Code())
	assert.Equal(t, "uniques in query shouldn't contain duplicates", err.Error())
}
