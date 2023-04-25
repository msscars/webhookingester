package matchers

import (
	"moscars-webhookingester-publisher/shared"
)

type ContainsHeaderKeyWithValueMatcher struct {
	Key   string
	Value string
}

func (h ContainsHeaderKeyWithValueMatcher) Match(request *shared.IncommingWebhook) bool {
	value, ok := request.Headers[h.Key]
	if ok && value == h.Value {
		return true
	}

	return false
}
