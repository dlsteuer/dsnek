package swu

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func init() {
	os.Setenv("DYNO", "svc-sendlog-event-consumers.1")
}

func TestGetDynoName(t *testing.T) {
	name, err := GetDynoAppName()
	assert.Nil(t, err)
	assert.Equal(t, "svc-sendlog-event-consumers", name)
}

func TestGetDynoCount(t *testing.T) {
	count, err := GetDynoCount()
	assert.Nil(t, err)
	assert.Equal(t, 0, count)
}

func TestGetDynoIndex(t *testing.T) {
	index, err := GetDynoIndex()
	assert.Nil(t, err)
	assert.Equal(t, 1, index)
}
