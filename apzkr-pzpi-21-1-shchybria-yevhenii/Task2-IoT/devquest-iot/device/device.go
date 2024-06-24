package device

type IDevice interface {
	GetDataFromSensors() (float64, error)
}