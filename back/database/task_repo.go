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
		INSERT INTO tasks (id_user, name, input_data, output_data, description, lvl)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	return r.DB.QueryRowContext(ctx, query,
		task.ID_USER,
		task.Name,
		task.Input_data,
		task.Output_data,
		task.Description,
		task.Lvl,
	).Scan(&task.ID)
}
func (r *TaskRepository) GetTaskByName(ctx context.Context, name string) (*models.Task, error) {
	task := &models.Task{}
	query := `
		SELECT id, id_user, name, input_data, output_data, description, lvl
		FROM tasks WHERE name = $1
	`
	row := r.DB.QueryRowContext(ctx, query, name)
	err := row.Scan(&task.ID, &task.ID_USER, &task.Name, &task.Input_data, &task.Output_data, &task.Description, &task.Lvl)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return task, err
}
func (r *TaskRepository) UpdateTask(ctx context.Context, task *models.Task) error {
	query := `
		UPDATE tasks
		SET input_data = $1, output_data = $2, description = $3, lvl = $4 
		WHERE name = $5
	`

	_, err := r.DB.ExecContext(ctx, query,
		task.Input_data,
		task.Output_data,
		task.Description,
		task.Lvl,
		task.Name,
	)

	return err
}
