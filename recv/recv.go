package main

import (
	"log"
	"rmq/common"
	"rmq/person"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	_, ch, q, err := common.Connect("amqp://guest:guest@localhost:5672/")

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

	var forever chan struct{}

	go func() {
		for d := range msgs {
			p := person.Person{}
			err := proto.Unmarshal(d.Body, &p)
			failOnError(err, "Failed to unmarshal")
			id, err := uuid.Parse(p.GetId())
			failOnError(err, "Failed to parse id")
			log.Printf("Received a message: Name: %s, Email: %s, ID: %s",
				p.String(), p.GetEmail(), id)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
