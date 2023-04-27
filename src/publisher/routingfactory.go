package main

import (
	"log"
	"moscars-webhookingester-publisher/matchers"
	"moscars-webhookingester-publisher/publisher"
	"strings"
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
	switch strings.ToLower(m.Type) {
	case "property":
		return matchers.PropertyMatcher{Path: m.Path}
	case "propertyvalue":
		return matchers.PropertyValueMatcher{Path: m.Path, Value: m.Value}
	case "headerkey":
		return matchers.HeaderKeyMatcher{Key: m.Key}
	case "headervalue":
		return matchers.HeaderValueMatcher{Key: m.Key, Value: m.Value}
	}

	return nil
}

func publisherFactory(m PublisherConfig) publisher.Publisher {
	switch m.Type {
	case "http":
		return publisher.HttpPublisher{Uri: m.Uri, BodySelector: m.BodySelector}
	case "nats":
		return publisher.NatsPublisher{Subject: m.Subject, BodySelector: m.BodySelector}
	case "kafka":
		return publisher.KafkaPublisher{Subject: m.Subject, BodySelector: m.BodySelector}
	}

	return nil
}
