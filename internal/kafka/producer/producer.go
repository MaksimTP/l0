package main

import (
	"log"
	"main/internal/configs"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func StartProducer(data []byte) {
	configs.InitConfig()
	conf := configs.GetConfig()

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": conf.Kafka.Server,
	})

	if err != nil {
		log.Fatalln("couldnt start kafka producer", err)
	}
	defer p.Close()
	topic := conf.Kafka.Topic
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          data,
	}, nil)

	if err != nil {
		log.Fatalln("couldnt produce a message, err:", err)
	}
	p.Flush(-1)
}

func main() {
	data, err := os.ReadFile("model.json")
	if err != nil {
		log.Println(err)
	}
	StartProducer(data)
}
