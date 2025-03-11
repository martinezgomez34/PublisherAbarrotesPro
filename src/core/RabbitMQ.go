package core

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Queue   amqp.Queue
}

func NewRabbitMQ(queueName string) *RabbitMQ {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatalf("Error connecting to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Error creating RabbitMQ channel: %v", err)
	}

	queue, err := ch.QueueDeclare(
		queueName,
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatalf("Error declaring RabbitMQ queue: %v", err)
	}

	return &RabbitMQ{
		Conn:    conn,  //itamar teomitzi cruz xd
		Channel: ch,
		Queue:   queue,
	}
}

func (r *RabbitMQ) PublishMessage(message []byte) error {
	err := r.Channel.Publish(
		"",           // exchange
		r.Queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
	if err != nil {
		log.Printf("Error publishing message to queue %s: %v", r.Queue.Name, err)
		return err
	}
	log.Printf("Message published to queue %s", r.Queue.Name)
	return nil
}

func (r *RabbitMQ) Close() {
	if err := r.Channel.Close(); err != nil {
		log.Printf("Error closing RabbitMQ channel: %v", err)
	}
	if err := r.Conn.Close(); err != nil {
		log.Printf("Error closing RabbitMQ connection: %v", err)
	}
}