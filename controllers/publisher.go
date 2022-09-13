package controllers

import (
	"fmt"
	"time"
	"watermillTutorial/handlers"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

// PublishToKafka is a controller function to publish all message to kafka
func PublishToKafka(host, topic string) {
	publisher := handlers.PublisherHandler(host)
	i := 0
	for {
		msg := message.NewMessage(watermill.NewUUID(), []byte(fmt.Sprintf("Hello, world! [%d]", i)))

		if err := publisher.Publish(topic, msg); err != nil {
			panic(err)
		}

		i += 1
		time.Sleep(2 * time.Second)
	}
}
