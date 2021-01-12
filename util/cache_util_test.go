package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUUID(t *testing.T) {
	uuid := GetUUID()
	assert.NotNil(t, uuid)
}
