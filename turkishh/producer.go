package turkishh

import (
	"github.com/streadway/amqp"
)


func GetConnection(url string) *amqp.Connection {
	conn, err := amqp.Dial(url)
	FailOnError(err, "Failed to connect to RabbitMQ")
	return conn
}


func GetChannel(connection *amqp.Connection) *amqp.Channel {
	ch, err := connection.Channel()
	FailOnError(err, "Failed to open a channel")
	return ch
}

func DeclareTopic(name string, channel *amqp.Channel) amqp.Queue {
	q, err := channel.QueueDeclare(
		name, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	FailOnError(err, "Failed to declare a queue")
	return q
}

func Publish(message string,queue amqp.Queue,channel *amqp.Channel)  {
	var err error
	err = channel.Publish(
		"",     // exchange
		queue.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	FailOnError(err, "Failed to publish a message")
}