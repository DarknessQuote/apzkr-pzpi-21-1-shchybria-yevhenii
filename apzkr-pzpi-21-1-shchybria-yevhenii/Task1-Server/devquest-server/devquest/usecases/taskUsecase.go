package usecases

import (
	"devquest-server/devquest/domain/entities"
	"devquest-server/devquest/domain/models"
	"devquest-server/devquest/domain/repositories"
	"errors"

	"github.com/google/uuid"
)

type TaskUsecase struct {
	taskRepo repositories.TaskRepo
	projectRepo repositories.ProjectRepo
	userRepo repositories.UserRepo
}

func NewTaskUsecase(tRepo repositories.TaskRepo, pRepo repositories.ProjectRepo, uRepo repositories.UserRepo) *TaskUsecase {
	return &TaskUsecase{taskRepo: tRepo, projectRepo: pRepo, userRepo: uRepo}
}

func (t *TaskUsecase) GetProjectTasks(projectID uuid.UUID) ([]*models.GetTaskDTO, error) {
	existingProject, err := t.projectRepo.GetProjectByID(projectID)
	if err != nil {
		return nil, err
	}
	if existingProject == nil {
		return nil, errors.New("project does not exist")
	}

	projectTasks, err := t.taskRepo.GetProjectTasks(projectID)
	if err != nil {
		return nil, err
	}
	
	return projectTasks, nil
}

func (t *TaskUsecase) CreateNewTask(newTask models.CreateTaskDTO) error {
	existingProject, err := t.projectRepo.GetProjectByID(newTask.ProjectID)
	if err != nil {
		return err
	}
	if existingProject == nil {
		return errors.New("project does not exist")
	}

	taskCategory, err := t.taskRepo.GetTaskCategoryByID(newTask.CategoryID)
	if err != nil {
		return err
	}
	if taskCategory == nil {
		return errors.New("category does not exist")
	}

	toDoStatus, err := t.taskRepo.GetTaskStatusByName("To do")
	if err != nil {
		return err
	}
	if toDoStatus == nil {
		return errors.New("error getting status")
	}

	addTask := entities.Task{
		ID: uuid.New(),
		Name: newTask.Name,
		Description: newTask.Description,
		Points: newTask.Points,
		ExpectedTime: newTask.ExpectedTime,
		CategoryID: newTask.CategoryID,
		StatusID: toDoStatus.ID,
		ProjectID: newTask.ProjectID,
	}

	err = t.taskRepo.AddTask(addTask)
	if err != nil {
		return err
	}

	return nil
}

func (t *TaskUsecase) UpdateTask(taskID uuid.UUID, updatedTask models.UpdateTaskDTO) error {
	taskToUpdate, err := t.taskRepo.GetTaskByID(taskID)
	if err != nil {
		return err
	}
	if taskToUpdate == nil {
		return errors.New("task does not exist")
	}

	err = t.taskRepo.UpdateTask(taskID, updatedTask)
	if err != nil {
		return err
	}

	return nil
}

func (t *TaskUsecase) DeleteTask(taskID uuid.UUID) error {
	err := t.taskRepo.DeleteTask(taskID)
	return err
}

func (t *TaskUsecase) AcceptTask(taskID uuid.UUID, developerID uuid.UUID) error {
	task, err := t.taskRepo.GetTaskByID(taskID)
	if err != nil {
		return err
	}
	if task == nil {
		return errors.New("task does not exist")
	}

	isDeveloperOnTheSameProject, err := t.projectRepo.CheckDeveloperOnProject(task.ProjectID, developerID)
	if err != nil {
		return err
	}
	if !isDeveloperOnTheSameProject {
		return errors.New("developer is not on this project")
	}

	if task.StatusName != "To do" {
		return errors.New("task is already in progress or done")
	}

	doingStatus, err := t.taskRepo.GetTaskStatusByName("Doing")
	if err != nil {
		return err
	}

	acceptedTask := models.AcceptTaskDTO{
		StatusID: doingStatus.ID,
		DeveloperID: developerID,
	}

	err = t.taskRepo.AcceptTask(taskID, acceptedTask)
	if err != nil {
		return err
	}

	return nil
}

func (t *TaskUsecase) CompleteTask(taskID uuid.UUID, developerID uuid.UUID) error {
	task, err := t.taskRepo.GetTaskByID(taskID)
	if err != nil {
		return err
	}
	if task == nil {
		return errors.New("task does not exist")
	}

	if task.StatusName != "Doing" {
		return errors.New("task is not in progress")
	}

	if task.DeveloperID != developerID {
		return errors.New("task cannot be completed by another developer")
	}

	doneStatus, err := t.taskRepo.GetTaskStatusByName("Done")
	if err != nil {
		return err
	}

	completedTask := models.CompleteTaskDTO{
		StatusID: doneStatus.ID,
	}

	err = t.taskRepo.CompleteTask(taskID, completedTask)
	if err != nil {
		return err
	}

	err = t.projectRepo.UpdateDeveloperProjectPoints(task.ProjectID, developerID, task.Points)
	if err != nil {
		return err
	}

	return nil
}

func (t *TaskUsecase) GetTaskCategories() ([]*entities.TaskCategory, error) {
	categories, err := t.taskRepo.GetTaskCategories()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (t *TaskUsecase) GetTaskCategoryByID(categoryID uuid.UUID) (*entities.TaskCategory, error) {
	category, err := t.taskRepo.GetTaskCategoryByID(categoryID)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (t *TaskUsecase) GetTaskStatusByID(statusID uuid.UUID) (*entities.TaskStatus, error) {
	status, err := t.taskRepo.GetTaskStatusByID(statusID)
	if err != nil {
		return nil, err
	}

	return status, nil
}