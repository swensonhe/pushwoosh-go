package pushwoosh

import (
	"net/http"
	"bytes"
	"encoding/json"
	"strings"
)

// Client represents a pushwoosh client.
type Client struct {
	httpClient *http.Client

	// applicationCode is the Pushwoosh application code.
	applicationCode string

	// applicationToken is the Pushwoosh application token.
	applicationToken string
}

type ContentAvailability int

type SilentSetting int

const (
	baseUrl = "https://cp.pushwoosh.com/json/1.3"
	registerDeviceEndpoint = "/registerDevice"
	createMessageEndpoint = "/createMessage"
	unregisterDeviceEndpoint = "/unregisterDevice"
)

const (
	NoContentAvailable ContentAvailability = iota
	ContentAvailable
)

const (
	SilentSettingOff SilentSetting = iota
	SilentSettingOn
)

// Confirm that Client implements the Service interface.
var _ Service = &Client{}

func NewClient(applicationCode, applicationToken string) *Client {
	return &Client{
		httpClient: &http.Client{},
		applicationCode: applicationCode,
		applicationToken: applicationToken,
	}
}

// RegisterDevice registers a device with Pushwoosh.
func (c *Client) RegisterDevice(device *Device) error {
	url := baseUrl + registerDeviceEndpoint

	device.Application = c.applicationCode
	device.PushToken = strings.ToLower(device.PushToken)

	requestObject := &Request{
		Request: device,
	}

	data, err := json.Marshal(requestObject)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	_, err = c.httpClient.Do(request)
	return err
}

// Unregister device unregisters a device.
func (c *Client) UnregisterDevice(device *Device) error {
	url := baseUrl + unregisterDeviceEndpoint

	if device.Hwid == "" {
		return ErrHwidRequired
	}

	device.Application = c.applicationCode

	requestObject := &Request{
		Request: device,
	}

	data, err := json.Marshal(requestObject)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	_, err = c.httpClient.Do(request)
	return err
}

// CreateMessage creates a Pushwoosh message.
func (c *Client) CreateMessage(message *Message) error {
	url := baseUrl + createMessageEndpoint

	message.Application = c.applicationCode
	message.Auth = c.applicationToken

	// convert devices to lowercase strings
	for _, notification := range message.Notifications {
		devices := []string{}
		for _, device := range notification.Devices {
			devices = append(devices, strings.ToLower(device))
		}
		notification.Devices = devices
	}

	requestObject := &Request{
		Request: message,
	}

	data, err := json.Marshal(requestObject)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	_, err = c.httpClient.Do(request)

	return err
}
