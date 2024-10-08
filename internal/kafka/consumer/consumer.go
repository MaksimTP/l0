package consumer

import (
	"fmt"
	"main/internal/cache"
	"main/internal/configs"
	"main/internal/db"
	"main/internal/types"

	"github.com/rs/zerolog/log"

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
		log.Err(err)
	}
	defer c.Close()

	topic := conf.Kafka.Topic
	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		log.Err(err)
	} else {
		log.Debug().Msg("subscribed to the topic with a name " + topic)
	}

	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			log.Err(err)
			continue
		}
		order, err := types.ReadJSON(msg.Value)
		if err != nil {
			log.Err(err)
		} else {
			database.InsertData(order)
			cache_.SaveData(order)
		}
		log.Info().Msg(fmt.Sprintf("Received Order: %+v", order))
	}
}
