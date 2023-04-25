package publisher

import (
	"log"
	"moscars-webhookingester-publisher/shared"
)

type NatsPublisher struct {
	Subject      string
	BodySelector string
}

func (p NatsPublisher) Publish(request *shared.IncommingWebhook) bool {
	log.Println("NatsPublisher received:", request)

	body := getBodyPart(request, p.BodySelector)

	// natsUrl, natsUrlFound := os.LookupEnv("WEBHOOKINGESTER_NATSURL")

	// if !natsUrlFound {
	// 	log.Println("No NATS url configured using default one")
	// 	natsUrl = nats.DefaultURL
	// }

	// nc, err := nats.Connect(natsUrl)
	// if err != nil {
	// 	log.Println(err)
	// 	return false
	// }

	// natsConnection, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	// if err != nil {
	// 	log.Println(err)
	// 	return false
	// }

	// defer natsConnection.Close()

	natsConnection := shared.GetNatsConnection()

	if err := natsConnection.Publish(p.Subject, &body); err != nil {
		log.Println(err)
		return false
	}

	return true
}
