package b08

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type customError struct {
	apiError
}

func TestServeError(t *testing.T) {
	err := MethodNotAllowed("GET", []string{"POST", "PUT"})
	recorder := httptest.NewRecorder()
	ServeError(recorder, nil, err)
	assert.Equal(t, http.StatusMethodNotAllowed, recorder.Code)
	assert.Equal(t, "POST,PUT", recorder.Header().Get("Allow"))
	assert.Equal(t, "application/json", recorder.Header().Get("content-type"))
	assert.Equal(t, `{"code":405,"message":"method GET is not allowed, but [POST,PUT] are"}`, recorder.Body.String())

	err = NotFound("")
	recorder = httptest.NewRecorder()
	ServeError(recorder, nil, err)
	assert.Equal(t, http.StatusNotFound, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("content-type"))
	assert.Equal(t, `{"code":404,"message":"Not found"}`, recorder.Body.String())

	err = InvalidTypeName("someType")
	recorder = httptest.NewRecorder()
	ServeError(recorder, nil, err)
	assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("content-type"))
	assert.Equal(t, `{"code":601,"message":"someType is an invalid type name"}`, recorder.Body.String())

	func() {
		oldDefaultHTTPCode := DefaultHTTPCode
		defer func() { DefaultHTTPCode = oldDefaultHTTPCode }()
		DefaultHTTPCode = http.StatusBadRequest

		err = InvalidTypeName("someType")
		recorder = httptest.NewRecorder()
		ServeError(recorder, nil, err)
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		assert.Equal(t, "application/json", recorder.Header().Get("content-type"))
		assert.Equal(t, `{"code":601,"message":"someType is an invalid type name"}`, recorder.Body.String())
	}()

	simpleErr := fmt.Errorf("server internal error")
	recorder = httptest.NewRecorder()
	ServeError(recorder, nil, simpleErr)
	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("content-type"))
	assert.Equal(t, `{"code":500,"message":"server internal error"}`, recorder.Body.String())

	compositeErr := &CompositeError{
		Errors: []error{
			fmt.Errorf("first error"),
			fmt.Errorf("another error"),
		},
	}
	recorder = httptest.NewRecorder()
	ServeError(recorder, nil, compositeErr)
	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("content-type"))
	assert.Equal(t, `{"code":500,"message":"first error"}`, recorder.Body.String())

	compositeErr = &CompositeError{
		Errors: []error{
			New(600, "myApiError"),
			New(601, "MyOtherApiError"),
		},
	}
	recorder = httptest.NewRecorder()
	ServeError(recorder, nil, compositeErr)
	assert.Equal(t, CompositeErrorCode, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("content-type"))
	assert.Equal(t, `{"code":600,"message":"myApiError"}`, recorder.Body.String())

	compositeErr = &CompositeError{
		Errors: []error{
			&CompositeError{
				Errors: []error{
					New(600, "myApiError"),
					New(601, "MyOtherApiError"),
				},
			},
		},
	}
	recorder = httptest.NewRecorder()
	ServeError(recorder, nil, compositeErr)
	assert.Equal(t, CompositeErrorCode, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("content-type"))
	assert.Equal(t, `{"code":600,"message":"myApiError"}`, recorder.Body.String())

	compositeErr = &CompositeError{
		Errors: []error{
			&CompositeError{
				Errors: []error{},
			},
		},
	}
	recorder = httptest.NewRecorder()
	ServeError(recorder, nil, compositeErr)
	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("content-type"))
	assert.Equal(t, `{"code":500,"message":"Unknown error"}`, recorder.Body.String())

	recorder = httptest.NewRecorder()
	ServeError(recorder, nil, nil)
	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("content-type"))
	assert.Equal(t, `{"code":500,"message":"Unknown error"}`, recorder.Body.String())

	recorder = httptest.NewRecorder()
	var foo *customError
	ServeError(recorder, nil, foo)
	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("content-type"))
	assert.Equal(t, `{"code":500,"message":"Unknown error"}`, recorder.Body.String())
}

func TestAPIErrors(t *testing.T) {
	err := New(402, "this failed %s", "test")
	assert.Error(t, err)
	assert.EqualValues(t, 402, err.Code())
	assert.EqualValues(t, "this failed test", err.Error())

	err = NotFound("this failed %d", 1)
	assert.Error(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Code())
	assert.EqualValues(t, "this failed 1", err.Error())

	err = NotFound("")
	assert.Error(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Code())
	assert.EqualValues(t, "Not found", err.Error())

	err = NotImplemented("not implemented")
	assert.Error(t, err)
	assert.EqualValues(t, http.StatusNotImplemented, err.Code())
	assert.EqualValues(t, "not implemented", err.Error())

	err = MethodNotAllowed("GET", []string{"POST", "PUT"})
	assert.Error(t, err)
	assert.EqualValues(t, http.StatusMethodNotAllowed, err.Code())
	assert.EqualValues(t, "method GET is not allowed, but [POST,PUT] are", err.Error())

	err = InvalidContentType("application/test", []string{"application/json", "application/xml"})
	assert.Error(t, err)
	assert.EqualValues(t, http.StatusUnsupportedMediaType, err.Code())
	assert.EqualValues(t, "unsupported media type \"application/test\", only [application/json application/xml] are allowed", err.Error())

	err = InvalidResponseFormat("application/test", []string{"application/json", "application/xml"})
	assert.Error(t, err)
	assert.EqualValues(t, http.StatusNotAcceptable, err.Code())
	assert.EqualValues(t, "unsupported media type requested, only [application/json application/xml] are available", err.Error())
}

func TestValidateName(t *testing.T) {
	v := &Validation{Name: "myValidation", message: "myMessage"}

	vn := v.ValidateName("")
	assert.EqualValues(t, "myValidation", vn.Name)
	assert.EqualValues(t, "myMessage", vn.message)

	vn = v.ValidateName("myValidateName")
	assert.EqualValues(t, "myValidation", vn.Name)
	assert.EqualValues(t, "myMessage", vn.message)

	v.Name = ""

	vn = v.ValidateName("")
	assert.EqualValues(t, "", vn.Name)
	assert.EqualValues(t, "myMessage", vn.message)

	vn = v.ValidateName("myValidateName")
	assert.EqualValues(t, "myValidateName", vn.Name)
	assert.EqualValues(t, "myValidateNamemyMessage", vn.message)
}
