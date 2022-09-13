// Sources for https://watermill.io/docs/getting-started/
package main

import (
	"log"
	"watermillTutorial/controllers"

	"github.com/ThreeDotsLabs/watermill/message"
)

func main() {
	var (
		host          string = "34.66.147.156:9092"
		topic         string = "allobank"
		consumerGroup string = "allobank-group"
	)

	controllers.PublishToKafka(host, topic)

	go process(controllers.SubscribeFromKafka(host, consumerGroup))

}

func process(messages <-chan *message.Message) {
	for msg := range messages {
		log.Printf("received message: %s, payload: %s", msg.UUID, string(msg.Payload))

		// we need to Acknowledge that we received and processed the message,
		// otherwise, it will be resent over and over again.
		msg.Ack()
	}
}
