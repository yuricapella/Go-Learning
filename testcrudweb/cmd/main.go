package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"testcrudweb/internal/handler"
	"testcrudweb/internal/user"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	startServer()
}

func startServer() {
	db := connectToDatabase()
	createTableIfNotExists(db)
	defer db.Close()

	userService := user.NewUserService(user.NewSQLiteRepository(db))
	userHandler := handler.UserHandler{Service: userService}

	handler.Get("/", handler.Root)
	handler.Post("/users", userHandler.CreateUser)
	handler.Get("/users/get", userHandler.FindUserById)

	fmt.Println("Servidor Go em funcionamento!")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func connectToDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func createTableIfNotExists(db *sql.DB) {
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS users (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL
		)
	`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Erro ao criar tabela:", err)
	}
}
