package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yuricapella/Go-Learning/1_golang_do_zero/projeto_2_devbook/src/autenticacao"
	"github.com/yuricapella/Go-Learning/1_golang_do_zero/projeto_2_devbook/src/respostas"
)

/*
Pacote que serve para ser utilizado pelas rotas, ao inves de repetir codigo em todas as rotas, elas chamam o middleware

*/

// Logger - escreve as informações no terminal
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {

	return func(responseWriter http.ResponseWriter, request *http.Request) {

		fmt.Printf("\n %s %s %s\n", request.Method, request.RequestURI, request.Host)
		proximaFuncao(responseWriter, request)
	}
}

// Autenticar - verifica se o usuario fazendo a requisição esta autenticado e se estiver retorna a função recebida
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {

	return func(responseWriter http.ResponseWriter, request *http.Request) {

		if erro := autenticacao.ValidarToken(request); erro != nil {

			respostas.Erro(responseWriter, http.StatusUnauthorized, erro)
			return
		}
		proximaFuncao(responseWriter, request)
	}
}

// VerificarUsuario - verifica se o usuarioID do token corresponde ao usuarioID da URL
func VerificarUsuario(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {

		usuarioIDNoToken, erro := autenticacao.ExtrairUsuarioID(request)
		if erro != nil {
			respostas.Erro(responseWriter, http.StatusUnauthorized, erro)
			return
		}

		parametros := mux.Vars(request)
		usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
		if erro != nil {
			respostas.Erro(responseWriter, http.StatusBadRequest, errors.New("ID de usuário inválido"))
			return
		}

		if usuarioID != usuarioIDNoToken {
			respostas.Erro(responseWriter, http.StatusForbidden, errors.New("Não é possível realizar esta operação em outro usuário"))
			return
		}

		proximaFuncao(responseWriter, request)
	}
}
