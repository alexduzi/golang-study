package main

import (
	"github.com/alexduzi/golang-study/eventutils/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	rabbitmq.Publish(ch, "Hello World, RabbitMQ!", "amq.direct")
}
