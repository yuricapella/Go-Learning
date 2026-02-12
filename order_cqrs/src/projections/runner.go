package projections

import (
	"context"
	"log"

	"github.com/yuricapella/Go-Learning/order_cqrs/src/customer/repositories"
	customerUtils "github.com/yuricapella/Go-Learning/order_cqrs/src/customer/utils"
	"github.com/yuricapella/Go-Learning/order_cqrs/src/database"
)

type ProjectionConsumer func(context.Context) error

// StartAllConsumers initializes and starts all projection consumers as long-running goroutines
func StartAllConsumers(ctx context.Context) {
	_, mongoDatabase, err := database.ConnectMongoDB()
	if err != nil {
		log.Fatal("MongoDB connection failed: ", err)
	}
	customerCollection := mongoDatabase.Collection(customerUtils.CollectionCustomers)
	customerRepo := repositories.NewMongoDBReadRepository(customerCollection)

	for _, consumer := range listConsumers(customerRepo) {
		go func(l ProjectionConsumer) {
			if err := l(ctx); err != nil {
				log.Printf("Projection listener stopped with error: %v", err)
			}
		}(consumer)
	}
}

// listConsumers returns all projection consumers configured for the system
func listConsumers(customerRepo *repositories.MongoDBReadRepository) []ProjectionConsumer {
	return []ProjectionConsumer{
		func(ctx context.Context) error { return ConsumeCustomerCreatedEvent(ctx, customerRepo) },
	}
}
