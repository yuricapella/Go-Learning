package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/yuricapella/Go-Learning/1_golang_do_zero/projeto_2_devbook/src/autenticacao"
	"github.com/yuricapella/Go-Learning/1_golang_do_zero/projeto_2_devbook/src/banco"
	"github.com/yuricapella/Go-Learning/1_golang_do_zero/projeto_2_devbook/src/modelos"
	"github.com/yuricapella/Go-Learning/1_golang_do_zero/projeto_2_devbook/src/repositorios"
	"github.com/yuricapella/Go-Learning/1_golang_do_zero/projeto_2_devbook/src/respostas"
	"github.com/yuricapella/Go-Learning/1_golang_do_zero/projeto_2_devbook/src/seguranca"
)

// Login - responsavel por autenticar o usuario na API
func Login(responseWriter http.ResponseWriter, request *http.Request) {
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

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarioSalvoNoBanco, erro := repositorio.BuscarPorEmail(usuario.Email)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}

	if erro = seguranca.VerificarSenha(usuario.Senha, usuarioSalvoNoBanco.Senha); erro != nil {
		respostas.Erro(responseWriter, http.StatusUnauthorized, erro)
		return
	}

	token, erro := autenticacao.CriarToken(usuarioSalvoNoBanco.ID)
	if erro != nil {
		respostas.Erro(responseWriter, http.StatusInternalServerError, erro)
		return
	}

	responseWriter.Write([]byte(token))
}
