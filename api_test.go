package smartmeter

import (
	"fmt"
	"testing"
	"time"
)

const (
	testKey                    = ""
	testUrl                    = "http://dev.api.iglucose.com"
	testOrderNumber            = "abcd12345771"
	testDeviceId               = "999970400000000"
	testReadingId              = 500000003181282
	testReadingFilterStartDate = "2022-05-12T00:00:00"
	testReadingFilterEndDate   = "2022-05-12T18:05:00"
)

/**
* Devices Service
 */
func TestGetDevices(t *testing.T) {
	client := NewClient(testUrl, testKey)

	GetDevicesResponse, err := client.GetDevices()
	if err != nil {
		t.Errorf(err.Error())
	}

	if GetDevicesResponse.Status.StatusCode != 200 {
		t.Errorf(GetDevicesResponse.Status.StatusMessage)
	}
}

func TestValidateDevice(t *testing.T) {
	client := NewClient(testUrl, testKey)

	ValidateDeviceResponse, err := client.ValidateDevice(testDeviceId)
	if err != nil {
		t.Errorf(err.Error())
	}

	if ValidateDeviceResponse.Status.StatusCode != 200 {
		t.Errorf(ValidateDeviceResponse.Status.StatusMessage)
	}
}

/**
* Fulfillment Service
 */

func TestCreateOrder(t *testing.T) {
	client := NewClient(testUrl, testKey)

	orderRequest := new(CreateOrderRequest)
	orderRequest.Address1 = "1234 Fake Street"
	orderRequest.City = "Cityland"
	orderRequest.Country = "United States"
	orderRequest.CustomerId = "12353456"
	orderRequest.CustomerName = "Test Patient"
	orderRequest.OrderNumber = testOrderNumber
	orderRequest.ShippingMethod = "MAIL"
	orderRequest.State = "AZ"
	orderRequest.ZipCode = "54767"

	lineItem := new(LineItem)
	lineItem.LineName = "iGlucose Monitoring System"
	lineItem.Quantity = 1
	lineItem.Sku = "06MGLUKITUSLTE"

	orderRequest.Lines = append(orderRequest.Lines, *lineItem)

	orderResponse, err := client.CreateOrder(*orderRequest)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(orderResponse.Status.StatusMessage)
}

func TestGetOrderDetailByOrderNumber(t *testing.T) {
	client := NewClient(testUrl, testKey)

	GetOrderDetailByOrderNumberResponse, err := client.GetOrderDetailByOrderNumber(testOrderNumber)
	if err != nil {
		t.Errorf(err.Error())
	}

	if GetOrderDetailByOrderNumberResponse.Status.StatusCode != 200 {
		t.Errorf("failed to retrieve order")
	}
}

func TestDeleteOrder(t *testing.T) {
	client := NewClient(testUrl, testKey)

	DeleteOrderResponse, err := client.DeleteOrder(testOrderNumber)
	if err != nil {
		t.Error(err.Error())
	}

	if DeleteOrderResponse.Status.StatusCode != 200 {
		t.Errorf("failed to delete order")
	}
}

// Readings Service
func TestGetReadingsByFilter(t *testing.T) {
	client := NewClient(testUrl, testKey)

	startDate, err := time.Parse("2006-01-02T15:04:05", testReadingFilterStartDate)
	if err != nil {
		t.Errorf(err.Error())
	}

	endDate, err := time.Parse("2006-01-02T15:04:05", testReadingFilterEndDate)
	if err != nil {
		t.Errorf(err.Error())
	}

	startApiTime := ApiTime(startDate)
	endApiTime := ApiTime(endDate)

	filterRequest := new(GetReadingsByFilterRequest)
	filterRequest.ApiKey = testKey
	filterRequest.DeviceIds = append(filterRequest.DeviceIds, testDeviceId)
	filterRequest.DateStart = &startApiTime
	filterRequest.DateEnd = &endApiTime

	readingsResponse, err := client.GetReadingsByFilter(*filterRequest)
	if err != nil {
		t.Errorf(err.Error())
	}

	if readingsResponse.Status.StatusCode != 200 {
		t.Errorf(readingsResponse.Status.StatusMessage)
	}
}

func TestGetReadingsByMeterIds(t *testing.T) {
	client := NewClient(testUrl, testKey)

	filterRequest := new(GetReadingsByMeterIdsRequest)
	filterRequest.ApiKey = testKey
	filterRequest.DeviceIds = append(filterRequest.DeviceIds, testDeviceId)

	readingsResponse, err := client.GetReadingsByMeterIds(*filterRequest)
	if err != nil {
		t.Errorf(err.Error())
	}

	if readingsResponse.Status.StatusCode != 200 {
		t.Errorf(readingsResponse.Status.StatusMessage)
	}
}

func TestGetReadingsByReadingIds(t *testing.T) {
	client := NewClient(testUrl, testKey)

	filterRequest := new(GetReadingsByReadingIdsRequest)
	filterRequest.ApiKey = testKey
	filterRequest.ReadingIds = append(filterRequest.ReadingIds, testReadingId)

	readingsResponse, err := client.GetReadingsByReadingIds(*filterRequest)
	if err != nil {
		t.Errorf(err.Error())
	}

	if readingsResponse.Status.StatusCode != 200 {
		t.Errorf(readingsResponse.Status.StatusMessage)
	}
}
