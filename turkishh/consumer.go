package turkishh

import (
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"runtime"
)


func Consumer(queue amqp.Queue, channel *amqp.Channel,collection *mongo.Collection ) {

	msgs, err := channel.Consume(
		queue.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	FailOnError(err, "Failed to register a consumer")

    runtime.GOMAXPROCS(8)
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			go Save(string([]byte(d.Body)),collection)

		}
	}()
   <-forever

}


