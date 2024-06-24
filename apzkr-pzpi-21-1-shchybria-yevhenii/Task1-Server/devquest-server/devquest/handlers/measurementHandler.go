package handlers

import (
	"devquest-server/devquest/domain/models"
	"devquest-server/devquest/usecases"
	"devquest-server/devquest/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type MeasurementHttpHandler struct {
	measurementUsecase usecases.MeasurementUsecase
}

func NewMeasurementHttpHandler(mUsecase usecases.MeasurementUsecase) *MeasurementHttpHandler {
	return &MeasurementHttpHandler{measurementUsecase: mUsecase}
}

func (m *MeasurementHttpHandler) AddOwnerToDevice(w http.ResponseWriter, r *http.Request) {
	deviceID, err := uuid.Parse(r.URL.Query().Get("deviceID"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	var deviceOwnerInfo models.AddOwnerToDeviceDTO
	deviceOwnerInfo.DeviceID = deviceID
	err = utils.ReadJSON(w, r, &deviceOwnerInfo)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	userID, err := m.measurementUsecase.AddOwnerToDevice(deviceOwnerInfo)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	res := struct {
		UserID uuid.UUID `json:"user_id"`
	} {
		UserID: userID,
	}

	_ = utils.WriteJSON(w, http.StatusAccepted, res)
}

func (m *MeasurementHttpHandler) AddMeasurementResult(w http.ResponseWriter, r *http.Request) {
	deviceID, err := uuid.Parse(r.URL.Query().Get("deviceID"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	var newMeasurement models.CreateMeasurementDTO
	newMeasurement.DeviceID = deviceID
	err = utils.ReadJSON(w, r, &newMeasurement)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	err = m.measurementUsecase.AddMeasurementResult(newMeasurement)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	res := utils.JSONResponse{
		Error: false,
		Message: "measurement successfully added",
	}

	_ = utils.WriteJSON(w, http.StatusAccepted, res)
}

func (m *MeasurementHttpHandler) GetLatestMeasurementsForDeveloper(w http.ResponseWriter, r *http.Request) {
	developerID, err := uuid.Parse(chi.URLParam(r, "developer_id"))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	latestMeasurements, err := m.measurementUsecase.GetLatestMeasurementsForDeveloper(developerID)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	_ = utils.WriteJSON(w, http.StatusAccepted, latestMeasurements)
}