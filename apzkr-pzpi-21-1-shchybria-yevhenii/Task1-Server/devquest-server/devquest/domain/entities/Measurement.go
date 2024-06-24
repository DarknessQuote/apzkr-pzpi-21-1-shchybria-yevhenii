package entities

import (
	"time"

	"github.com/google/uuid"
)

type (
	MeasurementDevice struct {
		ID uuid.UUID `json:"id"`
		TypeID uuid.UUID `json:"type_id"`
		OwnerID uuid.UUID `json:"owner_id"`
	}

	Measurement struct {
		ID uuid.UUID `json:"id"`
		DeviceID uuid.UUID `json:"device_id"`
		MeasuredAt time.Time `json:"measured_at"`
		Value float64 `json:"value"`
	}

	MeasurementType struct {
		ID uuid.UUID `json:"id"`
		Name string `json:"name"`
		MinimumNorm float64 `json:"minimum_norm"`
		MaximumNorm float64 `json:"maximum_norm"`
	}
)