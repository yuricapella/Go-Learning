package main

import (
	"fmt"
	"log"
	"net/http"
	"testcrudweb/internal/handler"
)

func main() {
	startServer()
}

func startServer() {
	handler.Get("/", handler.Root)
	fmt.Println("Servidor Go em funcionamento!")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
