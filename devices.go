package smartmeter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

/**
* http://api.iglucose.com/static/documentation/#-devices-service
 */

type GetDevicesResponse struct {
	Devices []string       `json:"devices,omitempty"`
	Status  ResponseStatus `json:"status,omitempty"`
}

type ValidateDeviceRequest struct {
	ApiKey   string `json:"api_key,omitempty"`
	DeviceId string `json:"device_id,omitempty"`
}

type ValidateDeviceResponse struct {
	DeviceId    string `json:"device_id,omitempty"`
	DeviceModel string `json:"device_model,omitempty"`
	Imei        string `json:"imei,omitempty"`
	IsValid     bool   `json:"is_valid,omitempty"`
	Shortcode   string `json:"short_code,omitempty"`

	Status ResponseStatus `json:"status,omitempty"`
}

// This service operates using a GET operation to
// pull data by API key. Only those devices associated
// with the API key will be returned by the service.
func (c *Client) GetDevices() (GetDevicesResponse, error) {
	var devicesResponse GetDevicesResponse

	url, err := url.Parse(c.Url)
	if err != nil {
		return devicesResponse, err
	}

	values := url.Query()
	values.Add("api_key", c.Key)

	url.RawQuery = values.Encode()
	url.Path = "/devices/"

	resp, err := http.Get(url.String())
	if err != nil {
		return devicesResponse, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return devicesResponse, err
	}

	err = json.Unmarshal(body, &devicesResponse)

	return devicesResponse, err
}

// This service operates using a POST operation to pull
// data by API key and device ID. Only devices associated
// with the API key will return an “is_valid” value of “true”.
func (c *Client) ValidateDevice(deviceId string) (ValidateDeviceResponse, error) {
	var validateResponse ValidateDeviceResponse

	validateRequest := new(ValidateDeviceRequest)
	validateRequest.ApiKey = c.Key
	validateRequest.DeviceId = deviceId

	dataBytes, err := json.Marshal(validateRequest)
	if err != nil {
		return validateResponse, err
	}

	url, err := url.Parse(c.Url)
	if err != nil {
		return validateResponse, err
	}

	url.Path = "/devices/validate/"

	resp, err := http.Post(url.String(), "application/json", bytes.NewBuffer(dataBytes))
	if err != nil {
		return validateResponse, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return validateResponse, err
	}

	fmt.Println(string(body))

	err = json.Unmarshal(body, &validateResponse)

	return validateResponse, err
}
