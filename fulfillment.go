package smartmeter

/**
* http://api.iglucose.com/static/documentation/#-fulfillment-service
 */

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type GetOrderDetailByOrderNumberResponse struct {
	Order  Order               `json:"order,omitempty"`
	Status OrderResponseStatus `json:"status,omitempty"`
}

type CreateOrderRequest struct {
	Order
}

type CreateOrderResponse struct {
	Orders []Order             `json:"orders,omitempty"`
	Status OrderResponseStatus `json:"status,omitempty"`
}

type DeleteOrderResponse struct {
	Orders []Order             `json:"orders,omitempty"`
	Status OrderResponseStatus `json:"status,omitempty"`
}

type LineItem struct {
	BoxNumber                int        `json:"box_number,omitempty"`
	ControlSolutionLotNumber string     `json:"control_solution_lot_number,omitempty"`
	DeviceModel              string     `json:"device_model,omitempty"`
	Id                       int        `json:"id,omitempty"`
	Imei                     string     `json:"imei,omitempty"`
	IsShippable              ApiBoolean `json:"is_shippable,omitempty"`
	LineItem                 int        `json:"line_item,omitempty"`
	LineKey                  string     `json:"line_key,omitempty"`
	LineName                 string     `json:"line_name,omitempty"`
	LotNumber                string     `json:"lot_number,omitempty"`
	OrderNumber              string     `json:"order_number,omitempty"`
	Qty                      int        `json:"qty,omitempty"`
	Quantity                 int        `json:"quantity,omitempty"`
	ParentLineKey            string     `json:"parent_line_key,omitempty"`
	PartNumber               string     `json:"part_number,omitempty"`
	SerialNumber             string     `json:"serial_number,omitempty"`
	Sku                      string     `json:"sku,omitempty"`
	TrackingNumber           string     `json:"tracking_number,omitempty"`
	UdcNumber                string     `json:"udc_number,omitempty"`
}

// type OrderBoolean bool

// type OrderDateTime time.Time

type Order struct {
	Address1          string     `json:"address1"`
	Address2          string     `json:"address2"`
	Carrier           string     `json:"carrier,omitempty"`
	City              string     `json:"city,omitempty"`
	Country           string     `json:"country,omitempty"`
	CustomerId        string     `json:"customer_id,omitempty"`
	CustomerName      string     `json:"customer_name,omitempty"`
	DateCreated       *ApiTime   `json:"date_created,omitempty"`
	DateShipped       *ApiTime   `json:"date_shipped,omitempty"`
	Id                int        `json:"id,omitempty"`
	IsRefill          ApiBoolean `json:"is_refill,omitempty"`
	IsReplacement     ApiBoolean `json:"is_replacement,omitempty"`
	IsSample          ApiBoolean `json:"is_sample,omitempty"`
	IsTransfer        ApiBoolean `json:"is_transfer,omitempty"`
	LastUpdated       *ApiTime   `json:"last_updated,omitempty"`
	Lines             []LineItem `json:"lines,omitempty"`
	Notes             string     `json:"notes,omitempty"`
	OrderNumber       string     `json:"order_number,omitempty"`
	PhoneNumber       string     `json:"phone_number,omitempty"`
	PoNumber          string     `json:"po_number,omitempty"`
	RefillingDeviceId string     `json:"refilling_device_id,omitempty"`
	ReplacingDeviceId string     `json:"replacing_device_id,omitempty"`
	RmaId             string     `json:"rma_id,omitempty"`
	ShippingMethod    string     `json:"shipping_method,omitempty"`
	Source            string     `json:"source,omitempty"`
	State             string     `json:"state,omitempty"`
	Status            string     `json:"status,omitempty"`
	TrackingNumber    string     `json:"tracking_number,omitempty"`
	ZipCode           string     `json:"zipcode,omitempty"`
}

type OrderResponseStatus struct {
	ResponseStatus
	OrdersInResponse int `json:"orders_in_response,omitempty"`
}

func (c *Client) CreateOrder(ord CreateOrderRequest) (CreateOrderResponse, error) {
	var orderResponse CreateOrderResponse

	ord.DateCreated = nil
	ord.LastUpdated = nil

	dataBytes, err := json.Marshal(ord)
	if err != nil {
		return orderResponse, err
	}

	url, err := url.Parse(c.Url)
	if err != nil {
		return orderResponse, err
	}

	values := url.Query()
	values.Add("api_key", c.Key)

	url.RawQuery = values.Encode()
	url.Path = "/fulfillment/"

	resp, err := http.Post(url.String(), "application/json", bytes.NewBuffer(dataBytes))
	if err != nil {
		return orderResponse, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return orderResponse, err
	}

	err = json.Unmarshal(body, &orderResponse)

	return orderResponse, err
}

func (c *Client) DeleteOrder(orderNumber string) (DeleteOrderResponse, error) {
	var deleteResponse DeleteOrderResponse

	url, err := url.Parse(c.Url)
	if err != nil {
		return deleteResponse, err
	}

	values := url.Query()
	values.Add("api_key", c.Key)
	values.Add("order_number", orderNumber)

	url.RawQuery = values.Encode()
	url.Path = "/fulfillment/number"

	fmt.Println(url.String())

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url.String(), nil)
	if err != nil {
		return deleteResponse, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return deleteResponse, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return deleteResponse, err
	}

	err = json.Unmarshal(body, &deleteResponse)

	return deleteResponse, err
}

func (c *Client) GetOrderDetailByOrderNumber(orderNumber string) (GetOrderDetailByOrderNumberResponse, error) {
	var orderResponse GetOrderDetailByOrderNumberResponse

	url, err := url.Parse(c.Url)
	if err != nil {
		return orderResponse, err
	}

	values := url.Query()
	values.Add("api_key", c.Key)
	values.Add("order_number", orderNumber)

	url.RawQuery = values.Encode()
	url.Path = "/fulfillment/number"

	resp, err := http.Get(url.String())
	if err != nil {
		return orderResponse, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return orderResponse, err
	}

	err = json.Unmarshal(body, &orderResponse)

	return orderResponse, err
}

// func (ob *OrderBoolean) UnmarshalJSON(data []byte) error {
// 	switch string(data) {
// 	case "true":
// 		*ob = true
// 	case "1":
// 		*ob = true
// 	default:
// 		*ob = false
// 	}
// 	return nil
// }

// func (odt *OrderDateTime) UnmarshalJSON(data []byte) error {
// 	var t time.Time

// 	t, err := time.Parse("2006-01-02T15:04:05", string(bytes.Trim(data, "\"")))
// 	if err != nil {
// 		return err
// 	}

// 	*odt = OrderDateTime(t)

// 	return nil
// }
