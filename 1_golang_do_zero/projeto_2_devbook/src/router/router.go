package router

import (
	"github.com/gorilla/mux"
	"github.com/yuricapella/Go-Learning/1_golang_do_zero/projeto_2_devbook/src/router/rotas"
)

func Gerar() *mux.Router {
	router := mux.NewRouter()

	return rotas.Configurar(router)
}
