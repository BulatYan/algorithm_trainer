package database

import (
	"context"
	"database/sql"
	"ped_poject/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users (name, password_hash, email)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	return r.DB.QueryRowContext(ctx, query,
		user.Name,
		user.PasswordHash,
		user.Email,
	).Scan(&user.ID)
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}

	query := `
		SELECT id, name, password_hash, email
		FROM users WHERE email = $1
	`

	row := r.DB.QueryRowContext(ctx, query, email)
	err := row.Scan(&user.ID, &user.Name, &user.PasswordHash, &user.Email)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return user, err
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	user := &models.User{}

	query := `
		SELECT id, name, password_hash, email
		FROM users WHERE id = $1
	`

	row := r.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&user.ID, &user.Name, &user.PasswordHash, &user.Email)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return user, err
}
