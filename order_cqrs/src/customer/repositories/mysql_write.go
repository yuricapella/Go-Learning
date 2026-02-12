package repositories

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/yuricapella/Go-Learning/order_cqrs/src/customer/commands"
	"github.com/yuricapella/Go-Learning/order_cqrs/src/responses"
)

type MySQLWriteRepository struct {
	db *sql.DB
}

// NewMySQLWriteRepository creates a new MySQL write repository instance
func NewMySQLWriteRepository(db *sql.DB) *MySQLWriteRepository {
	return &MySQLWriteRepository{db: db}
}

// Insert inserts a new customer into MySQL and returns the generated ID and created_at timestamp
func (repository *MySQLWriteRepository) Insert(command commands.Create) (int64, time.Time, error) {
	statement, err := repository.db.Prepare("INSERT INTO customers (name, email) VALUES (?, ?)")
	if err != nil {
		return 0, time.Time{}, fmt.Errorf("%w: %v", responses.ErrFailedToCreate, err)
	}
	defer statement.Close()

	result, err := statement.Exec(command.Name, command.Email)
	if err != nil {
		return 0, time.Time{}, fmt.Errorf("%w: %v", responses.ErrFailedToCreate, err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, time.Time{}, fmt.Errorf("%w: %v", responses.ErrFailedToRetrieveID, err)
	}

	var createdAt time.Time
	err = repository.db.QueryRow(
		"SELECT created_at FROM customers WHERE id = ?", id,
	).Scan(&createdAt)
	if err != nil {
		return 0, time.Time{}, fmt.Errorf("%w: %v", responses.ErrFailedToCreate, err)
	}

	return id, createdAt, nil
}
