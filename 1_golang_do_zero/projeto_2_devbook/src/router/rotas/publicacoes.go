package rotas

import (
	"net/http"

	"github.com/yuricapella/Go-Learning/1_golang_do_zero/projeto_2_devbook/src/controllers"
)

var rotasPublicacoes = []Rota{
	{
		URI:                            "/publicacoes",
		Metodo:                         http.MethodPost,
		Funcao:                         controllers.CriarPublicacao,
		RequerAutenticacao:             true,
		RequerMesmoUsuarioNaRequisicao: false,
	},
	{
		URI:                            "/publicacoes",
		Metodo:                         http.MethodGet,
		Funcao:                         controllers.BuscarPublicacoes,
		RequerAutenticacao:             true,
		RequerMesmoUsuarioNaRequisicao: false,
	},
	{
		URI:                            "/publicacoes/{publicacaoId}",
		Metodo:                         http.MethodGet,
		Funcao:                         controllers.BuscarPublicacaoPorID,
		RequerAutenticacao:             true,
		RequerMesmoUsuarioNaRequisicao: false,
	},
	{
		URI:                            "/publicacoes/{publicacaoId}",
		Metodo:                         http.MethodPut,
		Funcao:                         controllers.AtualizarPublicacao,
		RequerAutenticacao:             true,
		RequerMesmoUsuarioNaRequisicao: false,
	},
	{
		URI:                            "/publicacoes/{publicacaoId}",
		Metodo:                         http.MethodDelete,
		Funcao:                         controllers.DeletarPublicacao,
		RequerAutenticacao:             true,
		RequerMesmoUsuarioNaRequisicao: false,
	},
	{
		URI:                            "/usuarios/{usuarioId}/publicacoes",
		Metodo:                         http.MethodGet,
		Funcao:                         controllers.BuscarPublicacoesPorUsuario,
		RequerAutenticacao:             true,
		RequerMesmoUsuarioNaRequisicao: false,
	},
	{
		URI:                            "/publicacoes/{publicacaoId}/curtir",
		Metodo:                         http.MethodPost,
		Funcao:                         controllers.CurtirPublicacao,
		RequerAutenticacao:             true,
		RequerMesmoUsuarioNaRequisicao: false,
	},
	{
		URI:                            "/publicacoes/{publicacaoId}/descurtir",
		Metodo:                         http.MethodPost,
		Funcao:                         controllers.DescurtirPublicacao,
		RequerAutenticacao:             true,
		RequerMesmoUsuarioNaRequisicao: false,
	},
}
