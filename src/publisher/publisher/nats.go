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

	natsConnection := shared.GetNatsConnection()

	if err := natsConnection.Publish(p.Subject, &body); err != nil {
		log.Println(err)
		return false
	}

	return true
}
