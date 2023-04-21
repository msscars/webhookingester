package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	var config Config
	config.GetConf()

	routings := []Routing{}

	for i, h := range config.Routings {

		matcher := MatcherFactory(h.Matcher)
		publisher := PublisherFactory(h.Publisher)

		if matcher != nil {
			log.Println("Added", h.Matcher.Type, "from config at position", i)
			routings = append(routings, Routing{Matcher: matcher, Publisher: publisher})
		} else {
			log.Println("Could not create matcher from config at position", i)
		}
	}

	expectedApiKey, expectedApiKeyFound := os.LookupEnv("WEBHOOKINGESTER_APIKEY")

	if !expectedApiKeyFound {
		log.Println("Deliver endpoint running without authentication!")
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

		Route(routings, c.Request)

		c.Status(202)

	})

	r.Run()
}
