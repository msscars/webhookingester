package main

import (
	"encoding/json"
	"log"
	"moscars-webhookingester-publisher/shared"
	"os"
	"time"

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

	defer nc.Drain()

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}
	defer ec.Close()

	js, _ := ec.Conn.JetStream()

	streamName := "PUBLISH_QUEUE"

	js.AddStream(&nats.StreamConfig{
		Name:     streamName,
		Subjects: []string{"publish-queue"},
	})

	if _, err := js.QueueSubscribe("publish-queue", "publishers", func(msg *nats.Msg) {

		var incommingWebHook *shared.IncommingWebhook
		err := json.Unmarshal(msg.Data, &incommingWebHook)

		if err != nil {
			log.Println(err)
		} else {
			log.Println("Message received:", incommingWebHook)
			Route(routings, incommingWebHook)
			log.Println("Message routed")
		}

		msg.Ack()
	}, nats.AckExplicit()); err != nil {
		log.Println(err)
	}

	for {
		log.Println("still alive")
		time.Sleep(10 * time.Second)
	}
}
