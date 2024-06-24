package postgres

import (
	"context"
	"database/sql"
	"devquest-server/devquest/domain/entities"
	"devquest-server/devquest/domain/models"
	"devquest-server/devquest/domain/repositories"
	"devquest-server/devquest/infrastructure"

	"github.com/google/uuid"
)

type TaskPostgresRepo struct {
	db infrastructure.Database
}

func NewTaskPostgresRepo(db infrastructure.Database) repositories.TaskRepo {
	return &TaskPostgresRepo{db: db}
}

func (t *TaskPostgresRepo) GetTaskByID(taskID uuid.UUID) (*models.GetTaskDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), t.db.GetDBTimeout())
	defer cancel()

	query := `
		SELECT t.id, t.name, t.description, t.points, t.expected_time, t.accepted_time, t.completed_time, tc.id, tc.name, ts.id, ts.name, t.project_id, t.developer_id
		FROM tasks t
		LEFT JOIN task_categories tc ON t.category_id = tc.id
		LEFT JOIN task_statuses ts ON t.status_id = ts.id
		WHERE t.id = $1
	`

	row := t.db.GetDB().QueryRowContext(ctx, query, taskID)

	var task models.GetTaskDTO
	err := row.Scan(&task.ID, &task.Name, &task.Description, &task.Points, &task.ExpectedTime, &task.AcceptedTime, &task.CompletedTime, &task.CategoryID, &task.CategoryName, &task.StatusID, &task.StatusName, &task.ProjectID, &task.DeveloperID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
			return nil, err
	}

	return &task, nil
}

func (t *TaskPostgresRepo) GetProjectTasks(projectID uuid.UUID) ([]*models.GetTaskDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), t.db.GetDBTimeout())
	defer cancel()

	query := `
		SELECT t.id, t.name, t.description, t.points, t.expected_time, t.accepted_time, t.completed_time, tc.id, tc.name, ts.id, ts.name, t.project_id, t.developer_id
		FROM tasks t
		LEFT JOIN task_categories tc ON t.category_id = tc.id
		LEFT JOIN task_statuses ts ON t.status_id = ts.id
		WHERE project_id = $1
		ORDER BY t.expected_time
	`

	rows, err := t.db.GetDB().QueryContext(ctx, query, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*models.GetTaskDTO
	for rows.Next() {
		var task models.GetTaskDTO

		err := rows.Scan(&task.ID, &task.Name, &task.Description, &task.Points, &task.ExpectedTime, &task.AcceptedTime, &task.CompletedTime, &task.CategoryID, &task.CategoryName, &task.StatusID, &task.StatusName, &task.ProjectID, &task.DeveloperID)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (t *TaskPostgresRepo) AddTask(newTask entities.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), t.db.GetDBTimeout())
	defer cancel()

	execute := `
		INSERT INTO tasks
		(id, name, description, points, expected_time, category_id, status_id, project_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := t.db.GetDB().ExecContext(ctx, execute, newTask.ID, newTask.Name, newTask.Description, newTask.Points, newTask.ExpectedTime, newTask.CategoryID, newTask.StatusID, newTask.ProjectID)
	if err != nil {
			return err
	}

	return nil
}

func (t *TaskPostgresRepo) UpdateTask(taskID uuid.UUID, updatedTask models.UpdateTaskDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), t.db.GetDBTimeout())
	defer cancel()

	execute := `
		UPDATE tasks
		SET name = $1, description = $2, points = $3, expected_time = $4, category_id = $5
		WHERE id = $6
	`

	_, err := t.db.GetDB().ExecContext(ctx, execute, updatedTask.Name, updatedTask.Description, updatedTask.Points, updatedTask.ExpectedTime, updatedTask.CategoryID, taskID)
	if err != nil {
		return err
	}

	return nil
}

func (t *TaskPostgresRepo) DeleteTask(taskID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), t.db.GetDBTimeout())
	defer cancel()

	execute := `
		DELETE FROM tasks
		WHERE id = $1
	`

	_, err := t.db.GetDB().ExecContext(ctx, execute, taskID)
	if err != nil {
		return err
	}

	return nil
}

func (t *TaskPostgresRepo) AcceptTask(taskID uuid.UUID, acceptedTask models.AcceptTaskDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), t.db.GetDBTimeout())
	defer cancel()

	execute := `
		UPDATE tasks
		SET accepted_time = now(), status_id = $1, developer_id = $2
		WHERE id = $3
	`

	_, err := t.db.GetDB().ExecContext(ctx, execute, acceptedTask.StatusID, acceptedTask.DeveloperID, taskID)
	if err != nil {
		return err
	}

	return nil
}

func (t *TaskPostgresRepo) CompleteTask(taskID uuid.UUID, completedTask models.CompleteTaskDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), t.db.GetDBTimeout())
	defer cancel()

	execute := `
		UPDATE tasks
		SET completed_time = now(), status_id = $1
		WHERE id = $2
	`

	_, err := t.db.GetDB().ExecContext(ctx, execute, completedTask.StatusID, taskID)
	if err != nil {
		return err
	}

	return nil
}

func (t *TaskPostgresRepo) GetTaskCategories() ([]*entities.TaskCategory, error) {
	ctx, cancel := context.WithTimeout(context.Background(), t.db.GetDBTimeout())
	defer cancel()

	query := `
		SELECT id, name
		FROM task_categories
	`

	rows, err := t.db.GetDB().QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*entities.TaskCategory
	for rows.Next() {
		var category entities.TaskCategory

		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}

		categories = append(categories, &category)
	}

	return categories, nil
}

func (t *TaskPostgresRepo) GetTaskCategoryByID(categoryID uuid.UUID) (*entities.TaskCategory, error) {
	ctx, cancel := context.WithTimeout(context.Background(), t.db.GetDBTimeout())
	defer cancel()

	query := `
		SELECT id, name
		FROM task_categories
		WHERE id = $1
	`

	row := t.db.GetDB().QueryRowContext(ctx, query, categoryID)

	var taskCategory entities.TaskCategory
	err := row.Scan(&taskCategory.ID, &taskCategory.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
			return nil, err
	}

	return &taskCategory, nil
}

func (t *TaskPostgresRepo) GetTaskStatusByID(statusID uuid.UUID) (*entities.TaskStatus, error) {
	ctx, cancel := context.WithTimeout(context.Background(), t.db.GetDBTimeout())
	defer cancel()

	query := `
		SELECT id, name
		FROM task_statuses
		WHERE id = $1
	`

	row := t.db.GetDB().QueryRowContext(ctx, query, statusID)

	var taskStatus entities.TaskStatus
	err := row.Scan(&taskStatus.ID, &taskStatus.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
			return nil, err
	}

	return &taskStatus, nil
}

func (t *TaskPostgresRepo) GetTaskStatusByName(statusName string) (*entities.TaskStatus, error) {
	ctx, cancel := context.WithTimeout(context.Background(), t.db.GetDBTimeout())
	defer cancel()

	query := `
		SELECT id, name
		FROM task_statuses
		WHERE name = $1
	`

	row := t.db.GetDB().QueryRowContext(ctx, query, statusName)

	var taskStatus entities.TaskStatus
	err := row.Scan(&taskStatus.ID, &taskStatus.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
			return nil, err
	}

	return &taskStatus, nil
}