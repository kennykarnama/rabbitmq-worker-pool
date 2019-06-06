package main

import (
	"log"

	"github.com/kennykarnama/rabbitmq-worker-pool/config"

	"github.com/kennykarnama/rabbitmq-worker-pool/pool"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	cfg := config.Get()
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		cfg.QueueName, // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	p := pool.NewPool(cfg.NumWorkers)

	p.Run()

	go func() {
		var counter int32 = 1
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			p.JobQueue <- pool.Job{
				ID:        counter,
				Resources: string(d.Body),
			}
			counter++
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
