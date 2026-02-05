package controllers

import (
	"fmt"
	"net/http"
)

func CriarUsuario(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("Criando usuário")

	responseWriter.WriteHeader(http.StatusCreated)
	responseWriter.Write([]byte("Usuário criado com sucesso"))

	fmt.Println("Usuário criado com sucesso")
}

func BuscarUsuarios(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("Buscando todos os usuários")

	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write([]byte("Usuários encontrados com sucesso"))

	fmt.Println("Usuários encontrados com sucesso")
}

func BuscarUsuarioPorID(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("Buscando usuário por ID")

	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write([]byte("Usuário encontrado com sucesso"))

	fmt.Println("Usuário encontrado com sucesso")
}

func AtualizarUsuario(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("Atualizando usuário")

	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write([]byte("Usuário atualizado com sucesso"))

	fmt.Println("Usuário atualizado com sucesso")
}

func DeletarUsuario(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("Deletando usuário")

	responseWriter.WriteHeader(http.StatusNoContent)
	responseWriter.Write([]byte("Usuário deletado com sucesso"))

	fmt.Println("Usuário deletado com sucesso")
}
