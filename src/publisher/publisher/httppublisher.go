package publisher

import (
	"bytes"
	"encoding/json"
	"log"
	"moscars-webhookingester-publisher/shared"
	"net/http"
)

type HttpPublisher struct {
	Uri          string
	BodySelector string
}

func (p HttpPublisher) Publish(request *shared.IncommingWebhook) bool {
	log.Println("HttpPublisher received:", request)

	body := getBodyPart(request, p.BodySelector)

	bodyByte, err := json.Marshal(body)
	if err != nil {
		log.Println(err)
		return false
	}

	req, err := http.NewRequest("POST", p.Uri, bytes.NewBuffer(bodyByte))
	if err != nil {
		log.Println(err)
		return false
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return false
	}
	defer response.Body.Close()

	log.Println("response Status:", response.Status)
	log.Println("response Headers:", response.Header)

	return true
}
