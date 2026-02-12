package responses

import "errors"

var (
	ErrCustomerNotFound     = errors.New("customer not found")
	ErrInvalidCustomerID    = errors.New("invalid customer ID")
	ErrInvalidCommand       = errors.New("invalid command")
	ErrInternalError        = errors.New("internal error")
	ErrFailedToCreate       = errors.New("failed to create customer")
	ErrFailedToRetrieveID   = errors.New("failed to retrieve ID")
	ErrFailedToConnectMySQL = errors.New("failed to connect to MySQL")
	ErrFailedToConnectMongo = errors.New("failed to connect to MongoDB")
	ErrFailedToPublishEvent = errors.New("failed to publish event")
)
