package apns

import (
	"testing"
)

// Create a new Payload that specifies simple text,
// a badge counter, and a custom notification sound.
func mockPayload() (payload *Payload) {
	payload = new(Payload)
	payload.Alert = "You have mail!"
	payload.Badge = 42
	payload.Sound = "bingbong.aiff"
	return
}

// Create a new DictionaryAlert. Apple recommends you not use
// the more complex alert style unless absolutely necessary.
func mockDictionaryAlert() (dict *DictionaryAlert) {
	args := make([]string, 1)
	args[0] = "localized args"

	dict = new(DictionaryAlert)
	dict.Body = "Complex Message"
	dict.ActionLocKey = "Play a Game!"
	dict.LocKey = "localized key"
	dict.LocArgs = args
	dict.LaunchImage = "image.jpg"
	return
}

func TestBasicAlert(t *testing.T) {
	payload := mockPayload()
	envelope := new(Envelope)

	envelope.AddPayload(payload)

	bytes, _ := envelope.ToBytes()
	json, _ := envelope.PayloadJSON()
	if len(bytes) != 82 {
		t.Error("expected 82 bytes; got", len(bytes))
	}
	if len(json) != 69 {
		t.Error("expected 69 bytes; got", len(json))
	}
}

func TestDictionaryAlert(t *testing.T) {
	dict := mockDictionaryAlert()
	payload := mockPayload()
	payload.Alert = dict

	envelope := new(Envelope)
	envelope.AddPayload(payload)

	bytes, _ := envelope.ToBytes()
	json, _ := envelope.PayloadJSON()
	if len(bytes) != 207 {
		t.Error("expected 207 bytes; got", len(bytes))
	}
	if len(json) != 194 {
		t.Error("expected 194 bytes; got", len(bytes))
	}
}

func TestCustomParameters(t *testing.T) {
	payload := mockPayload()
	envelope := new(Envelope)

	envelope.AddPayload(payload)
	envelope.Set("foo", "bar")

	if envelope.Get("foo") != "bar" {
		t.Error("unable to set a custom property")
	}
	if envelope.Get("not_set") != nil {
		t.Error("expected a missing key to return nil")
	}

	bytes, _ := envelope.ToBytes()
	json, _ := envelope.PayloadJSON()
	if len(bytes) != 94 {
		t.Error("expected 94 bytes; got", len(bytes))
	}
	if len(json) != 81 {
		t.Error("expected 81 bytes; got", len(json))
	}
}
