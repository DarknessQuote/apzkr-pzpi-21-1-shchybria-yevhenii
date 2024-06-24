package network

import (
	"bytes"
	"devquest-iot/management"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type RequestSender struct {
	connection *management.HttpConnection
	config *management.DeviceConfig
}

func NewRequestSender(conn *management.HttpConnection, conf *management.DeviceConfig) *RequestSender {
	return &RequestSender{connection: conn, config: conf}
}

func (r *RequestSender) RegisterOwner(username string, password string) (*RegisterOwnerResponse, error) {
	jsonBody := []byte(fmt.Sprintf(`{
		"device_type": "%s",
		"username": "%s",
		"password": "%s"
	}`, r.config.DeviceSettings.Type, username, password))

	bodyReader := bytes.NewReader(jsonBody)
	requestURL := fmt.Sprintf("%s/measure/add-owner?deviceID=%s", r.config.ConnectionSettings.ServerHost, r.config.DeviceSettings.ID)

	req, err := http.NewRequest(http.MethodPut, requestURL, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := r.connection.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var resBody RegisterOwnerResponse
	err = json.NewDecoder(res.Body).Decode(&resBody)
	if err != nil {
		return nil, err
	}

	return &resBody, nil
}

func (r *RequestSender) SendMeasurement(value float64) (*AddMeasurementResponse, error) {
	jsonBody := []byte(fmt.Sprintf(`{
		"measured_at": "%s",
		"value": %.2f
	}`, time.Now().Format(time.RFC3339), value))

	bodyReader := bytes.NewReader(jsonBody)
	requestURL := fmt.Sprintf("%s/measure?deviceID=%s", r.config.ConnectionSettings.ServerHost, r.config.DeviceSettings.ID)

	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := r.connection.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var resBody AddMeasurementResponse
	err = json.NewDecoder(res.Body).Decode(&resBody)
	if err != nil {
		return nil, err
	}

	return &resBody, nil
}