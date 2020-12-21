package hellomodule

import (
	"github.com/streadway/amqp"
)

// Initialize amqp
func Initialize() {
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

	body := "Hello World!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

	defer ch.Close()

	defer conn.Close()
}
