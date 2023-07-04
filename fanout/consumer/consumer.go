package main

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	exchange := []string{"fanout-exchange"}
	kind := "fanout"
	queueName := []string{"fanout"}

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

	err = ch.QueueBind(q.Name, "", exchange[0], false, nil)

	if err != nil {
		log.Fatalf("queue binding: %s", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("queue connection: %s", err)
	}

	fmt.Println("consumer initialized")

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			fmt.Printf("Received Message: %s\n", msg.Body)
		}
	}()

	<-forever
}
