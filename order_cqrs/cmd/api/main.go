package main

import (
	"fmt"
	"log"

	"github.com/yuricapella/Go-Learning/order_cqrs/src/config"
	"github.com/yuricapella/Go-Learning/order_cqrs/src/database"
	"github.com/yuricapella/Go-Learning/order_cqrs/src/eventbus"
	"github.com/yuricapella/Go-Learning/order_cqrs/src/router"
)

func main() {
	fmt.Println("Order CQRS API")
	config.Load()
	fmt.Println("Running API on port", config.APIPort)

	_, rabbitChannel, err := database.ConnectRabbitMQ()
	if err != nil {
		log.Fatal(err)
	}
	eventbus.SetChannel(rabbitChannel)

	router := router.SetupRouter()

	log.Fatal(router.Run(fmt.Sprintf(":%d", config.APIPort)))
}
