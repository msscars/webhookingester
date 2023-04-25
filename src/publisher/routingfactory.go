package main

import (
	"log"
	"moscars-webhookingester-publisher/matchers"
	"moscars-webhookingester-publisher/publisher"
)

func CreateRoutings(config Config) []Routing {
	routings := []Routing{}

	for i, h := range config.Routings {

		matcher := matcherFactory(h.Matcher)
		publisher := publisherFactory(h.Publisher)

		if matcher != nil && publisher != nil {
			log.Println("Added routing with", h.Matcher.Type, "and", h.Publisher.Type, "from config at position", i)
			routings = append(routings, Routing{Matcher: matcher, Publisher: publisher})
		} else {
			log.Println("Could not create matcher or publisher from config at position", i)
		}
	}

	return routings
}

func matcherFactory(m MatcherConfig) matchers.Matcher {
	switch m.Type {
	case "ContainsPropertyWithValue":
		return matchers.ContainsPropertyWithValueMatcher{Path: m.Path, Value: m.Value}
	case "ContainsHeaderKey":
		return matchers.ContainsHeaderKeyMatcher{Key: m.Key}
	case "ContainsHeaderKeyWithValue":
		return matchers.ContainsHeaderKeyWithValueMatcher{Key: m.Key, Value: m.Value}
	}

	return nil
}

func publisherFactory(m PublisherConfig) publisher.Publisher {
	switch m.Type {
	case "http":
		return publisher.HttpPublisher{Uri: m.Uri, BodySelector: m.BodySelector}
	case "nats":
		return publisher.NatsPublisher{Subject: m.Subject, BodySelector: m.BodySelector}
	}

	return nil
}
