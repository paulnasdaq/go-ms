package main

import (
	"github.com/rabbitmq/amqp091-go"
	"log"
)

func main() {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	q, err := ch.QueueDeclare("notifications", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	err = ch.QueueBind(q.Name, "test.*", "amq.topic", false, nil)
	if err != nil {
		panic(err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer tag
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	forever := make(chan bool)

	// Receive messages
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Println("Waiting for messages. To exit press CTRL+C")
	<-forever
}
