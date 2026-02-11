package database

import (
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/yuricapella/Go-Learning/order_cqrs/src/config"
)

// ConnectRabbitMQ opens connection with RabbitMQ and returns connection and channel
// Retries connection up to 5 times with 2 second intervals
func ConnectRabbitMQ() (*amqp.Connection, *amqp.Channel, error) {
	connectionString := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		config.RabbitMQUser,
		config.RabbitMQPassword,
		config.RabbitMQHost,
		config.RabbitMQPort,
	)

	var conn *amqp.Connection
	var err error
	maxRetries := 5
	retryInterval := 2 * time.Second

	for i := 0; i < maxRetries; i++ {
		conn, err = amqp.Dial(connectionString)
		if err == nil {
			break
		}

		if i < maxRetries-1 {
			fmt.Printf("Failed to connect to RabbitMQ (attempt %d/%d), retrying in %v...\n", i+1, maxRetries, retryInterval)
			time.Sleep(retryInterval)
		}
	}

	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to RabbitMQ after %d attempts: %w", maxRetries, err)
	}

	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, nil, fmt.Errorf("failed to open RabbitMQ channel: %w", err)
	}

	fmt.Println("Connected to RabbitMQ")

	return conn, channel, nil
}
