package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yuricapella/Go-Learning/order_cqrs/src/config"
)

// ConnectMySQL opens connection with MySQL database and returns it
func ConnectMySQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.MySQLConnectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	fmt.Println("Connected to MySQL")

	return db, nil
}
