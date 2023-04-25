package shared

import (
	"log"
	"os"
	"sync"

	"github.com/nats-io/nats.go"
)

type IncommingWebhook struct {
	Headers map[string]string `json:"headers"`
	Body    interface{}       `json:"body"`
}

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *nats.EncodedConn

func GetNatsConnection() *nats.EncodedConn {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			log.Println("Creating nats connection")

			natsUrl, natsUrlFound := os.LookupEnv("WEBHOOKINGESTER_NATSURL")

			if !natsUrlFound {
				log.Println("No NATS url configured using default one")
				natsUrl = nats.DefaultURL
			}

			nc, err := nats.Connect(natsUrl)
			if err != nil {
				log.Fatal(err)
			}

			ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
			if err != nil {
				log.Fatal(err)
			}

			singleInstance = ec
		} else {
			log.Println("Using existing nats connection")
		}
	} else {
		log.Println("Using existing nats connection")
	}

	return singleInstance
}
