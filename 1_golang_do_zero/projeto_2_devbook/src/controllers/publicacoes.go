package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yuricapella/Go-Learning/1_golang_do_zero/projeto_2_devbook/src/autenticacao"
	"github.com/yuricapella/Go-Learning/1_golang_do_zero/projeto_2_devbook/src/banco"
	"github.com/yuricapella/Go-Learning/1_golang_do_zero/projeto_2_devbook/src/modelos"
	"github.com/yuricapella/Go-Learning/1_golang_do_zero/projeto_2_devbook/src/repositorios"
	"github.com/yuricapella/Go-Learning/1_golang_do_zero/projeto_2_devbook/src/respostas"
)

func CriarPublicacao(responseWriter http.ResponseWriter, request *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(request)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusUnauthorized, erro)
		return
	}

	corpoRequisicao, erro := io.ReadAll(request.Body)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacao modelos.Publicacao
	if erro = json.Unmarshal(corpoRequisicao, &publicacao); erro != nil {
		respostas.Erro(responseWriter, http.StatusBadRequest, erro)
		return
	}

	publicacao.AutorID = usuarioID

	if erro = publicacao.Preparar(); erro != nil {
		respostas.Erro(responseWriter, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacao.ID, erro = repositorio.Criar(publicacao)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(responseWriter, http.StatusCreated, publicacao)
}

func BuscarPublicacoes(responseWriter http.ResponseWriter, request *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(request)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusUnauthorized, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacoes, erro := repositorio.Buscar(usuarioID)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(responseWriter, http.StatusOK, publicacoes)
}

func BuscarPublicacaoPorID(responseWriter http.ResponseWriter, request *http.Request) {
	parametros := mux.Vars(request)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacao, erro := repositorio.BuscarPorID(publicacaoID)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(responseWriter, http.StatusOK, publicacao)
}

func AtualizarPublicacao(responseWriter http.ResponseWriter, request *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(request)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(request)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacaoSalvaNoBanco, erro := repositorio.BuscarPorID(publicacaoID)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}

	if publicacaoSalvaNoBanco.AutorID != usuarioID {
		respostas.Erro(responseWriter, http.StatusForbidden, errors.New("Não é possível atualizar uma publicação que não é sua"))
		return
	}

	corpoRequisicao, erro := io.ReadAll(request.Body)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacao modelos.Publicacao
	if erro = json.Unmarshal(corpoRequisicao, &publicacao); erro != nil {
		respostas.Erro(responseWriter, http.StatusBadRequest, erro)
		return
	}

	if erro = publicacao.Preparar(); erro != nil {
		respostas.Erro(responseWriter, http.StatusBadRequest, erro)
		return
	}

	if erro = repositorio.Atualizar(publicacaoID, publicacao); erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(responseWriter, http.StatusNoContent, nil)
}

func DeletarPublicacao(responseWriter http.ResponseWriter, request *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(request)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(request)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacaoSalvaNoBanco, erro := repositorio.BuscarPorID(publicacaoID)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}

	if publicacaoSalvaNoBanco.AutorID != usuarioID {
		respostas.Erro(responseWriter, http.StatusForbidden, errors.New("Não é possível deletar uma publicação que não é sua"))
		return
	}

	erro = repositorio.Deletar(publicacaoID)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(responseWriter, http.StatusNoContent, nil)
}

// BuscarPublicacoesPorUsuario - busca todas as publicações de um usuário
func BuscarPublicacoesPorUsuario(responseWriter http.ResponseWriter, request *http.Request) {
	parametros := mux.Vars(request)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacoes, erro := repositorio.BuscarPorUsuario(usuarioID)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(responseWriter, http.StatusOK, publicacoes)
}

// CurtirPublicacao - adiciona uma curtida na publicação
func CurtirPublicacao(responseWriter http.ResponseWriter, request *http.Request) {
	parametros := mux.Vars(request)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	erro = repositorio.Curtir(publicacaoID)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(responseWriter, http.StatusNoContent, nil)
}

// DescurtirPublicacao - remove uma curtida da publicação
func DescurtirPublicacao(responseWriter http.ResponseWriter, request *http.Request) {
	parametros := mux.Vars(request)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	erro = repositorio.Descurtir(publicacaoID)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(responseWriter, http.StatusNoContent, nil)
}
