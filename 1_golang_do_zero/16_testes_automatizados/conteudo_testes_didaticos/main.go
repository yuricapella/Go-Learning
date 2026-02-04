package main

import (
	"fmt"

	"github.com/yuricapella/Go-Learning/1_golang_do_zero/16_testes_automatizados/conteudo_testes_didaticos/testes"
)

func main() {
	fmt.Println("=== TESTES AUTOMATIZADOS EM GO ===\n")

	testes.DemonstrarTestesBasicos()
	testes.DemonstrarTestesTabela()
	testes.DemonstrarSubtestes()
	testes.DemonstrarComandosTest()
	testes.DemonstrarCobertura()
}
