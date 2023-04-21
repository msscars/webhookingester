package publisher

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/PaesslerAG/jsonpath"
)

type HttpPublisher struct {
	Uri          string
	BodySelector string
}

func (p HttpPublisher) Publish(request *http.Request) bool {
	v := interface{}(nil)
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return false
	}

	request.Body = io.NopCloser(bytes.NewBuffer(body))

	if p.BodySelector != "" {
		json.Unmarshal(body, &v)

		selectedBody, _ := jsonpath.Get(p.BodySelector, v)

		body, _ = json.Marshal(selectedBody)
	}

	request, error := http.NewRequest("POST", p.Uri, bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		return false
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	// body, _ := ioutil.ReadAll(response.Body)
	// fmt.Println("response Body:", string(body))

	return true
}
