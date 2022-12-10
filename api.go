package smartmeter

import (
	"bytes"
	"encoding/json"
	"time"
)

type ApiBoolean bool

type ApiTime time.Time

type Client struct {
	Key string
	Url string
}

type ResponseStatus struct {
	StatusCode    int    `json:"status_code,omitempty"`
	StatusMessage string `json:"status_message,omitempty"`
}

func NewClient(url string, key string) *Client {
	client := new(Client)
	client.Key = key
	client.Url = url
	return client
}

func (ab *ApiBoolean) MarshalJSON() ([]byte, error) {
	bValue := bool(*ab)
	return json.Marshal(bValue)
}

func (ab *ApiBoolean) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case "true":
		*ab = true
	case "1":
		*ab = true
	default:
		*ab = false
	}
	return nil
}

func (at *ApiTime) MarshalJSON() ([]byte, error) {
	aValue := time.Time(*at)
	return json.Marshal(aValue)
}

func (at *ApiTime) UnmarshalJSON(data []byte) error {
	var t time.Time

	t, err := time.Parse("2006-01-02T15:04:05", string(bytes.Trim(data, "\"")))
	if err != nil {
		return err
	}

	*at = ApiTime(t)

	return nil
}
