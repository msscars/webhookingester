package matchers

import (
	"net/http"
)

type ContainsHeaderKeyMatcher struct {
	Key string
}

func (h ContainsHeaderKeyMatcher) Match(request *http.Request, requestString string) bool {
	_, ok := request.Header[h.Key]
	if ok {
		return true
	}

	return false
}
