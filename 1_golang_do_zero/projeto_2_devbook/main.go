package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yuricapella/Go-Learning/1_golang_do_zero/projeto_2_devbook/src/config"
	"github.com/yuricapella/Go-Learning/1_golang_do_zero/projeto_2_devbook/src/router"
)

func main() {
	fmt.Println("Projeto 2: DevBook")
	config.Carregar()
	fmt.Println("Rodando API na porta", config.Porta)
	router := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), router))
}
