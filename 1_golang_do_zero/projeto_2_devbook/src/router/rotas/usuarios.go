package rotas

import (
	"net/http"

	"github.com/yuricapella/Go-Learning/1_golang_do_zero/projeto_2_devbook/src/controllers"
)

var rotasUsuarios = []Rota{
	{
		URI:                            "/usuarios",
		Metodo:                         http.MethodPost,
		Funcao:                         controllers.CriarUsuario,
		RequerAutenticacao:             false,
		RequerMesmoUsuarioNaRequisicao: false,
	},
	{
		URI:                            "/usuarios",
		Metodo:                         http.MethodGet,
		Funcao:                         controllers.BuscarUsuarios,
		RequerAutenticacao:             true,
		RequerMesmoUsuarioNaRequisicao: false,
	},
	{
		URI:                            "/usuarios/{usuarioId}",
		Metodo:                         http.MethodGet,
		Funcao:                         controllers.BuscarUsuarioPorID,
		RequerAutenticacao:             true,
		RequerMesmoUsuarioNaRequisicao: false,
	},
	{
		URI:                            "/usuarios/{usuarioId}",
		Metodo:                         http.MethodPut,
		Funcao:                         controllers.AtualizarUsuario,
		RequerAutenticacao:             true,
		RequerMesmoUsuarioNaRequisicao: true,
	},
	{
		URI:                            "/usuarios/{usuarioId}",
		Metodo:                         http.MethodDelete,
		Funcao:                         controllers.DeletarUsuario,
		RequerAutenticacao:             true,
		RequerMesmoUsuarioNaRequisicao: true,
	},
	{
		URI:                            "/usuarios/{usuarioId}/seguir",
		Metodo:                         http.MethodPost,
		Funcao:                         controllers.SeguirUsuario,
		RequerAutenticacao:             true,
		RequerMesmoUsuarioNaRequisicao: false,
	},
	{
		URI:                            "/usuarios/{usuarioId}/parar-de-seguir",
		Metodo:                         http.MethodPost,
		Funcao:                         controllers.PararDeSeguirUsuario,
		RequerAutenticacao:             true,
		RequerMesmoUsuarioNaRequisicao: false,
	},
	{
		URI:                            "/usuarios/{usuarioId}/seguidores",
		Metodo:                         http.MethodGet,
		Funcao:                         controllers.BuscarSeguidores,
		RequerAutenticacao:             true,
		RequerMesmoUsuarioNaRequisicao: false,
	},
	{
		URI:                            "/usuarios/{usuarioId}/seguindo",
		Metodo:                         http.MethodGet,
		Funcao:                         controllers.BuscarSeguidos,
		RequerAutenticacao:             true,
		RequerMesmoUsuarioNaRequisicao: false,
	},
	{
		URI:                            "/usuarios/{usuarioId}/atualizar-senha",
		Metodo:                         http.MethodPost,
		Funcao:                         controllers.AtualizarSenha,
		RequerAutenticacao:             true,
		RequerMesmoUsuarioNaRequisicao: true,
	},
}
