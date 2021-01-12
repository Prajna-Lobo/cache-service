package configuration

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig("./config.json")
	assert.NotNil(t, config)
	assert.Nil(t, err)

	assert.Equal(t, "nokia", config.Kafka.Topic)
	assert.Equal(t, "localhost:9092", config.Kafka.Broker)

	assert.Equal(t, "nokia", config.Database.MongoCollectionName)
	assert.Equal(t, "admin", config.Database.Name)
	assert.Equal(t, "localhost", config.Database.Host)
}

func TestLoadConfigShouldReturnError(t *testing.T) {
	_, err := LoadConfig("")
	assert.NotNil(t, err)
}
