package main

import (
	"bufio"
	"bytes"
	"log"
	"net/http"
	"sync"

	"github.com/nats-io/nats.go"
)

func main() {
	var config Config
	config.GetConf()

	routings := CreateRoutings(config)

	nc, err := nats.Connect(nats.DefaultURL)
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
