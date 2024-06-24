package repositories

import (
	"devquest-server/devquest/domain/entities"
	"devquest-server/devquest/domain/models"

	"github.com/google/uuid"
)

type TaskRepo interface {
	GetTaskByID(taskID uuid.UUID) (*models.GetTaskDTO, error)
	GetProjectTasks(projectID uuid.UUID) ([]*models.GetTaskDTO, error)
	AddTask(newTask entities.Task) error
	UpdateTask(taskID uuid.UUID, updatedTask models.UpdateTaskDTO) error
	DeleteTask(taskID uuid.UUID) error

	AcceptTask(taskID uuid.UUID, acceptedTask models.AcceptTaskDTO) error
	CompleteTask(taskID uuid.UUID, completedTask models.CompleteTaskDTO) error

	GetTaskCategories() ([]*entities.TaskCategory, error)
	GetTaskCategoryByID(categoryID uuid.UUID) (*entities.TaskCategory, error)
	GetTaskStatusByID(statusID uuid.UUID) (*entities.TaskStatus, error)
	GetTaskStatusByName(statusName string) (*entities.TaskStatus, error)
}