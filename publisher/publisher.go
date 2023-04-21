package publisher

import "net/http"

type Publisher interface {
	Publish(request *http.Request) bool
}
