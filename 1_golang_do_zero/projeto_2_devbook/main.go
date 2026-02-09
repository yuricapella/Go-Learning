package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yuricapella/Go-Learning/1_golang_do_zero/projeto_2_devbook/src/config"
	"github.com/yuricapella/Go-Learning/1_golang_do_zero/projeto_2_devbook/src/router"
)

/*
Gera uma string em base64 para usar como secret
func init() {
	chave := make([]byte, 64)

	if _, erro := rand.Read(chave); erro != nil {
		log.Fatal(erro)
	}

	stringBase64 := base64.StdEncoding.EncodeToString(chave)
	fmt.Println(stringBase64)
}
*/

func main() {
	fmt.Println("Projeto 2: DevBook")
	config.Carregar()
	fmt.Println("Rodando API na porta", config.Porta)
	router := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), router))
}
