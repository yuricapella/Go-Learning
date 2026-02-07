package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yuricapella/Go-Learning/1_golang_do_zero/projeto_2_devbook/src/middlewares"
)

// Rota representa todas as rotas da API
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool

	// Se true, a rota requer que o usuário seja o mesmo que o usuário autenticado
	RequerMesmoUsuario bool
}

// Coloca todas as rotas dentro do router
func Configurar(router *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotaLogin)

	for _, rota := range rotas {

		if rota.RequerAutenticacao {
			var handler http.HandlerFunc = rota.Funcao

			if rota.RequerMesmoUsuario {
				handler = middlewares.VerificarUsuario(handler)
			}

			router.HandleFunc(
				rota.URI,
				middlewares.Logger(middlewares.Autenticar(handler)),
			).Methods(rota.Metodo)
		} else {

			router.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}

		router.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}
	return router
}
