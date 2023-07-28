package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// To catch the message from the terminal
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the message to be sent?")
	messagePayload, _ := reader.ReadString('\n')

	// To setup a connection with RabbitMQ
	// 5672 is the RabbitMQ port
	// 15672 is the management API port
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// To create a channel in order to process the messages
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// To declare a queue
	// If the queue does not exist, it's created
	queueName := "simpleQueue"
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// We set the payload for the message and publish the message to the queue
	body := messagePayload
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

	log.Printf("Sent message: %s", body)
}
