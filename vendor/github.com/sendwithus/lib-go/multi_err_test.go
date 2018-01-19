package swu

import (
	"testing"
	"errors"
	"github.com/stretchr/testify/assert"
)

func TestMultiError_Error(t *testing.T) {
	err := MultiError{errors.New("err1")}
	assert.Equal(t, "err1", err.Error())
}

func TestMultiError_ErrorList(t *testing.T) {
	err := MultiError{errors.New("err1"), errors.New("err2")}
	assert.Equal(t, "err1, err2", err.Error())
}

func TestMultiError_Append(t *testing.T) {
	err2 := errors.New("err2")
	err := MultiError{errors.New("err1")}
	assert.Equal(t, "err1, err2", err.Append(err2).Error())
}
