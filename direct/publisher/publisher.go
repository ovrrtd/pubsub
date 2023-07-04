package main

import (
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")

	exchange := []string{"direct-exchange"}
	kind := "direct"
	queueName := []string{"direct"}
	routingKey := []string{"direct-1"}

	if err != nil {
		log.Fatalf("connection.open source: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("channel connection: %s", err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(exchange[0], kind, true, false, false, false, nil)

	if err != nil {
		log.Fatalf("queue connection: %s", err)
	}

	q, err := ch.QueueDeclare(queueName[0], true, false, false, false, nil)

	if err != nil {
		log.Fatalf("queue connection: %s", err)
	}

	err = ch.QueueBind(q.Name, routingKey[0], exchange[0], false, nil)

	if err != nil {
		log.Fatalf("queue binding: %s", err)
	}

	fmt.Println("publisher initialized")

	err = ch.Publish(exchange[0], routingKey[0], false, false, amqp.Publishing{
		Body:        []byte(fmt.Sprintf("direct rmq implementation %v", time.Now().String())),
		ContentType: "text/plain",
	})
	fmt.Println("publisher publish message")

	if err != nil {
		log.Fatalf("rmq publish: %s", err)
	}
}
