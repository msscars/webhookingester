package publisher

import (
	"moscars-webhookingester-publisher/shared"

	"github.com/PaesslerAG/jsonpath"
)

func getBodyPart(incommingWebhook *shared.IncommingWebhook, bodySelector string) interface{} {
	body := incommingWebhook.Body

	if bodySelector != "" {
		body, _ = jsonpath.Get(bodySelector, body)
	}

	return body
}
