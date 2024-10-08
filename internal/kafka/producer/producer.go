package main

import (
	"fmt"
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
	} else {
		log.Info().Msg(fmt.Sprintf("Produced e message: %s", string(data)))
	}
	p.Flush(-1)
}

func main() {
	logger.InitLogger()
	filename := ""
	if len(os.Args) > 1 {
		filename = os.Args[len(os.Args)-1]
	} else {
		filename = "model.json"
	}
	data, err := os.ReadFile("testdata/" + filename)
	if err != nil {
		log.Err(err)
	}
	StartProducer(data)
}
