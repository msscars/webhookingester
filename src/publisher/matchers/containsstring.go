package matchers

import (
	"moscars-webhookingester-publisher/shared"
	"strings"
)

type ContainsStringMatcher struct {
	Token string
}

func (h ContainsStringMatcher) Match(request *shared.IncommingWebhook, requestString string) bool {
	if strings.Contains(requestString, h.Token) {
		return true
	}

	return false
}
