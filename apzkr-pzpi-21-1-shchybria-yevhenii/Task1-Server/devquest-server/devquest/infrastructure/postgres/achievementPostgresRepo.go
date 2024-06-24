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

type AchievementPostgresRepo struct {
	db infrastructure.Database
}

func NewAchievementPostgresRepo(db infrastructure.Database) repositories.AchievementRepo {
	return &AchievementPostgresRepo{db: db}
}

func (a *AchievementPostgresRepo) GetAchievementByID(achievementID uuid.UUID) (*entities.Achievement, error) {
	ctx, cancel := context.WithTimeout(context.Background(), a.db.GetDBTimeout())
	defer cancel()

	query := `
		SELECT id, name, description, points, project_id
		FROM achievements
		WHERE id = $1
	`

	row := a.db.GetDB().QueryRowContext(ctx, query, achievementID)

	var achievement entities.Achievement
	err := row.Scan(&achievement.ID, &achievement.Name, &achievement.Description, &achievement.Points, &achievement.ProjectID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
			return nil, err
	}

	return &achievement, nil
}

func (a *AchievementPostgresRepo) GetProjectAchievements(projectID uuid.UUID) ([]*entities.Achievement, error) {
	ctx, cancel := context.WithTimeout(context.Background(), a.db.GetDBTimeout())
	defer cancel()

	query := `
		SELECT id, name, description, points, project_id
		FROM achievements
		WHERE project_id = $1
		ORDER BY name
	`

	rows, err := a.db.GetDB().QueryContext(ctx, query, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var achievements []*entities.Achievement
	for rows.Next() {
		var achievement entities.Achievement

		err := rows.Scan(&achievement.ID, &achievement.Name, &achievement.Description, &achievement.Points, &achievement.ProjectID)
		if err != nil {
			return nil, err
		}

		achievements = append(achievements, &achievement)
	}

	return achievements, nil
}

func (a *AchievementPostgresRepo) GetDevelopersAchievements(developerID uuid.UUID) ([]*entities.Achievement, error) {
	ctx, cancel := context.WithTimeout(context.Background(), a.db.GetDBTimeout())
	defer cancel()

	query := `
		SELECT a.id, a.name, a.description, a.points, a.project_id
		FROM achievements a
		LEFT JOIN achievements_users au ON a.id = au.achievement_id
		WHERE au.developer_id = $1
		ORDER BY a.name
	`

	rows, err := a.db.GetDB().QueryContext(ctx, query, developerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var achievements []*entities.Achievement
	for rows.Next() {
		var achievement entities.Achievement

		err := rows.Scan(&achievement.ID, &achievement.Name, &achievement.Description, &achievement.Points, &achievement.ProjectID)
		if err != nil {
			return nil, err
		}

		achievements = append(achievements, &achievement)
	}

	return achievements, nil
}

func (a *AchievementPostgresRepo) AddAchievement(newAchievement entities.Achievement) error {
	ctx, cancel := context.WithTimeout(context.Background(), a.db.GetDBTimeout())
	defer cancel()

	execute := `
		INSERT INTO achievements
		(id, name, description, points, project_id)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := a.db.GetDB().ExecContext(ctx, execute, newAchievement.ID, newAchievement.Name, newAchievement.Description, newAchievement.Points, newAchievement.ProjectID)
	if err != nil {
			return err
	}

	return nil
}

func (a *AchievementPostgresRepo) UpdateAchievement(achievementID uuid.UUID, updatedAchievement models.UpdateAchievementDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), a.db.GetDBTimeout())
	defer cancel()

	execute := `
		UPDATE achievements
		SET name = $1, description = $2, points = $3
		WHERE id = $4
	`

	_, err := a.db.GetDB().ExecContext(ctx, execute, updatedAchievement.Name, updatedAchievement.Description, updatedAchievement.Points, achievementID)
	if err != nil {
		return err
	}

	return nil
}

func (a *AchievementPostgresRepo) DeleteAchievement(achievementID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), a.db.GetDBTimeout())
	defer cancel()

	execute := `
		DELETE FROM achievements
		WHERE id = $1
	`

	_, err := a.db.GetDB().ExecContext(ctx, execute, achievementID)
	if err != nil {
		return err
	}

	return nil
}

func (a *AchievementPostgresRepo) CheckAchievementOnProject(projectID uuid.UUID, achievementID uuid.UUID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), a.db.GetDBTimeout())
	defer cancel()

	query := `
		SELECT id
		FROM achievements
		WHERE project_id = $1
	`

	row := a.db.GetDB().QueryRowContext(ctx, query, projectID)

	var achievementOnProjectID uuid.UUID
	err := row.Scan(&achievementOnProjectID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
			return false, err
	}

	return true, nil
}

func (a *AchievementPostgresRepo) GiveAchievementToDeveloper(achievementID uuid.UUID, developerID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), a.db.GetDBTimeout())
	defer cancel()

	execute := `
		INSERT INTO achievements_users
		(achievement_id, developer_id)
		VALUES ($1, $2)
	`

	_, err := a.db.GetDB().ExecContext(ctx, execute, achievementID, developerID)
	if err != nil {
			return err
	}

	return nil
}