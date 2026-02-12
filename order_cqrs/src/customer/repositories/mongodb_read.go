package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/yuricapella/Go-Learning/order_cqrs/src/customer/viewmodels"
	"github.com/yuricapella/Go-Learning/order_cqrs/src/responses"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoDBReadRepository struct {
	collection *mongo.Collection
}

// NewMongoDBReadRepository creates a new MongoDB read repository instance
func NewMongoDBReadRepository(collection *mongo.Collection) *MongoDBReadRepository {
	return &MongoDBReadRepository{collection: collection}
}

// Insert adds a customer view model to MongoDB with idempotency check
// If the customer already exists, it returns nil without error
func (repository *MongoDBReadRepository) Insert(ctx context.Context, customerView viewmodels.CustomerView) error {
	_, err := repository.GetByID(ctx, customerView.ID)
	if err == nil {
		return nil
	}

	if !errors.Is(err, responses.ErrCustomerNotFound) {
		return fmt.Errorf("%w: %v", responses.ErrInternalError, err)
	}

	_, insertError := repository.collection.InsertOne(ctx, customerView)
	if insertError != nil {
		return fmt.Errorf("%w: %v", responses.ErrInternalError, insertError)
	}
	return nil
}

// GetByID retrieves a customer view model from MongoDB by ID
// Returns ErrCustomerNotFound if the customer does not exist
func (repository *MongoDBReadRepository) GetByID(ctx context.Context, id int64) (viewmodels.CustomerView, error) {
	var customer viewmodels.CustomerView

	filter := bson.M{"id": id}

	err := repository.collection.FindOne(ctx, filter).Decode(&customer)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return viewmodels.CustomerView{}, responses.ErrCustomerNotFound
		}
		return viewmodels.CustomerView{}, fmt.Errorf("%w: %v", responses.ErrInternalError, err)
	}

	return customer, nil
}
