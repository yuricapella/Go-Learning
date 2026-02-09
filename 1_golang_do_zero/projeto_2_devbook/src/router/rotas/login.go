package rotas

import (
	"net/http"

	"github.com/yuricapella/Go-Learning/1_golang_do_zero/projeto_2_devbook/src/controllers"
)

var rotaLogin = Rota{
	URI:                "/login",
	Metodo:             http.MethodPost,
	Funcao:             controllers.Login,
	RequerAutenticacao: false,
}
