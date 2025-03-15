package main

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	consume()
}

func consume() {
	connection, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("%s : %s", "failed to connect to rabbitmq", err)
	}

	ch, err := connection.Channel()
	if err != nil {
		log.Fatalf("%s : %s", "failed to open a channel", err)
	}

	q, err := ch.QueueDeclare(
		"publisher",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("%s : %s", "failed to declare queue", err)
	}

	fmt.Println("Queue and channel are declared successfully")

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("%s : %s", "failed to register a consumer", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			d.Ack(false)
		}
	}()

	fmt.Println("Running")

	<-forever

}
