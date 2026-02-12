package eventbus

import (
	"encoding/json"
	"errors"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

var rabbitMQChannel *amqp.Channel

func SetChannel(channel *amqp.Channel) {
	rabbitMQChannel = channel
}

func PublishEvent(queue string, event interface{}) error {
	fmt.Println("Publishing event to queue", queue)

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

// ConsumeEvent sets up a consumer for a RabbitMQ queue
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
				fmt.Printf("Error processing message from queue %s: %v\n", queue, err)
			}
		}
	}()

	return nil
}
