package repositorios

import (
	"database/sql"
	"fmt"

	"github.com/yuricapella/Go-Learning/1_golang_do_zero/projeto_2_devbook/src/modelos"
)

// representa um repositorio de usuarios
type usuarios struct {
	db *sql.DB
}

// Cria um repositorio de usuarios
func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db: db}
}

// Criar insere um usuário no banco de dados
func (repositorio usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	UltimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(UltimoIDInserido), nil
}

// Buscar - Busca todos os usuários que contenham o nome ou nick no banco de dados
func (repositorio usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)

	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome like ? or nick like ?",
		nomeOuNick, nomeOuNick,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// BuscarPorID - Busca um usuário pelo ID
func (repositorio usuarios) BuscarPorID(ID uint64) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id = ?",
		ID,
	)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

// Atualizar - Atualiza um usuário no banco de dados
func (repositorio usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {
	statement, erro := repositorio.db.Prepare(
		"update usuarios set nome = ?, nick = ?, email = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
		return erro
	}

	return nil
}

// Deletar - Deleta um usuário no banco de dados
func (repositorio usuarios) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

// BuscarPorEmail - busca um usuario por email e retorna o seu id e senha com hash
func (repositorio usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"select id, senha from usuarios where email = ?",
		email,
	)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario
	if linhas.Next() {
		if erro := linhas.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

// Seguir - permite que um usuário siga outro
func (repositorio usuarios) Seguir(seguidoID, seguidorID uint64) error {

	// insert ignore - insere o registro se não existir, caso exista, não retorna erro
	statement, erro := repositorio.db.Prepare(
		"insert ignore into seguidores (seguido_id, seguidor_id) values (?, ?)",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(seguidoID, seguidorID); erro != nil {
		return erro
	}

	return nil
}

// PararDeSeguir - permite que um usuário pare de seguir outro
func (repositorio usuarios) PararDeSeguir(seguidoID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from seguidores where seguido_id = ? and seguidor_id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(seguidoID, seguidorID); erro != nil {
		return erro
	}

	return nil
}

// BuscarSeguidores - busca todos os seguidores de um usuário
func (repositorio usuarios) BuscarSeguidores(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
	select u.id, u.nome, u.nick, u.email, u.criadoEm 
	from usuarios u inner join seguidores s on u.id = s.seguidor_id where s.seguido_id = ?`, usuarioID,
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var seguidores []modelos.Usuario

	for linhas.Next() {
		var seguidor modelos.Usuario
		if erro := linhas.Scan(&seguidor.ID, &seguidor.Nome, &seguidor.Nick, &seguidor.Email, &seguidor.CriadoEm); erro != nil {
			return nil, erro
		}
		seguidores = append(seguidores, seguidor)
	}
	return seguidores, nil
}

// BuscarSeguindo - busca todos os usuários que um usuário está seguindo
func (repositorio usuarios) BuscarSeguidos(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
	select u.id, u.nome, u.nick, u.email, u.criadoEm 
	from usuarios u inner join seguidores s on u.id = s.seguido_id where s.seguidor_id = ?`, usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuariosSeguidos []modelos.Usuario

	for linhas.Next() {
		var usuarioSeguido modelos.Usuario
		if erro := linhas.Scan(&usuarioSeguido.ID, &usuarioSeguido.Nome, &usuarioSeguido.Nick, &usuarioSeguido.Email, &usuarioSeguido.CriadoEm); erro != nil {
			return nil, erro
		}
		usuariosSeguidos = append(usuariosSeguidos, usuarioSeguido)
	}
	return usuariosSeguidos, nil
}

// BuscarSenha - busca a senha de um usuário
func (repositorio usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	linhas, erro := repositorio.db.Query("select senha from usuarios where id = ?", usuarioID)
	if erro != nil {
		return "", erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario
	if linhas.Next() {
		if erro := linhas.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}

	return usuario.Senha, nil
}

// AtualizarSenha - atualiza a senha de um usuário
func (repositorio usuarios) AtualizarSenha(usuarioID uint64, senha string) error {
	statement, erro := repositorio.db.Prepare("update usuarios set senha = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(senha, usuarioID); erro != nil {
		return erro
	}
	return nil
}
