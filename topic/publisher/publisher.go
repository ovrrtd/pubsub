package main

import (
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")

	exchange := []string{"topic-exchange"}
	routingKey := "topic.deliver"

	if err != nil {
		log.Fatalf("connection.open source: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("channel connection: %s", err)
	}
	defer ch.Close()

	if err != nil {
		log.Fatalf("queue binding: %s", err)
	}

	fmt.Println("publisher initialized")

	err = ch.Publish(exchange[0], routingKey, false, false, amqp.Publishing{
		Body:        []byte(fmt.Sprintf("topic rmq implementation %v", time.Now().String())),
		ContentType: "text/plain",
	})
	fmt.Println("publisher publish message")

	if err != nil {
		log.Fatalf("rmq publish: %s", err)
	}
}
