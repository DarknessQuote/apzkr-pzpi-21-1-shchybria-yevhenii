package management

import (
	"net/http"
	"time"
)

type HttpConnection struct {
	Client *http.Client
	serverHost string
}

func NewHttpConnection(config *DeviceConfig) *HttpConnection {
	tr := &http.Transport{
		MaxIdleConns: config.ConnectionSettings.MaxIdleConns,
		IdleConnTimeout: config.ConnectionSettings.IdleConnTimeout * time.Second,
		DisableCompression: config.ConnectionSettings.DisableCompression,
	}

	return &HttpConnection{
		Client: &http.Client{Transport: tr, Timeout: config.ConnectionSettings.ConnTimeout * time.Second},
		serverHost: config.ConnectionSettings.ServerHost,
	}
}