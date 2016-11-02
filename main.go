package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	go client()
	go server()

	var a string
	fmt.Scanln(&a)
}

func client() {
	conn, ch, q := getQueue()
	defer conn.Close()
	defer ch.Close()

	msgs, err := ch.Consume(
		q.Name, //queue string
		"",     //consumer string
		true,
		false,
		false,
		false,
		nil)

	failOnError(err, "Failed to register a consumer")

	for msg := range msgs {
		log.Printf("Recieved message with message: %s", msg.Body)
	}
}

func server() {
	conn, ch, q := getQueue()
	defer conn.Close()
	defer ch.Close()

	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Hello RabbitMQ"),
	}
	for {
		ch.Publish("", q.Name, false, false, msg)
	}
}

func getQueue() (*amqp.Connection, *amqp.Channel, *amqp.Queue) {
	conn, err := amqp.Dial("amqp://guest@localhost:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	q, err := ch.QueueDeclare("hello",
		false, //durable bool - this saves the msgs to disk to avoid loss if server re starts
		false, //autoDelete bool - msgs will be deleted if their is no actuve consumer otherwise msgs will be kept around until consumer comes around to recieve it
		false, //exclusive bool - queue will be exclusive to the connection which is used to set it.
		false, //noWait bool - it will return only a pre existing queue that matches the provided configuration, if non matched it will return error
		nil)   //args ampq.Table - args to be used in multiple cases like if performing header matching etc, since I am not using it our queue will be automatically bound to the default exchange
	failOnError(err, "Failed to declare a queue")

	return conn, ch, &q
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
