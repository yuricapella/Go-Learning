package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yuricapella/Go-Learning/1_golang_do_zero/projeto_2_devbook/src/config"
)

// Conectar abre a conexão com o banco de dados e a retorna
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
