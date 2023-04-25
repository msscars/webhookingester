package matchers

import (
	"moscars-webhookingester-publisher/shared"
)

type ContainsHeaderKeyMatcher struct {
	Key string
}

func (h ContainsHeaderKeyMatcher) Match(request *shared.IncommingWebhook) bool {
	_, ok := request.Headers[h.Key]
	if ok {
		return true
	}

	return false
}
