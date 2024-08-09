package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Subscribe to "updates" subject
	_, err = nc.Subscribe("updates", func(m *nats.Msg) {
		fmt.Printf("Received message: %s\n", string(m.Data))
	})
	if err != nil {
		log.Fatal(err)
	}

	// Keep the connection alive
	select {}
}
