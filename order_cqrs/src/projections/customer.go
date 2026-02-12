package projections

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/yuricapella/Go-Learning/order_cqrs/src/customer/events"
	"github.com/yuricapella/Go-Learning/order_cqrs/src/customer/repositories"
	customerUtils "github.com/yuricapella/Go-Learning/order_cqrs/src/customer/utils"
	"github.com/yuricapella/Go-Learning/order_cqrs/src/eventbus"
)

// ConsumeCustomerCreatedEvent sets up consumer for customer created events
func ConsumeCustomerCreatedEvent(ctx context.Context, customerMongoDBRepository *repositories.MongoDBReadRepository) error {

	handler := func(messageBody []byte) error {

		var createdEvent events.Created
		if err := json.Unmarshal(messageBody, &createdEvent); err != nil {
			return fmt.Errorf("failed to unmarshal customer created event: %w", err)
		}

		if err := customerMongoDBRepository.InsertCreatedEvent(ctx, createdEvent); err != nil {
			return fmt.Errorf("failed to insert customer created event into MongoDB: %w", err)
		}

		fmt.Printf("Customer created event processed: ID=%d, Name=%s, Email=%s\n", createdEvent.ID, createdEvent.Name, createdEvent.Email)
		return nil
	}

	return eventbus.ConsumeEvent(customerUtils.QueueCustomerCreated, handler)
}
