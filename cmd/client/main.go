package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/Ultimace1314/learn-pub-sub-starter/internal/gamelogic"
	"github.com/Ultimace1314/learn-pub-sub-starter/internal/pubsub"
	"github.com/Ultimace1314/learn-pub-sub-starter/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril client...")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ")
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Peril game server connected to RabbitMQ!")

	username, err := gamelogic.ClientWelcome()
	if err != nil {
		fmt.Println("Failed to get username")
		os.Exit(1)
	}

	_, queue, err := pubsub.DeclareAndBind(
		conn,
		routing.ExchangePerilDirect,
		routing.PauseKey+"."+username,
		routing.PauseKey,
		pubsub.SimpleQueueTransient,
	)
	if err != nil {
		log.Fatalf("could not subscribe to pause: %v", err)
	}
	fmt.Printf("Queue %v declared and bound!\n", queue.Name)

	// wait for ctrl+c
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
}
