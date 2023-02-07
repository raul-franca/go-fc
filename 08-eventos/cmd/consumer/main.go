package main

import (
	"fmt"
	"github.com/go-fc/08-eventos/rabbitMQ"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	//ch channel do rabbitMQ
	ch, err := rabbitMQ.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	msgs := make(chan amqp.Delivery)

	go rabbitMQ.Consume(ch, msgs, "orders")

	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false) //Avisa para o rabbitMQ que o que a mensagem foi lida.
	}
}
