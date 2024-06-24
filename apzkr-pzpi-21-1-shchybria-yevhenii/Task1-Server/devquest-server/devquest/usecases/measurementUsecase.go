package usecases

import (
	"devquest-server/devquest/domain/entities"
	"devquest-server/devquest/domain/models"
	"devquest-server/devquest/domain/repositories"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type MeasurementUsecase struct {
	measurementRepo repositories.MeasurementRepo
	userRepo repositories.UserRepo
}

func NewMeasurementUsecase(mRepo repositories.MeasurementRepo, uRepo repositories.UserRepo) *MeasurementUsecase {
	return &MeasurementUsecase{measurementRepo: mRepo, userRepo: uRepo}
}

func (m *MeasurementUsecase) AddOwnerToDevice(deviceOwnerInfo models.AddOwnerToDeviceDTO) (uuid.UUID, error) {
	deviceType, err := m.measurementRepo.GetTypeByName(deviceOwnerInfo.DeviceType)
	if err != nil {
		return uuid.Nil, err
	}
	if deviceType == nil {
		return uuid.Nil, errors.New("device type does not exist")
	}

	device, err := m.measurementRepo.CheckForDevice(deviceOwnerInfo.DeviceID, deviceType.ID)
	if err != nil {
		return uuid.Nil, err
	}
	if !device {
		return uuid.Nil, errors.New("device does not exist")
	}

	user, err := m.userRepo.GetUserByUsername(deviceOwnerInfo.Username)
	if err != nil {
		return uuid.Nil, err
	}
	if user == nil {
		return uuid.Nil, errors.New("user does not exist")
	}
	
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(deviceOwnerInfo.Password))
	if err != nil {
		return uuid.Nil, err
	}

	isDeveloper, err := m.userRepo.CheckUserRole(user.ID, "Developer")
	if err != nil {
		return uuid.Nil, err
	}
	if !isDeveloper {
		return uuid.Nil, errors.New("can't register device as a non-developer")
	}

	err = m.measurementRepo.AddOwnerToDevice(deviceOwnerInfo.DeviceID, user.ID)
	if err != nil {
		return uuid.Nil, err
	}

	return user.ID, nil
}

func (m *MeasurementUsecase) AddMeasurementResult(measurementInfo models.CreateMeasurementDTO) error {
	device, err := m.measurementRepo.GetDeviceByID(measurementInfo.DeviceID)
	if err != nil {
		return err
	}
	if device == nil {
		return errors.New("device does not exist")
	}

	measurement := entities.Measurement{
		ID: uuid.New(),
		DeviceID: measurementInfo.DeviceID,
		MeasuredAt: measurementInfo.MeasuredAt,
		Value: measurementInfo.Value,
	}

	err = m.measurementRepo.AddMeasurementResult(measurement)
	if err != nil {
		return err
	}

	return nil
}

func (m *MeasurementUsecase) GetLatestMeasurementsForDeveloper(developerID uuid.UUID) ([]*models.GetMeasurementDTO, error) {
	latestMeasurements, err := m.measurementRepo.GetLatestMeasurementsForDeveloper(developerID)
	if err != nil {
		return nil, err
	}

	var latestMeasurementsDTO []*models.GetMeasurementDTO
	for _, measurement := range(latestMeasurements) {
		device, _ := m.measurementRepo.GetDeviceByID(measurement.DeviceID)
		deviceType, _ := m.measurementRepo.GetTypeByID(device.TypeID)

		message := "Results are normal"
		if (measurement.Value < deviceType.MinimumNorm) {
			message = "Results are below optimal levels"
		} else if (measurement.Value < deviceType.MaximumNorm) {
			message = "Results are above optimal levels"
		}

		measurementDTO := &models.GetMeasurementDTO{
			ID: measurement.ID,
			TypeName: deviceType.Name,
			MeasuredAt: measurement.MeasuredAt,
			Value: measurement.Value,
			Message: message,
		}

		latestMeasurementsDTO = append(latestMeasurementsDTO, measurementDTO)
	}

	return latestMeasurementsDTO, nil
}