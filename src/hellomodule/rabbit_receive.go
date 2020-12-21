package hellomodule

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

// ReceiveMessages messages from amqp
func ReceiveMessages() {
	conn, err := amqp.Dial("amqps://zcoqmbte:2hLTwEGHEPy9iXbWz9NWSi5DZch8eq1a@flamingo.rmq.cloudamqp.com/zcoqmbte")
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			var bird Bird
			json.Unmarshal(d.Body, &bird)

			d.Ack(true)

			log.Printf("Received a message with Species: %s", bird.Species)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	failOnError(err, "Failed to publish a message")

	defer ch.Close()

	defer conn.Close()
}
