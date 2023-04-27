package publisher

import (
	"encoding/json"
	"log"
	"moscars-webhookingester-publisher/shared"

	"gopkg.in/Shopify/sarama.v1"
)

type KafkaPublisher struct {
	Subject      string
	BodySelector string
}

func (p KafkaPublisher) Publish(request *shared.IncommingWebhook) bool {
	log.Println("KafkaPublisher received:", request)

	body := getBodyPart(request, p.BodySelector)
	b, err := json.Marshal(body)

	connection, err := shared.GetKafkaConnection()
	if err != nil {
		log.Println(err)
		return false
	}

	msg := &sarama.ProducerMessage{
		Topic: p.Subject,
		Value: sarama.ByteEncoder(b),
	}

	partition, offset, err := connection.SendMessage(msg)
	if err != nil {
		log.Println(err)
		return false
	}

	log.Println("kafka message sent to partion", partition, "with offset", offset)

	return true
}
