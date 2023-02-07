package main

import "github.com/go-fc/08-eventos/rabbitMQ"

func main() {
	ch, err := rabbitMQ.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	rabbitMQ.Publish(ch, "Hello World!", "amq.direct")
}
