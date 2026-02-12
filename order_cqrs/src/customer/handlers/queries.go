package handlers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuricapella/Go-Learning/order_cqrs/src/customer/queries"
	"github.com/yuricapella/Go-Learning/order_cqrs/src/customer/repositories"
	customerUtils "github.com/yuricapella/Go-Learning/order_cqrs/src/customer/utils"
	"github.com/yuricapella/Go-Learning/order_cqrs/src/database"
	"github.com/yuricapella/Go-Learning/order_cqrs/src/responses"
)

// GetByID handles HTTP GET requests to retrieve a customer by ID from the read database
func GetByID(ginContext *gin.Context) {
	var query queries.GetByID
	if err := ginContext.ShouldBindUri(&query); err != nil {
		responses.Error(ginContext, http.StatusBadRequest, responses.ErrInvalidCustomerID)
		return
	}

	mongoClient, mongoDatabase, err := database.ConnectMongoDB()
	if err != nil {
		log.Printf("failed to connect to MongoDB: %v", err)
		responses.Error(ginContext, http.StatusInternalServerError, responses.ErrFailedToConnectMongo)
		return
	}
	defer mongoClient.Disconnect(context.Background())

	collection := mongoDatabase.Collection(customerUtils.CollectionCustomers)
	mongoDBReadRepository := repositories.NewMongoDBReadRepository(collection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	customerView, err := mongoDBReadRepository.GetByID(ctx, query.ID)
	if err != nil {
		if errors.Is(err, responses.ErrCustomerNotFound) {
			responses.Error(ginContext, http.StatusNotFound, responses.ErrCustomerNotFound)
			return
		}
		log.Printf("query error: %v", err)
		responses.Error(ginContext, http.StatusInternalServerError, responses.ErrInternalError)
		return
	}

	responses.JSON(ginContext, http.StatusOK, customerView)
}
