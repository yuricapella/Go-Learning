package crud

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yuricapella/Go-Learning/1_golang_do_zero/21_banco_de_dados/domain/usuario"
	"github.com/yuricapella/Go-Learning/1_golang_do_zero/21_banco_de_dados/mysql"
)

func BuscarUsuarios(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("Listando todos os usuários")

	db, erro := mysql.Conectar()
	if erro != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("select id, nome, email from usuarios")
	if erro != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("Erro ao buscar usuários"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario.Usuario

	for linhas.Next() {
		var usuario usuario.Usuario
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			responseWriter.WriteHeader(http.StatusInternalServerError)
			responseWriter.Write([]byte("Erro ao scanear usuário"))
			return
		}
		usuarios = append(usuarios, usuario)
	}

	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write([]byte(fmt.Sprintf("Usuários encontrados: %d\n", len(usuarios))))
	if erro := json.NewEncoder(responseWriter).Encode(usuarios); erro != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("Erro ao converter usuários para JSON"))
		return
	}

	fmt.Printf("usuarios encontrados: %+v\n", usuarios)
}

func BuscarUsuarioPorID(responseWriter http.ResponseWriter, request *http.Request) {
	parametros := mux.Vars(request)

	ID, erro := strconv.ParseInt(parametros["id"], 10, 32)
	if erro != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("Erro ao converter o ID para int"))
		return
	}

	db, erro := mysql.Conectar()
	if erro != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}
	defer db.Close()

	linha, erro := db.Query("select id, nome, email from usuarios where id = ?", ID)
	if erro != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("Erro ao buscar usuário"))
		return
	}
	defer linha.Close()

	var usuario usuario.Usuario

	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			responseWriter.WriteHeader(http.StatusInternalServerError)
			responseWriter.Write([]byte("Erro ao scanear usuário"))
			return
		}
	}

	if usuario.ID == 0 {
		responseWriter.WriteHeader(http.StatusNotFound)
		responseWriter.Write([]byte("Usuário não encontrado"))
		return
	}

	if erro := json.NewEncoder(responseWriter).Encode(usuario); erro != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("Erro ao converter usuário para JSON"))
		return
	}
}

func CriarUsuario(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("Criando um novo usuário")

	corpoRequisicao, erro := io.ReadAll(request.Body)
	if erro != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("Erro ao ler o corpo da requisição"))
		return
	}
	fmt.Println(string(corpoRequisicao))

	var usuario usuario.Usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	fmt.Println(usuario)

	db, erro := mysql.Conectar()
	if erro != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("Erro ao preparar a declaração"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("Erro ao executar a inserção"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("Erro ao obter o ID inserido"))
		return
	}

	responseWriter.WriteHeader(http.StatusCreated)
	responseWriter.Write([]byte(fmt.Sprintf("Usuário criado com sucesso com ID: %d", idInserido)))

	fmt.Printf("usuario criado: %+v\n", usuario)
}

func AtualizarUsuario(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("Atualizando um usuário")

	parametros := mux.Vars(request)

	ID, erro := strconv.ParseInt(parametros["id"], 10, 32)
	if erro != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("Erro ao converter o ID para int"))
		return
	}

	corpoRequisicao, erro := io.ReadAll(request.Body)
	if erro != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("Erro ao ler o corpo da requisição"))
		return
	}
	fmt.Println(string(corpoRequisicao))

	var usuario usuario.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		responseWriter.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, erro := mysql.Conectar()
	if erro != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if erro != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("Erro criar o statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("Erro ao executar a atualização"))
		return
	}

	responseWriter.WriteHeader(http.StatusNoContent)

	fmt.Printf("usuario atualizado: %+v\n", usuario)
}

func DeletarUsuario(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("Deletando um usuário")

	parametros := mux.Vars(request)

	ID, erro := strconv.ParseInt(parametros["id"], 10, 32)
	if erro != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("Erro ao converter o ID para int"))
		return
	}

	db, erro := mysql.Conectar()
	if erro != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		responseWriter.Write([]byte("Erro ao executar a deleção"))
		return
	}

	responseWriter.WriteHeader(http.StatusNoContent)

	fmt.Printf("usuario deletado com ID: %d\n", ID)
}
