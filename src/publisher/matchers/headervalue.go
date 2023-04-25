package matchers

import (
	"moscars-webhookingester-publisher/shared"
)

type HeaderValueMatcher struct {
	Key   string
	Value string
}

func (h HeaderValueMatcher) Match(request *shared.IncommingWebhook) bool {
	value, ok := request.Headers[h.Key]
	if ok && value == h.Value {
		return true
	}

	return false
}
