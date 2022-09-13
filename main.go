// Sources for https://watermill.io/docs/getting-started/
package main

import (
	"log"
	"os"
	"watermillTutorial/controllers"

	"github.com/joho/godotenv"
)

func main() {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Fatalf("Error read env file with err: %s", errEnv)
	}
	var (
		host          string = os.Getenv("KAFKA_HOST")
		topic         string = os.Getenv("KAFKA_TOPIC")
		consumerGroup string = os.Getenv("KAFKA_CONSUMER_GROUP")
	)

	controllers.SubscribeFromKafka(host, consumerGroup, topic)
	controllers.PublishToKafka(host, topic)
}
