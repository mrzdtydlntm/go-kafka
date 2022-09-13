package controllers

import (
	"context"
	"log"
	"watermillTutorial/handlers"

	"github.com/ThreeDotsLabs/watermill/message"
)

// SubscribeFromKafka is a controller function to subscribe all messages from kafka
func SubscribeFromKafka(host, consumerGroup, topic string) {
	subscriber := handlers.SubscriberHandler(host, consumerGroup)
	messages, err := subscriber.Subscribe(context.Background(), topic)
	if err != nil {
		panic(err)
	}

	go process(messages)
}

func process(messages <-chan *message.Message) {
	for msg := range messages {
		log.Printf("received message: %s, payload: %s", msg.UUID, string(msg.Payload))

		// we need to Acknowledge that we received and processed the message,
		// otherwise, it will be resent over and over again.
		msg.Ack()
	}
}
