package matchers

import (
	"moscars-webhookingester-publisher/shared"
)

type HeaderKeyMatcher struct {
	Key string
}

func (h HeaderKeyMatcher) Match(request *shared.IncommingWebhook) bool {
	_, ok := request.Headers[h.Key]
	if ok {
		return true
	}

	return false
}
