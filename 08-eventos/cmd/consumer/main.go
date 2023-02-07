package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/raul-franca/go-fc/08-eventos/rabbitMQ"
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
