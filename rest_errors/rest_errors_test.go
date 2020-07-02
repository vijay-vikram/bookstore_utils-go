package rest_errors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewInternalServerError(t *testing.T) {
	restErr := NewInternalServerError("this is the message", errors.New("database error"))
	assert.NotNil(t, restErr)
	assert.EqualValues(t, http.StatusInternalServerError, restErr.Status)
	assert.EqualValues(t, "internal_server_error", restErr.Error)
	assert.EqualValues(t, "this is the message", restErr.Message)

	assert.NotNil(t, restErr.Causes)
	assert.EqualValues(t, 1, len(restErr.Causes))
	assert.EqualValues(t, "database error", restErr.Causes[0])
}

func TestNewBadRequestError(t *testing.T) {
	restErr := NewBadRequestError("this is the message")
	assert.NotNil(t, restErr)
	assert.EqualValues(t, http.StatusBadRequest, restErr.Status)
	assert.EqualValues(t, "bad_request", restErr.Error)
	assert.EqualValues(t, "this is the message", restErr.Message)
}

func TestNewNotFoundError(t *testing.T) {
	restErr := NewNotFoundError("this is the message")
	assert.NotNil(t, restErr)
	assert.EqualValues(t, http.StatusNotFound, restErr.Status)
	assert.EqualValues(t, "not_found", restErr.Error)
	assert.EqualValues(t, "this is the message", restErr.Message)
}

func TestNewError(t *testing.T) {
	err := NewError("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, "this is the message", err.Error())
}
