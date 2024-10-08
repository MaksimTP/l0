package main

import (
	"main/internal/configs"
	"main/internal/logger"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func StartProducer(data []byte) {
	configs.InitConfig()
	conf := configs.GetConfig()

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": conf.Kafka.Server,
	})

	if err != nil {
		log.Err(err)
	}
	defer p.Close()
	topic := conf.Kafka.Topic
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          data,
	}, nil)

	if err != nil {
		log.Err(err)
	}
	p.Flush(-1)
}

func main() {
	logger.InitLogger()
	data, err := os.ReadFile("model.json")
	if err != nil {
		log.Err(err)
	}
	StartProducer(data)
}
