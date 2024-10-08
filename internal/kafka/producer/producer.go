package main

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	kafkaServer = "localhost:9092"
	kafkaTopic  = "orders"
)

func StartProducer(data []byte) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaServer,
	})

	if err != nil {
		log.Fatalln("couldnt start kafka producer", err)
	}
	defer p.Close()
	topic := kafkaTopic
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
