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
		Errors: []error {
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
		Errors: []error {
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
		Errors: []error {
			&CompositeError{
				Errors: []error {
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
