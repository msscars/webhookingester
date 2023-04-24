package main

import (
	"bytes"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"
)

func captureRequestData(req *http.Request) ([]byte, error) {
	var b = &bytes.Buffer{} // holds serialized representation
	var err error
	if err = req.Write(b); err != nil { // serialize request to HTTP/1.1 wire format
		return nil, err
	}
	return b.Bytes(), nil
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

		nc.Publish("publish-queue", b)

		c.Status(202)
	})

	r.Run()
}
