package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"
)

type IncommingWebhook struct {
	Headers map[string]string `json:"headers"`
	Body    interface{}       `json:"body"`
}

func captureRequestData(req *http.Request) (IncommingWebhook, error) {
	v := interface{}(nil)
	webhookRequest := IncommingWebhook{Headers: map[string]string{}}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
	}

	for name, values := range req.Header {
		for _, value := range values {
			webhookRequest.Headers[name] = value
		}
	}

	json.Unmarshal(body, &v)

	return webhookRequest, nil
}

func main() {
	godotenv.Load(".env")

	expectedApiKey, expectedApiKeyFound := os.LookupEnv("WEBHOOKINGESTER_APIKEY")
	natsUrl, natsUrlFound := os.LookupEnv("WEBHOOKINGESTER_NATSURL")

	if !expectedApiKeyFound {
		log.Println("Deliver endpoint running without authentication!")
	}

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
	defer ec.Close()

	r := gin.Default()

	r.GET("/healthz", func(c *gin.Context) {
		c.Status(200)
	})

	r.POST("/deliver", func(c *gin.Context) {

		if expectedApiKeyFound {
			actualApiKey := c.Query("key")
			if expectedApiKey != actualApiKey {
				c.Status(401)
				return
			}
		}

		b, err := captureRequestData(c.Request)

		if err != nil {
			c.Status(500)
			return
		}

		if err := ec.Publish("publish-queue", &b); err != nil {
			log.Println(err)
		}

		c.Status(202)
	})

	r.Run()
}
