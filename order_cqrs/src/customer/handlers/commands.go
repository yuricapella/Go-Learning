package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuricapella/Go-Learning/order_cqrs/src/customer/commands"
	"github.com/yuricapella/Go-Learning/order_cqrs/src/customer/events"
	"github.com/yuricapella/Go-Learning/order_cqrs/src/customer/repositories"
	customerUtils "github.com/yuricapella/Go-Learning/order_cqrs/src/customer/utils"
	"github.com/yuricapella/Go-Learning/order_cqrs/src/database"
	"github.com/yuricapella/Go-Learning/order_cqrs/src/eventbus"
	"github.com/yuricapella/Go-Learning/order_cqrs/src/responses"
)

// Create handles HTTP POST requests to create a new customer
// It validates the command, persists to MySQL, publishes an event, and returns the created ID
func Create(ginContext *gin.Context) {
	var command commands.Create
	if err := ginContext.ShouldBindJSON(&command); err != nil {
		responses.Error(ginContext, http.StatusBadRequest, responses.ErrInvalidCommand)
		return
	}

	db, err := database.ConnectMySQL()
	if err != nil {
		log.Printf("failed to connect to MySQL: %v", err)
		responses.Error(ginContext, http.StatusInternalServerError, responses.ErrFailedToConnectMySQL)
		return
	}
	defer db.Close()

	mysqlWriteRepository := repositories.NewMySQLWriteRepository(db)
	id, createdAt, err := mysqlWriteRepository.Insert(command)
	if err != nil {
		log.Printf("failed to create customer: %v", err)
		if errors.Is(err, responses.ErrFailedToCreate) {
			responses.Error(ginContext, http.StatusInternalServerError, responses.ErrFailedToCreate)
		} else {
			responses.Error(ginContext, http.StatusInternalServerError, responses.ErrInternalError)
		}
		return
	}

	event := events.Created{
		ID:        id,
		Name:      command.Name,
		Email:     command.Email,
		CreatedAt: createdAt,
	}
	err = eventbus.PublishEvent(customerUtils.QueueCustomerCreated, event)
	if err != nil {
		log.Printf("failed to publish event: %v", err)
		responses.Error(ginContext, http.StatusInternalServerError, responses.ErrFailedToPublishEvent)
		return
	}

	responses.Created(ginContext, fmt.Sprintf("/customers/%d", id), id)
}
