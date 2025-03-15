package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func submit(w http.ResponseWriter, r *http.Request) {
	message := r.PathValue("message")
	if message == "" {
		http.Error(w, "Message parameter is missing", http.StatusBadRequest)
		return
	}

	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	fmt.Println("Connecting to RabbitMQ")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	//Channel creation
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"publisher",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})

	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", message)
	w.WriteHeader(http.StatusAccepted)
}

func main() {
	mux := http.NewServeMux()
	fmt.Println("Route to publish message /publish/<message>")
	mux.HandleFunc("POST /publish/{message}", submit)

	fmt.Println("Running")
	err := http.ListenAndServe(":8080", mux)
	failOnError(err, "Failed to listen on port 8080")

}
