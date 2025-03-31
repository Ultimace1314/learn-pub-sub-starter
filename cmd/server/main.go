package main

import (
	"fmt"
	"log"

	"github.com/Ultimace1314/learn-pub-sub-starter/internal/pubsub"
	"github.com/Ultimace1314/learn-pub-sub-starter/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril server...")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ")
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Peril game server connected to RabbitMQ!")

	publishCh, err := conn.Channel()
	if err != nil {
		log.Fatalf("could not create channel: %v", err)
	}

	err = pubsub.PublishJSON(
		publishCh,
		routing.ExchangePerilDirect,
		routing.PauseKey,
		routing.PlayingState{
			IsPaused: true,
		},
	)
	if err != nil {
		log.Printf("could not publish time: %v", err)
	}
	fmt.Println("Pause message sent!")
}
