package controllers

import (
	"context"
	"watermillTutorial/handlers"

	"github.com/ThreeDotsLabs/watermill/message"
)

// SubscribeFromKafka is a controller function to subscribe all messages from kafka
func SubscribeFromKafka(host, consumerGroup string) <-chan *message.Message {
	subscriber := handlers.SubscriberHandler(host, consumerGroup)
	messages, err := subscriber.Subscribe(context.Background(), "allobank")
	if err != nil {
		panic(err)
	}

	return messages
}
