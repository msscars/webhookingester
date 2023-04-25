package matchers

import (
	"moscars-webhookingester-publisher/shared"

	"github.com/PaesslerAG/jsonpath"
)

type PropertyMatcher struct {
	Path string
}

func (h PropertyMatcher) Match(request *shared.IncommingWebhook) bool {
	_, err := jsonpath.Get(h.Path, request.Body)
	if err != nil {
		return false
	}

	return true
}
