package shared

import (
	"errors"
	"log"
	"os"
	"sync"

	"github.com/nats-io/nats.go"
	"gopkg.in/Shopify/sarama.v1"
)

type IncommingWebhook struct {
	Headers map[string]string `json:"headers"`
	Body    interface{}       `json:"body"`
}

var natsConnectionLock = &sync.Mutex{}

var natsConnectionInstance *nats.EncodedConn

var kafkaConnectionLock = &sync.Mutex{}

var kafkaConnectionInstance sarama.SyncProducer

func GetNatsConnection() *nats.EncodedConn {
	if natsConnectionInstance == nil {
		natsConnectionLock.Lock()
		defer natsConnectionLock.Unlock()
		if natsConnectionInstance == nil {
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

			natsConnectionInstance = ec
		} else {
			log.Println("Using existing nats connection")
		}
	} else {
		log.Println("Using existing nats connection")
	}

	return natsConnectionInstance
}

func GetKafkaConnection() (sarama.SyncProducer, error) {
	if kafkaConnectionInstance == nil {
		kafkaConnectionLock.Lock()
		defer kafkaConnectionLock.Unlock()
		if kafkaConnectionInstance == nil {
			log.Println("Creating Kafka connection")
			kafkaUrl, kafkaUrlFound := os.LookupEnv("WEBHOOKINGESTER_KAFKAURL")

			if !kafkaUrlFound {
				log.Println("No Kafka url configured using default one")
				return nil, errors.New("No Kafka url configured")
			}

			config := sarama.NewConfig()
			config.Producer.Return.Successes = true
			config.Producer.Retry.Max = 5

			conn, err := sarama.NewSyncProducer([]string{kafkaUrl}, config)
			if err != nil {
				return nil, err
			}

			kafkaConnectionInstance = conn
		} else {
			log.Println("Using existing Kafka connection")
		}
	} else {
		log.Println("Using existing Kafka connection")
	}

	return kafkaConnectionInstance, nil
}
