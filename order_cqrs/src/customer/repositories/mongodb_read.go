package repositories

import (
	"context"
	"fmt"

	"github.com/yuricapella/Go-Learning/order_cqrs/src/customer/events"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoDBReadRepository struct {
	collection *mongo.Collection
}

func NewMongoDBReadRepository(collection *mongo.Collection) *MongoDBReadRepository {
	return &MongoDBReadRepository{collection: collection}
}

func (repository *MongoDBReadRepository) InsertCreatedEvent(ctx context.Context, createdEvent events.Created) error {
	_, insertError := repository.collection.InsertOne(ctx, createdEvent)
	if insertError != nil {
		return fmt.Errorf("failed to insert created event: %w", insertError)
	}
	return nil
}

func (repository *MongoDBReadRepository) GetByID(ctx context.Context, id int64) (events.Created, error) {
	var customer events.Created
	filter := bson.M{"id": id}

	err := repository.collection.FindOne(ctx, filter).Decode(&customer)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return events.Created{}, fmt.Errorf("customer with ID %d not found", id)
		}
		return events.Created{}, fmt.Errorf("failed to find customer by ID: %w", err)
	}

	return customer, nil
}
