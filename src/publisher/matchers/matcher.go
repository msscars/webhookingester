package matchers

import "moscars-webhookingester-publisher/shared"

type Matcher interface {
	Match(request *shared.IncommingWebhook, requestString string) bool
}
