package user

import (
	"database/sql"
	"fmt"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) Repository {
	return &SQLiteRepository{db: db}
}

func (repository *SQLiteRepository) Create(newUser User) (User, error) {
	const insertUserSQL = `
		INSERT INTO users (id, name, created_at, updated_at)
		VALUES (?, ?, ?, ?)
	`
	_, err := repository.db.Exec(
		insertUserSQL, newUser.ID, newUser.Name, newUser.CreatedAt, newUser.UpdatedAt,
	)

	if err != nil {
		return User{}, fmt.Errorf("failed to insert user: %w", err)
	}
	return newUser, nil
}

func (repository *SQLiteRepository) FindById(id string) (User, error) {
	var foundUser User

	const selectUserSQL = `
		SELECT id, name, created_at, updated_at FROM users WHERE id = ?
	`

	err := repository.db.QueryRow(selectUserSQL, id).
		Scan(&foundUser.ID, &foundUser.Name, &foundUser.CreatedAt, &foundUser.UpdatedAt)

	if err != nil {

		if err == sql.ErrNoRows {
			return User{}, fmt.Errorf("user with id '%s' not found", id)
		}
		return User{}, fmt.Errorf("failed to query user: %w", err)
	}
	return foundUser, nil
}
