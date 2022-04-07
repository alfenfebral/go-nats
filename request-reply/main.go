package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	wait := make(chan bool)
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Fatalln(err)
	}

	defer nc.Close()

	nc.Subscribe("foo", func(msg *nats.Msg) {
		log.Println("Request received:", string(msg.Data))

		nc.Publish(msg.Reply, []byte("I'm here"))
		msg.Respond([]byte("I'm here"))
	})

	<-wait
}
