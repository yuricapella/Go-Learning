package main

import (
	"fmt"
	"log"
	"os"

	app "github.com/yuricapella/Go-Learning/1_golang_do_zero/projeto_1_linha_de_comando/app_aula"
)

func main() {
	fmt.Println("Iniciando aplicação de linha de comando")
	aplicacao := app.Gerar()

	// Método Run da aplicação recebe os.Args (parâmetro padrão) que contém os argumentos passados na linha de comando
	// O método Run retorna erro e por isso precisa ser tratado
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Aplicação executada com sucesso")
}
