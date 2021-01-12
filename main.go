package main

import (
	"cache-service/configuration"
	"cache-service/constants"
	"cache-service/factory"
	"cache-service/router"
	"log"
	"net/http"
)

func main() {
	configData, configErr := configuration.LoadConfig(constants.ConfigFileName)
	if configErr != nil {
		log.Fatal("could not read config file ", configErr)
	}
	factory.StartKafkaConsumer(configData)
	ginRouter := router.SetUpRouter(*configData)

	_ = http.ListenAndServe("0.0.0.0:8081", ginRouter)
}
