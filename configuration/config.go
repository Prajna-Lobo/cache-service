package configuration

import (
	"encoding/json"
	"os"
)

type Config interface {
	LoadConfig(fileName string)
}

type ConfigData struct {
	Database Database `json:"database"`
	Kafka    Kafka    `json:"kafka"`
}
type Database struct {
	Host                string `json:"db_host"`
	Name                string `json:"db_name"`
	Username            string `json:"db_username"`
	Password            string `json:"db_password"`
	MongoCollectionName string `json:"mongo_collection_name"`
}

type Kafka struct {
	Broker string `json:"broker"`
	Topic string `json:"topic"`
}

func LoadConfig(fileName string) (*ConfigData, error) {
	configuration := ConfigData{}
	file, err := os.Open(fileName)
	if err != nil {
		return &configuration, err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	return &configuration, err
}
