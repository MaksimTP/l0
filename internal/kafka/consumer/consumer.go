package consumer

import (
	"log"
	"main/internal/cache"
	"main/internal/db"
	"main/internal/types"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	kafkaServer  = "localhost:9092"
	kafkaTopic   = "orders"
	kafkaGroupId = "wb_product_service"
)

func StartConsumer(database db.DataBase, cache_ *cache.Cache) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaServer,
		"group.id":          kafkaGroupId,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	topic := kafkaTopic
	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		log.Println("couldnt subscribe to the topic with a name", topic)
	} else {
		log.Println("subscribed to the topic with a name", topic)
	}

	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			log.Println(err)
			continue
		}
		order, err := types.ReadJSON(msg.Value)
		if err != nil {
			log.Println("Error decoding message,", err)
		} else {
			database.InsertData(order)
			cache_.SaveData(order)
		}
		log.Printf("Received Order: %+v\n", order)
	}
}
