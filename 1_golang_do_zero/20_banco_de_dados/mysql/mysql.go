package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Conectar - Conecta ao banco de dados
func Conectar() (*sql.DB, error) {
	stringConexao := "yuri:root@/dbteste?charset=utf8&parseTime=True&loc=Local"

	dataBase, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = dataBase.Ping(); erro != nil {
		dataBase.Close()
		return nil, erro
	}

	fmt.Println("Conexão estabelecida com sucesso!")

	return dataBase, nil
}

func PrepararBancoDeDados() error {
	dataBase, erro := Conectar()
	if erro != nil {
		return erro
	}
	defer func() {
		dataBase.Close()
		fmt.Println("Conexão fechada")
	}()

	query := `
      CREATE TABLE IF NOT EXISTS usuarios (
        id INT AUTO_INCREMENT PRIMARY KEY,
        nome VARCHAR(100) NOT NULL,
        email VARCHAR(100) NOT NULL
      );
    `

	_, erro = dataBase.Exec(query)
	if erro != nil {
		return fmt.Errorf("erro ao preparar banco de dados: %v", erro)
	}
	return nil
}
