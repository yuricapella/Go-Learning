package main

import (
	"context"
	"fmt"
	"log"

	"github.com/yuricapella/Go-Learning/order_cqrs/src/config"
	"github.com/yuricapella/Go-Learning/order_cqrs/src/database"
	"github.com/yuricapella/Go-Learning/order_cqrs/src/eventbus"
	"github.com/yuricapella/Go-Learning/order_cqrs/src/projections"
	"github.com/yuricapella/Go-Learning/order_cqrs/src/router"
)

// main initializes the application, connects to databases and message broker,
// starts projection consumers, and runs the HTTP server
func main() {
	log.Println("Order CQRS API")
	config.Load()
	log.Printf("Running API on port %d", config.APIPort)

	_, rabbitChannel, err := database.ConnectRabbitMQ()
	if err != nil {
		log.Fatal(err)
	}
	eventbus.SetChannel(rabbitChannel)

	projections.StartAllConsumers(context.Background())

	router := router.SetupRouter()

	log.Fatal(router.Run(fmt.Sprintf(":%d", config.APIPort)))
}
