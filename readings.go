package smartmeter

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

/**
* http://api.iglucose.com/static/documentation/#-readings-service
 */

type GetReadingsByFilterRequest struct {
	ApiKey          string   `json:"api_key,omitempty"`
	DateStart       *ApiTime `json:"date_start,omitempty"`
	DateEnd         *ApiTime `json:"date_end,omitempty"`
	DeviceIds       []string `json:"device_ids,omitempty"`
	IngestDateStart *ApiTime `json:"ingest_date_start,omitempty"`
	IngestDateEnd   *ApiTime `json:"ingest_date_end,omitempty"`
	ReadingType     []string `json:"reading_type,omitempty"`
}

type GetReadingsByFilterResponse struct {
	Readings []Reading             `json:"readings,omitempty"`
	Status   ReadingResponseStatus `json:"status,omitempty"`
}

type GetReadingsByMeterIdsRequest struct {
	ApiKey    string   `json:"api_key,omitempty"`
	DeviceIds []string `json:"device_ids,omitempty"`
}

type GetReadingsByMeterIdsResponse struct {
	Readings []Reading             `json:"readings,omitempty"`
	Status   ReadingResponseStatus `json:"status,omitempty"`
}

type GetReadingsByReadingIdsRequest struct {
	ApiKey     string `json:"api_key,omitempty"`
	ReadingIds []int  `json:"reading_ids,omitempty"`
}

type GetReadingsByReadingIdsResponse struct {
	Readings []Reading             `json:"readings,omitempty"`
	Status   ReadingResponseStatus `json:"status,omitempty"`
}

type ReadingResponseStatus struct {
	ResponseStatus

	DevicesInResponse  int `json:"devices_in_response,omitempty"`
	ReadingsInResponse int `json:"readings_in_response,omitempty"`
}

type Reading struct {
	Battery          int     `json:"battery,omitempty"`
	BeforeMeal       bool    `json:"before_meal,omitempty"`
	BloodGlucoseMgdl float32 `json:"blood_glucose_mgdl,omitempty"`
	BloodGlucoseMmol float32 `json:"blood_glucose_mmol,omitempty"`
	DateRecorded     ApiTime `json:"date_recorded,omitempty"`
	DateReceived     ApiTime `json:"date_received,omitempty"`
	DeviceId         string  `json:"device_id,omitempty"`
	DeviceModel      string  `json:"device_model,omitempty"`
	DiastolicMmhg    int     `json:"diastolic_mmhg,omitempty"`
	EventFlag        string  `json:"event_flag,omitempty"`
	Irregular        bool    `json:"irregular,omitempty"`
	PulseBpm         int     `json:"pulse_bpm,omitempty"`
	ReadingId        int     `json:"reading_id,omitempty"`
	ReadingType      string  `json:"reading_type,omitempty"`
	ShortCode        string  `json:"short_code,omitempty"`
	Spo2             int     `json:"spo2,omitempty"`
	SystolicMmhg     int     `json:"systolic_mmhg,omitempty"`
	TareKg           int     `json:"tare_kg,omitempty"`
	TareLbs          int     `json:"tare_lbs,omitempty"`
	TimeZoneOffset   float32 `json:"time_zone_offset,omitempty"`
	WeightKg         int     `json:"weight_kg,omitempty"`
	WeightLbs        int     `json:"weight_lbs,omitempty"`
}

func (c *Client) GetReadingsByFilter(filterRequest GetReadingsByFilterRequest) (GetReadingsByFilterResponse, error) {
	var readingsResponse GetReadingsByFilterResponse

	dataBytes, err := json.Marshal(filterRequest)
	if err != nil {
		return readingsResponse, err
	}

	url, err := url.Parse(c.Url)
	if err != nil {
		return readingsResponse, err
	}

	url.Path = "/readings/"

	resp, err := http.Post(url.String(), "application/json", bytes.NewBuffer(dataBytes))
	if err != nil {
		return readingsResponse, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return readingsResponse, err
	}

	err = json.Unmarshal(body, &readingsResponse)

	return readingsResponse, err
}

func (c *Client) GetReadingsByMeterIds(readingsRequest GetReadingsByMeterIdsRequest) (GetReadingsByMeterIdsResponse, error) {
	var readingsResponse GetReadingsByMeterIdsResponse

	dataBytes, err := json.Marshal(readingsRequest)
	if err != nil {
		return readingsResponse, err
	}

	url, err := url.Parse(c.Url)
	if err != nil {
		return readingsResponse, err
	}

	url.Path = "/readings/"

	resp, err := http.Post(url.String(), "application/json", bytes.NewBuffer(dataBytes))
	if err != nil {
		return readingsResponse, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return readingsResponse, err
	}

	err = json.Unmarshal(body, &readingsResponse)

	return readingsResponse, err
}

func (c *Client) GetReadingsByReadingIds(readingsRequest GetReadingsByReadingIdsRequest) (GetReadingsByReadingIdsResponse, error) {
	var readingsResponse GetReadingsByReadingIdsResponse

	dataBytes, err := json.Marshal(readingsRequest)
	if err != nil {
		return readingsResponse, err
	}

	url, err := url.Parse(c.Url)
	if err != nil {
		return readingsResponse, err
	}

	url.Path = "/readings/"

	resp, err := http.Post(url.String(), "application/json", bytes.NewBuffer(dataBytes))
	if err != nil {
		return readingsResponse, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return readingsResponse, err
	}

	err = json.Unmarshal(body, &readingsResponse)

	return readingsResponse, err
}
