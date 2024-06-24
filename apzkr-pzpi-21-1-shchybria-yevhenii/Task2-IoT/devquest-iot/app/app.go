package app

import (
	"devquest-iot/device"
	"devquest-iot/management"
	"devquest-iot/network"
	"fmt"
	"log"
	"sync"
	"time"
)

type App struct {
	Config *management.DeviceConfig
	Device device.IDevice
	RequestSender network.RequestSender
	DataProcessor DataProcessor
}

func CreateAppInstance(conf *management.DeviceConfig) *App {
	device := device.GetDevice(conf)
	conn := management.NewHttpConnection(conf)
	reqSender := network.NewRequestSender(conn, conf)
	dataProcessor := NewDataProcessor(device)

	return &App{Config: conf, Device: device, RequestSender: *reqSender, DataProcessor: *dataProcessor}
}

func (a *App) Start() error {
	if a.Config.UserID == "" {
		err := a.registerDeviceOwner()
		if err != nil {
			return err
		}
	}

	var deviceError error
	ticker := time.NewTicker(a.Config.ConnectionSettings.RequestInterval * time.Second)
	quit := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for {
			select {
			case <-quit:
				wg.Done()
				return
			case <-ticker.C:
				deviceError = a.doAndSendMeasurement()
				if deviceError != nil {
					ticker.Stop()
					quit <- true
				}
			}
		}
	}()
	wg.Wait()

	return deviceError
}

func (a *App) registerDeviceOwner() error {
	var username, password string

	fmt.Print("Enter your username: ")
	_, err := fmt.Scanln(&username)
	if err != nil {
		return err
	}

	fmt.Print("Enter your password: ")
	_, err = fmt.Scanln(&password)
	if err != nil {
		return err
	}

	res, err := a.RequestSender.RegisterOwner(username, password)
	if err != nil {
		return err
	}

	a.Config.SetConfigValue("userID", res.UserID.String())

	return nil
}

func (a *App) doAndSendMeasurement() error {
	avgSensorValue, err := a.DataProcessor.GetAverageValueFromSensors()
	if err != nil {
		return err
	}

	res, err := a.RequestSender.SendMeasurement(avgSensorValue)
	if err != nil {
		return err
	}

	log.Println(res.Message)
	return nil
}