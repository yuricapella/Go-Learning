package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yuricapella/Go-Learning/1_golang_do_zero/20_banco_de_dados/crud"
	"github.com/yuricapella/Go-Learning/1_golang_do_zero/20_banco_de_dados/mysql"
)

func main() {
	if erro := mysql.PrepararBancoDeDados(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Banco de dados preparado com sucesso")

	router := mux.NewRouter()

	fmt.Println("Escutando na porta 8080")

	router.HandleFunc("/usuarios", crud.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", crud.BuscarUsuarioPorID).Methods(http.MethodGet)
	router.HandleFunc("/usuarios", crud.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/usuarios/{id}", crud.AtualizarUsuario).Methods(http.MethodPut)
	router.HandleFunc("/usuarios/{id}", crud.DeletarUsuario).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8080", router))

}
