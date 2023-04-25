package main

import (
	"log"
	"moscars-webhookingester-publisher/matchers"
	"moscars-webhookingester-publisher/publisher"
	"moscars-webhookingester-publisher/shared"
)

type Routing struct {
	Matcher   matchers.Matcher    `yaml:"matcher"`
	Publisher publisher.Publisher `yaml:"publisher"`
}

func Route(routings []Routing, request *shared.IncommingWebhook) bool {
	var match bool = false

	for i, h := range routings {
		if h.Matcher.Match(request) {
			match = true
			log.Println("matcher found at position", i)

			if h.Publisher.Publish(request) {
				log.Println("publisher at position", i, "succeeded")
			}
		}
	}

	return match
}
