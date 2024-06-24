package network

import "github.com/google/uuid"

type (
	RegisterOwnerResponse struct {
		UserID uuid.UUID `json:"user_id"`
	}

	AddMeasurementResponse struct {
		Message string `json:"message"`
	}
)