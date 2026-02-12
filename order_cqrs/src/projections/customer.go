package projections

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/yuricapella/Go-Learning/order_cqrs/src/customer/events"
	"github.com/yuricapella/Go-Learning/order_cqrs/src/customer/repositories"
	customerUtils "github.com/yuricapella/Go-Learning/order_cqrs/src/customer/utils"
	"github.com/yuricapella/Go-Learning/order_cqrs/src/customer/viewmodels"
	"github.com/yuricapella/Go-Learning/order_cqrs/src/eventbus"
)

// ConsumeCustomerCreatedEvent consumes customer_created events from RabbitMQ,
// converts them to view models, and persists them to MongoDB
func ConsumeCustomerCreatedEvent(ctx context.Context, customerMongoDBRepository *repositories.MongoDBReadRepository) error {

	handler := func(messageBody []byte) error {

		var createdEvent events.Created
		if err := json.Unmarshal(messageBody, &createdEvent); err != nil {
			return fmt.Errorf("failed to unmarshal customer created event: %w", err)
		}

		customerView := viewmodels.CustomerView{
			ID:        createdEvent.ID,
			Name:      createdEvent.Name,
			Email:     createdEvent.Email,
			CreatedAt: createdEvent.CreatedAt,
		}

		if err := customerMongoDBRepository.Insert(ctx, customerView); err != nil {
			return fmt.Errorf("failed to insert customer into MongoDB: %w", err)
		}

		return nil
	}

	return eventbus.ConsumeEvent(customerUtils.QueueCustomerCreated, handler)
}
