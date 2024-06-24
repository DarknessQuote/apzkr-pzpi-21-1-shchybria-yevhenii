package entities

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type (
	Task struct {
		ID uuid.UUID `json:"id"`
		Name string `json:"name"`
		Description string `json:"description"`
		Points int `json:"points"`
		ExpectedTime time.Time `json:"expected_time"`
		AcceptedTime sql.NullTime `json:"accepted_time"`
		CompletedTime sql.NullTime `json:"completed_time"`
		CategoryID uuid.UUID `json:"category_id"`
		StatusID uuid.UUID `json:"status_id"`
		ProjectID uuid.UUID `json:"project_id"`
		DeveloperID uuid.UUID `json:"developer_id"`
	}

	TaskCategory struct {
		ID uuid.UUID `json:"id"`
		Name string `json:"name"`
	}

	TaskStatus struct {
		ID uuid.UUID `json:"id"`
		Name string `json:"status"`
	}
)