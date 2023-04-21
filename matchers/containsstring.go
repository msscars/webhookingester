package matchers

import (
	"net/http"
	"strings"
)

type ContainsStringMatcher struct {
	Token string
}

func (h ContainsStringMatcher) Match(request *http.Request, requestString string) bool {
	if strings.Contains(requestString, h.Token) {
		return true
	}

	return false
}
