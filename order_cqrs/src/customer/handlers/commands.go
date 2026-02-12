package handlers

import (
	"errors"
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

func Create(ginContext *gin.Context) {
	var command commands.Create
	if err := ginContext.ShouldBindJSON(&command); err != nil {
		responses.Error(ginContext, http.StatusBadRequest, errors.New("invalid command"))
		return
	}

	db, err := database.ConnectMySQL()
	if err != nil {
		responses.Error(ginContext, http.StatusInternalServerError, errors.New("failed to connect to MySQL"))
		return
	}
	defer db.Close()

	mysqlWriteRepository := repositories.NewMySQLWriteRepository(db) // use sua conexão/banco
	id, createdAt, err := mysqlWriteRepository.Create(command)
	if err != nil {
		responses.Error(ginContext, http.StatusInternalServerError, errors.New("failed to create customer"))
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
		responses.Error(ginContext, http.StatusInternalServerError, errors.New("failed to publish event"))
		return
	}

	ginContext.JSON(http.StatusCreated, gin.H{"id": event.ID, "name": event.Name, "email": event.Email})
}
