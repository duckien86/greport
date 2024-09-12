package workqueues

import (
	"context"
	"fmt"
	"greport/common"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func createCnn() *amqp.Connection {
	user := common.GetEnv(common.RabbitmqUsername, "guest")
	pwd := common.GetEnv(common.RabbitmqPassword, "guest")
	host := common.GetEnv(common.RabbitmqHost, "localhost")
	port := common.GetEnv(common.RabbitmqPort, "5672")
	uri := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, pwd, host, port)
	conn, err := amqp.Dial(uri)
	failOnError(err, "Failed to connect to RabbitMQ")
	return conn
}

func createChannel(conn *amqp.Connection) *amqp.Channel {
	channel, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	return channel
}

func Publish(qName string, body string) {
	//`Dial` method creates a new connection to the RabbitMQ server.
	conn := createCnn()
	defer conn.Close()
	//`Channel` method creates a new channel on the connection.
	channel := createChannel(conn)
	defer channel.Close()
	//`QueueDeclare` declares a queue to hold messages and deliver to consumers.
	q, err := channel.QueueDeclare(qName, true, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")
	// declare a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//`Publishing` struct is used to publish a message.
	publishData := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         []byte(body),
	}
	//`Publish` method publishes a message to the queue.
	err = channel.PublishWithContext(ctx, "", q.Name, false, false, publishData)
	failOnError(err, "Failed to publish a message")
	log.Printf("[x] Sent %s", body)
}

func StartConsumer(queueList ...string) {
	//`Dial` method creates a new connection to the RabbitMQ server.
	conn := createCnn()
	defer conn.Close()
	//`Channel` method creates a new channel on the connection.
	channel := createChannel(conn)
	defer channel.Close()

	for _, queue := range queueList {
		//`QueueDeclare` declares a queue to hold messages and deliver to consumers.
		q, err := channel.QueueDeclare(queue, true, false, false, false, nil)
		failOnError(err, "Failed to declare a queue")
		//`Consume` method returns a channel in which the messages are delivered.
		msgs, err := channel.Consume(q.Name, "", true, false, false, false, nil)
		failOnError(err, "Failed to register a consumer")
		// start a goroutine to consume a message
		go func() {
			for delivery := range msgs {
				log.Printf("[x] Received a message: %s", delivery.Body)
			}
		}()
	}

	log.Printf("[*] Waiting for messages.")
	<-make(chan bool) // wait forever
}
