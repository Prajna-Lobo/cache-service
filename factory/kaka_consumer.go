package factory

import (
	"cache-service/configuration"
	"cache-service/repository"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/patrickmn/go-cache"
	"log"
	"os"
	"os/signal"
)

func StartKafkaConsumer(config *configuration.ConfigData) {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Consumer.Return.Errors = true
	brokers := []string{config.Kafka.Broker}

	kafkaConsumer, err := sarama.NewConsumer(brokers, kafkaConfig)
	if err != nil {
		log.Fatal(err)
	}

	topic := config.Kafka.Topic
	partitionConsumer, err := kafkaConsumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatal(err)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	msgCount := 0
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-partitionConsumer.Errors():
				fmt.Println(err)
			case msg := <-partitionConsumer.Messages():
				msgCount++
				fmt.Println("Received messages", string(msg.Key), string(msg.Value))
				mongoCollection := GetMongoCollection(config)
				storageRepository := repository.NewStorageRepository(mongoCollection)
				var c cache.Cache
				_ = storageRepository.ReloadCacheFromDb(&c)
			case <-signals:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()
}
