package models

import (
	"time"

	"github.com/google/uuid"
)

type (
	CreateMeasurementDTO struct {
		DeviceID uuid.UUID `json:"device_id"`
		MeasuredAt time.Time `json:"measured_at"`
		Value float64 `json:"value"`
	}

	GetMeasurementDTO struct {
		ID uuid.UUID `json:"id"`
		TypeName string `json:"type_name"`
		MeasuredAt time.Time `json:"measured_at"`
		Value float64 `json:"value"`
		Message string `json:"message"`
	}

	AddOwnerToDeviceDTO struct {
		DeviceID uuid.UUID `json:"device_id"`
		DeviceType string `json:"device_type"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
)