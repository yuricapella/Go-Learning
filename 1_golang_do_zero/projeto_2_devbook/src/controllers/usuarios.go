package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yuricapella/Go-Learning/1_golang_do_zero/projeto_2_devbook/src/banco"
	"github.com/yuricapella/Go-Learning/1_golang_do_zero/projeto_2_devbook/src/modelos"
	"github.com/yuricapella/Go-Learning/1_golang_do_zero/projeto_2_devbook/src/repositorios"
	"github.com/yuricapella/Go-Learning/1_golang_do_zero/projeto_2_devbook/src/respostas"
)

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
