package main

import (
	"bufio"
	"bytes"
	"log"
	"net/http"
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

	// Use a WaitGroup to wait for 10 messages to arrive
	wg := sync.WaitGroup{}
	wg.Add(10)

	// Create a queue subscription on "updates" with queue name "workers"
	if _, err := nc.QueueSubscribe("publish-queue", "publishers", func(m *nats.Msg) {
		var b = &bytes.Buffer{}
		var originalRequest *http.Request
		r := bufio.NewReader(b)
		r.Read(m.Data)
		if originalRequest, err = http.ReadRequest(r); err != nil { // deserialize request
			log.Println(err)
		}

		Route(routings, originalRequest)

		wg.Done()
	}); err != nil {
		log.Fatal(err)
	}

	// Wait for messages to come in
	wg.Wait()
}
