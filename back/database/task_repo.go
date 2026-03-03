package database

import (
	"context"
	"database/sql"
	"ped_poject/models"
)

type TaskRepository struct {
	DB *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) CreateTask(ctx context.Context, task *models.Task) error {
	query := `
		INSERT INTO tasks (name, description, lvl)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	return r.DB.QueryRowContext(ctx, query,
		task.Name,
		task.Description,
		task.Lvl,
	).Scan(&task.ID)
}
func (r *TaskRepository) GetTaskByName(ctx context.Context, name string) (*models.Task, error) {
	task := &models.Task{}
	query := `
		SELECT id, name, description, lvl
		FROM tasks WHERE name = $1
	`
	row := r.DB.QueryRowContext(ctx, query, name)
	err := row.Scan(&task.ID, &task.Name, &task.Description, &task.Lvl)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return task, err
}
func (r *TaskRepository) UpdateTask(ctx context.Context, task *models.Task) error {
	query := `
		UPDATE tasks
		SET description = $1, lvl = $2
		WHERE name = $3
	`

	_, err := r.DB.ExecContext(ctx, query,
		task.Description,
		task.Lvl,
		task.Name,
	)

	return err
}
