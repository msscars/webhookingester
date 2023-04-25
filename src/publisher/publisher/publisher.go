package publisher

import "moscars-webhookingester-publisher/shared"

type Publisher interface {
	Publish(request *shared.IncommingWebhook) bool
}
