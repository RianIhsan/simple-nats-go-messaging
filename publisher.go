package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Publish message to "updates" subject
	for i := 1; i <= 10; i++ {
		msg := fmt.Sprintf("Sangat Gacor %d", i)

		nc.Publish("updates", []byte(msg))
		fmt.Printf("Published: %s\n", msg)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("All messages published!")
}
