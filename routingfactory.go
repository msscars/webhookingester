package main

import (
	"moscars-webhookingester/matchers"
	"moscars-webhookingester/publisher"
)

func MatcherFactory(m MatcherConfig) matchers.Matcher {
	switch m.Type {
	case "ContainsString":
		return matchers.ContainsStringMatcher{Token: m.Token}
	case "ContainsHeaderKey":
		return matchers.ContainsHeaderKeyMatcher{Key: m.Key}
	}

	return nil
}

func PublisherFactory(m PublisherConfig) publisher.Publisher {
	switch m.Type {
	case "http":
		return publisher.HttpPublisher{Uri: m.Uri, BodySelector: m.BodySelector}
	}

	return nil
}
