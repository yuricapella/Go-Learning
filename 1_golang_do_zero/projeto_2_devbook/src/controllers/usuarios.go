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
	"github.com/yuricapella/Go-Learning/1_golang_do_zero/projeto_2_devbook/src/seguranca"
)

// verificarUsuarioNaRequisicao - verifica se o usuarioID do token corresponde ao usuarioID da URL
func verificarUsuarioNaRequisicao(responseWriter http.ResponseWriter, request *http.Request) bool {
	usuarioIDNoToken, erro := autenticacao.ExtrairUsuarioID(request)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusUnauthorized, erro)
		return false
	}

	parametros := mux.Vars(request)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusBadRequest, errors.New("ID de usuário inválido"))
		return false
	}

	if usuarioID != usuarioIDNoToken {
		respostas.Erro(responseWriter, http.StatusForbidden, errors.New("Não é possível realizar esta operação em outro usuário"))
		return false
	}

	return true
}

func CriarUsuario(responseWriter http.ResponseWriter, request *http.Request) {
	corpoRequest, erro := io.ReadAll(request.Body)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Erro(responseWriter, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar("cadastro"); erro != nil {
		respostas.Erro(responseWriter, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario.ID, erro = repositorio.Criar(usuario)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(responseWriter, http.StatusCreated, usuario)
}

func BuscarUsuarios(responseWriter http.ResponseWriter, request *http.Request) {
	nomeOuNick := request.URL.Query().Get("usuario")

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarios, erro := repositorio.Buscar(nomeOuNick)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(responseWriter, http.StatusOK, usuarios)
}

func BuscarUsuarioPorID(responseWriter http.ResponseWriter, request *http.Request) {
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

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario, erro := repositorio.BuscarPorID(usuarioID)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(responseWriter, http.StatusOK, usuario)
}

func AtualizarUsuario(responseWriter http.ResponseWriter, request *http.Request) {
	if !verificarUsuarioNaRequisicao(responseWriter, request) {
		return
	}

	parametros := mux.Vars(request)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusBadRequest, erro)
		return
	}

	corpoRequisicao, erro := io.ReadAll(request.Body)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		respostas.Erro(responseWriter, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar("edicao"); erro != nil {
		respostas.Erro(responseWriter, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	erro = repositorio.Atualizar(usuarioID, usuario)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(responseWriter, http.StatusNoContent, nil)
}

func DeletarUsuario(responseWriter http.ResponseWriter, request *http.Request) {
	if !verificarUsuarioNaRequisicao(responseWriter, request) {
		return
	}

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

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	erro = repositorio.Deletar(usuarioID)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(responseWriter, http.StatusNoContent, nil)
}

// SeguirUsuario - permite que um usuário siga outro
func SeguirUsuario(responseWriter http.ResponseWriter, request *http.Request) {
	seguidorID, erro := autenticacao.ExtrairUsuarioID(request)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(request)
	seguidoID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusBadRequest, erro)
		return
	}

	if seguidorID == seguidoID {
		respostas.Erro(responseWriter, http.StatusForbidden, errors.New("Não é possível seguir você mesmo"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	erro = repositorio.Seguir(seguidoID, seguidorID)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(responseWriter, http.StatusNoContent, nil)
}

// PararDeSeguirUsuario - permite que um usuário pare de seguir outro
func PararDeSeguirUsuario(responseWriter http.ResponseWriter, request *http.Request) {
	seguidorID, erro := autenticacao.ExtrairUsuarioID(request)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(request)
	seguidoID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusBadRequest, erro)
		return
	}

	if seguidorID == seguidoID {
		respostas.Erro(responseWriter, http.StatusForbidden, errors.New("Não é possível parar de seguir você mesmo"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	erro = repositorio.PararDeSeguir(seguidoID, seguidorID)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(responseWriter, http.StatusNoContent, nil)
}

// BuscarSeguidores - busca todos os seguidores de um usuário
func BuscarSeguidores(responseWriter http.ResponseWriter, request *http.Request) {
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

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	seguidores, erro := repositorio.BuscarSeguidores(usuarioID)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(responseWriter, http.StatusOK, seguidores)
}

// BuscarSeguidos - busca todos os usuários que um usuário está seguindo
func BuscarSeguidos(responseWriter http.ResponseWriter, request *http.Request) {
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

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuariosSeguidos, erro := repositorio.BuscarSeguidos(usuarioID)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(responseWriter, http.StatusOK, usuariosSeguidos)
}

// AtualizarSenha - atualiza a senha de um usuário
func AtualizarSenha(responseWriter http.ResponseWriter, request *http.Request) {
	if !verificarUsuarioNaRequisicao(responseWriter, request) {
		return
	}

	parametros := mux.Vars(request)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusBadRequest, erro)
		return
	}

	corpoRequisicao, erro := io.ReadAll(request.Body)

	var senha modelos.Senha
	if erro = json.Unmarshal(corpoRequisicao, &senha); erro != nil {
		respostas.Erro(responseWriter, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	senhaSalvaNoBanco, erro := repositorio.BuscarSenha(usuarioID)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}

	if erro = seguranca.VerificarSenha(senha.Atual, senhaSalvaNoBanco); erro != nil {
		respostas.Erro(responseWriter, http.StatusUnauthorized, errors.New("a senha atual não é válida"))
		return
	}

	senhaComHash, erro := seguranca.Hash(senha.Nova)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusBadRequest, erro)
		return
	}

	erro = repositorio.AtualizarSenha(usuarioID, string(senhaComHash))
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(responseWriter, http.StatusNoContent, nil)
}
