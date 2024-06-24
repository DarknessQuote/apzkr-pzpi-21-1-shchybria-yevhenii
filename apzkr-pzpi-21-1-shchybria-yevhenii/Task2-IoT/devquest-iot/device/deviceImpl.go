package device

import (
	"devquest-iot/management"
	"errors"
	"math/rand"
	"sync"
)

type Device struct {
	Type string
}

var (
	once sync.Once
	deviceInstance *Device
)

func GetDevice(config *management.DeviceConfig) IDevice {
	once.Do(func() {
		deviceInstance = &Device{Type: config.DeviceSettings.Type}
	})

	return deviceInstance
}

func (d *Device) GetDataFromSensors() (float64, error) {
	var minOptimumRange, maxOptimumRange int

	switch d.Type {
		case "Pulse":
			minOptimumRange = 60
			maxOptimumRange = 100
		case "Illumination":
			minOptimumRange = 400
			maxOptimumRange = 500
		case "Humidity":
			minOptimumRange = 30
			maxOptimumRange = 50
		default:
			return 0, errors.New("unsupported type of sensor")
	}

	return float64(rand.Intn(maxOptimumRange - minOptimumRange) + minOptimumRange), nil
}