package matchers

import (
	"encoding/json"
	"moscars-webhookingester-publisher/shared"
	"testing"
)

func Test_ExistingPropertyWithCorrectValue_Match(t *testing.T) {
	sut := PropertyValueMatcher{Path: "$.prop1", Value: "xyz"}

	var iw interface{}
	jsonstr := `{"prop1": "xyz"}`
	err := json.Unmarshal([]byte(jsonstr), &iw)
	if err != nil {
		panic(err)
	}

	if !sut.Match(&shared.IncommingWebhook{Body: iw}) {
		t.Fail()
	}
}

func Test_ExistingPropertyWithIncorrectValue_NoMatch(t *testing.T) {
	sut := PropertyValueMatcher{Path: "$.prop1", Value: "xyz"}

	var iw interface{}
	jsonstr := `{"prop1": "xxx"}`
	err := json.Unmarshal([]byte(jsonstr), &iw)
	if err != nil {
		panic(err)
	}

	if sut.Match(&shared.IncommingWebhook{Body: iw}) {
		t.Fail()
	}
}

func Test_NotExistingProperty_NoMatch(t *testing.T) {
	sut := PropertyValueMatcher{Path: "$.prop1", Value: "xyz"}

	var iw interface{}
	jsonstr := `{"prop2": "xxx"}`
	err := json.Unmarshal([]byte(jsonstr), &iw)
	if err != nil {
		panic(err)
	}

	if sut.Match(&shared.IncommingWebhook{Body: iw}) {
		t.Fail()
	}
}
