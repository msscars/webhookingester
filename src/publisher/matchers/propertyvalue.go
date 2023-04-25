package matchers

import (
	"moscars-webhookingester-publisher/shared"

	"github.com/PaesslerAG/jsonpath"
)

type PropertyValueMatcher struct {
	Path  string
	Value string
}

func (h PropertyValueMatcher) Match(request *shared.IncommingWebhook) bool {
	token, err := jsonpath.Get(h.Path, request.Body)
	if err != nil {
		return false
	}

	value := token.(string)

	if value == h.Value {
		return true
	}

	return false
}
