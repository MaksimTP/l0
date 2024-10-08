package consumer

import (
	"log"
	"main/internal/cache"
	"main/internal/configs"
	"main/internal/db"
	"main/internal/types"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func StartConsumer(database db.DataBase, cache_ *cache.Cache) {

	conf := configs.GetConfig()

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": conf.Kafka.Server,
		"group.id":          conf.Kafka.GroupID,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	topic := conf.Kafka.Topic
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
