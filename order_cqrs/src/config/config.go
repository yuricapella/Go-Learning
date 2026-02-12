package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// API Port where the server will run
	APIPort = 0

	// MySQL connection string for write database
	MySQLConnectionString = ""

	// MongoDB connection string for read database
	MongoDBConnectionString = ""

	// MongoDB database name
	MongoDBDatabaseName = ""

	// RabbitMQ connection parameters
	RabbitMQHost     = ""
	RabbitMQPort     = ""
	RabbitMQUser     = ""
	RabbitMQPassword = ""
)

// Load loads all environment variables and initializes configuration
func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	APIPort, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		APIPort = 8080
	}

	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	MySQLConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlUser,
		mysqlPassword,
		mysqlHost,
		mysqlPort,
		mysqlDatabase,
	)

	mongodbHost := os.Getenv("MONGODB_HOST")
	mongodbPort := os.Getenv("MONGODB_PORT")
	MongoDBDatabaseName = os.Getenv("MONGODB_DATABASE")

	MongoDBConnectionString = fmt.Sprintf("mongodb://%s:%s",
		mongodbHost,
		mongodbPort,
	)

	RabbitMQHost = os.Getenv("RABBITMQ_HOST")
	RabbitMQPort = os.Getenv("RABBITMQ_PORT")
	RabbitMQUser = os.Getenv("RABBITMQ_USER")
	RabbitMQPassword = os.Getenv("RABBITMQ_PASSWORD")
}
