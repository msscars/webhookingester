package shared

type IncommingWebhook struct {
	Headers map[string]string `json:"headers"`
	Body    interface{}       `json:"body"`
}
