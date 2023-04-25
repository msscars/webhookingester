package main

import (
	"log"
	"moscars-webhookingester-publisher/shared"
	"os"
	"sync"

	"github.com/nats-io/nats.go"
)

func main() {
	natsUrl, natsUrlFound := os.LookupEnv("WEBHOOKINGESTER_NATSURL")

	if !natsUrlFound {
		log.Println("No NATS url configured using default one")
		natsUrl = nats.DefaultURL
	}

	var config Config
	config.GetConf()

	routings := CreateRoutings(config)

	nc, err := nats.Connect(natsUrl)
	if err != nil {
		log.Fatal(err)
	}

	defer nc.Close()

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}
	defer ec.Close()

	// Use a WaitGroup to wait for 10 messages to arrive
	wg := sync.WaitGroup{}
	wg.Add(10)

	if _, err := ec.QueueSubscribe("publish-queue", "publishers", func(m *shared.IncommingWebhook) {
		log.Println("Message received:", m)

		Route(routings, m)

		log.Println("Message routed")

		wg.Done()
	}); err != nil {
		log.Println(err)
	}

	// Wait for messages to come in
	log.Println("Waiting for messages to come in")
	wg.Wait()
	log.Println("WaitGroup exceeded")
}
