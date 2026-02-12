package eventbus

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var rabbitMQChannel *amqp.Channel

// SetChannel sets the RabbitMQ channel for event publishing and consumption
func SetChannel(channel *amqp.Channel) {
	rabbitMQChannel = channel
}

// PublishEvent serializes an event to JSON and publishes it to the specified RabbitMQ queue
func PublishEvent(queue string, event interface{}) error {

	if rabbitMQChannel == nil {
		return errors.New("RabbitMQ channel not initialized")
	}

	body, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}
	_, err = rabbitMQChannel.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}
	return rabbitMQChannel.Publish("", queue, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
}

// ConsumeEvent registers a consumer for the specified queue and processes messages
// using the provided handler function in a goroutine
func ConsumeEvent(queue string, handler func([]byte) error) error {
	if rabbitMQChannel == nil {
		return errors.New("RabbitMQ channel not initialized")
	}

	_, err := rabbitMQChannel.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	messages, err := rabbitMQChannel.Consume(
		queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to register consumer: %w", err)
	}

	go func() {
		for message := range messages {
			if err := handler(message.Body); err != nil {
				log.Printf("Error processing message from queue %s: %v", queue, err)
			}
		}
	}()

	return nil
}
