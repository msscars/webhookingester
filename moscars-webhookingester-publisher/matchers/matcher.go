package matchers

import "net/http"

type Matcher interface {
	Match(request *http.Request, requestString string) bool
}
