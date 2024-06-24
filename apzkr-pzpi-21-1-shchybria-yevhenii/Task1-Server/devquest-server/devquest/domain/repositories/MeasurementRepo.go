package repositories

import (
	"devquest-server/devquest/domain/entities"

	"github.com/google/uuid"
)

type MeasurementRepo interface {
	AddMeasurementResult(newMeasurement entities.Measurement) error
	GetLatestMeasurementsForDeveloper(developerID uuid.UUID) ([]*entities.Measurement, error)
	
	GetDeviceByID(deviceID uuid.UUID) (*entities.MeasurementDevice, error)
	CheckForDevice(deviceID uuid.UUID, typeID uuid.UUID) (bool, error)
	AddOwnerToDevice(deviceID uuid.UUID, ownerID uuid.UUID) error

	GetTypeByID(typeID uuid.UUID) (*entities.MeasurementType, error)
	GetTypeByName(typeName string) (*entities.MeasurementType, error)
}