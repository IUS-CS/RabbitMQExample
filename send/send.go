package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"rmq/common"
	"rmq/person"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
)

var ch *amqp.Channel
var q *amqp.Queue
var err error

func main() {
	_, ch, q, err = common.Connect("amqp://guest:guest@localhost:5672/")
	checkErr(err)
	http.HandleFunc("/", handleRequest)
	log.Printf("Listening on http://127.0.0.1:8080")
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type requestBody struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	rb := requestBody{}
	rawBody := json.NewDecoder(r.Body)
	err := rawBody.Decode(&rb)
	if err != nil {
		log.Printf("Error decoding request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	p := person.Person{}
	p.SetName(rb.Name)
	p.SetId(uuid.New().ID())
	p.SetEmail(rb.Email)

	body, err := proto.Marshal(&p)
	checkErr(err)

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/protobuf",
			Body:        body,
		})
	log.Printf("sent %s\n", p.String())
}
