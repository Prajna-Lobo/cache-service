package factory

import (
	"cache-service/configuration"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCache(t *testing.T) {
	cache, err := GetCache()
	assert.NotNil(t, cache)
	assert.Nil(t, err)
}

func TestGetMongoCollection(t *testing.T) {
	config := configuration.ConfigData{}
	collection := GetMongoCollection(&config)
	assert.NotNil(t, collection)
}

func TestGetDbSession(t *testing.T) {
	config := configuration.ConfigData{}
	db, err := GetDbSession(&config)
	assert.NotNil(t, db)
	assert.Nil(t, err)
}
