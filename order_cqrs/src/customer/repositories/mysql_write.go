package repositories

import (
	"database/sql"
	"time"

	"github.com/yuricapella/Go-Learning/order_cqrs/src/customer/commands"
)

type MySQLWriteRepository struct {
	db *sql.DB
}

func NewMySQLWriteRepository(db *sql.DB) *MySQLWriteRepository {
	return &MySQLWriteRepository{db: db}
}

func (repository *MySQLWriteRepository) Create(command commands.Create) (int64, time.Time, error) {
	statement, err := repository.db.Prepare("INSERT INTO customers (name, email) VALUES (?, ?)")
	if err != nil {
		return 0, time.Time{}, err
	}
	defer statement.Close()

	result, err := statement.Exec(command.Name, command.Email)
	if err != nil {
		return 0, time.Time{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, time.Time{}, err
	}

	var createdAt time.Time
	err = repository.db.QueryRow(
		"SELECT created_at FROM customers WHERE id = ?", id,
	).Scan(&createdAt)
	if err != nil {
		return 0, time.Time{}, err
	}

	return id, createdAt, nil
}
