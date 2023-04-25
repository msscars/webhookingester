package publisher

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"moscars-webhookingester-publisher/shared"
	"net/http"

	"github.com/PaesslerAG/jsonpath"
)

type HttpPublisher struct {
	Uri          string
	BodySelector string
}

func (p HttpPublisher) Publish(request *shared.IncommingWebhook) bool {
	body := request.Body

	if p.BodySelector != "" {
		body, _ = jsonpath.Get(p.BodySelector, body)
	}

	buf := &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, body)

	req, error := http.NewRequest("POST", p.Uri, buf)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(req)
	if error != nil {
		return false
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)

	return true
}
