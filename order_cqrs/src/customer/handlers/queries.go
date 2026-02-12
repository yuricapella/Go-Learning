package handlers

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuricapella/Go-Learning/order_cqrs/src/customer/queries"
	"github.com/yuricapella/Go-Learning/order_cqrs/src/customer/repositories"
	customerUtils "github.com/yuricapella/Go-Learning/order_cqrs/src/customer/utils"
	"github.com/yuricapella/Go-Learning/order_cqrs/src/database"
	"github.com/yuricapella/Go-Learning/order_cqrs/src/responses"
)

func GetByID(ginContext *gin.Context) {
	var query queries.GetByID
	if err := ginContext.ShouldBindUri(&query); err != nil {
		responses.Error(ginContext, http.StatusBadRequest, errors.New("invalid customer ID"))
		return
	}

	mongoClient, mongoDatabase, err := database.ConnectMongoDB()
	if err != nil {
		responses.Error(ginContext, http.StatusInternalServerError, errors.New("failed to connect to MongoDB"))
		return
	}
	defer mongoClient.Disconnect(context.Background())

	collection := mongoDatabase.Collection(customerUtils.CollectionCustomers)
	mongoDBReadRepository := repositories.NewMongoDBReadRepository(collection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	customer, err := mongoDBReadRepository.GetByID(ctx, query.ID)
	if err != nil {
		responses.Error(ginContext, http.StatusNotFound, errors.New("customer not found"))
		return
	}

	responses.JSON(ginContext, http.StatusOK, customer)
}
