package middlewares

import (
	"fmt"
	"net/http"

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
