package main

import (
	"fmt"
	"log"
	"moscars-webhookingester-api/matchers"
	"moscars-webhookingester-api/publisher"
	"net/http"
	"net/http/httputil"
)

type Routing struct {
	Matcher   matchers.Matcher    `yaml:"matcher"`
	Publisher publisher.Publisher `yaml:"publisher"`
}

func Route(routings []Routing, request *http.Request) bool {
	var match bool = false

	requestDump, err := httputil.DumpRequest(request, true)
	if err != nil {
		fmt.Println(err)
	} else {
		reqString := string(requestDump)

		for i, h := range routings {
			if h.Matcher.Match(request, reqString) {
				match = true
				log.Println("matcher found at position", i)

				if h.Publisher.Publish(request) {
					log.Println("publisher at position", i, "succeeded")
				}
			}
		}
	}

	return match
}
