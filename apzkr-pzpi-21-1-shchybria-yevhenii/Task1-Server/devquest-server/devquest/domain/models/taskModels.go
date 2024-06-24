package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type (
	GetTaskDTO struct {
		ID uuid.UUID `json:"id"`
		Name string `json:"name"`
		Description string `json:"description"`
		Points int `json:"points"`
		ExpectedTime time.Time `json:"expected_time"`
		AcceptedTime sql.NullTime `json:"accepted_time"`
		CompletedTime sql.NullTime `json:"completed_time"`
		CategoryID uuid.UUID `json:"category_id"`
		CategoryName string `json:"category_name"`
		StatusID uuid.UUID `json:"status_id"`
		StatusName string `json:"status_name"`
		ProjectID uuid.UUID `json:"project_id"`
		DeveloperID uuid.UUID `json:"developer_id"`
	}

	CreateTaskDTO struct {
		Name         string  `json:"name"`
		Description  string  `json:"description"`
		Points       int     `json:"points"`
		ExpectedTime time.Time `json:"expected_time"`
		CategoryID uuid.UUID `json:"category_id"`
		ProjectID uuid.UUID `json:"project_id"`
	}

	UpdateTaskDTO struct {
		Name         string  `json:"name"`
		Description  string  `json:"description"`
		Points       int     `json:"points"`
		ExpectedTime time.Time `json:"expected_time"`
		CategoryID uuid.UUID `json:"category_id"`
	}

	AcceptTaskDTO struct {
		StatusID uuid.UUID `json:"status_id"`
		DeveloperID uuid.UUID `json:"developer_id"`
	}

	CompleteTaskDTO struct {
		StatusID uuid.UUID `json:"status_id"`
	}

	CreateTaskCategoryDTO struct {
		Name string `json:"string"`
	}
)