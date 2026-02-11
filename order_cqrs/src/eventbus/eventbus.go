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
