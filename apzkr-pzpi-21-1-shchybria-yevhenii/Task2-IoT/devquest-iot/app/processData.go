package app

import (
	"devquest-iot/device"
	"log"
	"math"
	"time"
)

type DataProcessor struct {
	device device.IDevice
}

func NewDataProcessor(dev device.IDevice) *DataProcessor {
	return &DataProcessor{device: dev}
}

func (p *DataProcessor) GetAverageValueFromSensors() (float64, error) {
	var sensorData []float64

	for range make([]int, 20) {
		data, err := p.device.GetDataFromSensors()
		if err != nil {
			return 0, err
		}

		log.Printf("Data point: %.2f", data)
		sensorData = append(sensorData, data)
		time.Sleep(time.Second * 1)
	}

	avgValue := getInterquartileMean(sensorData)
	return avgValue, nil
}

func getInterquartileMean(data []float64) float64 {
	quartileSize := len(data) / 4
	firstQuartileBound := quartileSize
	lastQuartileBound := len(data) - quartileSize
	middleQuartiles := data[firstQuartileBound:lastQuartileBound]

	var valueSum float64
	for _, value := range(middleQuartiles) {
		valueSum += value
	}

	mean := valueSum / float64(len(middleQuartiles))
	
	return math.Round(mean * 100) / 100
}