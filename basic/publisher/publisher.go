package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")

	if err != nil {
		log.Fatalf("connection.open source: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("channel connection: %s", err)
	}
	defer ch.Close()

	_, err = ch.QueueDeclare("basic", true, false, false, false, nil)

	if err != nil {
		log.Fatalf("queue connection: %s", err)
	}

	err = ch.Publish("", "basic", false, false, amqp.Publishing{
		Body:        []byte("basic rmq implementation"),
		ContentType: "text/plain",
	})

	if err != nil {
		log.Fatalf("rmq publish: %s", err)
	}
}
