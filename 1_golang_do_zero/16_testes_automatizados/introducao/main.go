package main

import (
	"fmt"

	"github.com/yuricapella/Go-Learning/1_golang_do_zero/16_testes_automatizados/introducao/enderecos"
)

/*
	o comando go test ./... fala para o go entrar em todos os pacotes do projeto e executar todos os testes
	a primeira vez que teste ele roda tudo

ok      github.com/yuricapella/Go-Learning/1_golang_do_zero/17_testes_automatizados/enderecos   0.564s
apos rodar ele usa cache para guardar o teste e somenta verifica se o teste foi alterado, se nao, usa o cache
ok      github.com/yuricapella/Go-Learning/1_golang_do_zero/17_testes_automatizados/enderecos   (cached)

tambem da para rodar os testes com go test -v assim mostra o nome das funções de teste que estao sendo executadas

	go test -v

=== RUN   TestTipoDeEndereco
--- PASS: TestTipoDeEndereco (0.00s)
=== RUN   TestTipoDeEnderecoComStruct
--- PASS: TestTipoDeEnderecoComStruct (0.00s)
PASS
ok      github.com/yuricapella/Go-Learning/1_golang_do_zero/17_testes_automatizados/enderecos   1.061s

Da para colocar t.Parallel para os testes rodarem ao mesmo tempo
go test -v
=== RUN   TestTipoDeEndereco
=== PAUSE TestTipoDeEndereco
=== RUN   TestTipoDeEnderecoComStruct
=== PAUSE TestTipoDeEnderecoComStruct
=== CONT  TestTipoDeEndereco
--- PASS: TestTipoDeEndereco (0.00s)
=== CONT  TestTipoDeEnderecoComStruct
--- PASS: TestTipoDeEnderecoComStruct (0.00s)
PASS
ok      github.com/yuricapella/Go-Learning/1_golang_do_zero/17_testes_automatizados/enderecos   1.273s

tambem podemos rodar os testes com o comando:
go test --cover
PASS
coverage: 100.0% of statements
ok      github.com/yuricapella/Go-Learning/1_golang_do_zero/17_testes_automatizados/enderecos   1.279s

para saber exatamente oque faltou cobrir, podemos gerar um arquivo txt para ter o relatorio
go test --coverprofile cobertura.txt
porem ele gera de um jeito dificil de ler
mode: set
github.com/yuricapella/Go-Learning/1_golang_do_zero/17_testes_automatizados/enderecos/enderecos.go:11.45,24.36 5 1
github.com/yuricapella/Go-Learning/1_golang_do_zero/17_testes_automatizados/enderecos/enderecos.go:24.36,25.40 1 1
github.com/yuricapella/Go-Learning/1_golang_do_zero/17_testes_automatizados/enderecos/enderecos.go:25.40,27.4 1 1
github.com/yuricapella/Go-Learning/1_golang_do_zero/17_testes_automatizados/enderecos/enderecos.go:30.2,30.29 1 1
github.com/yuricapella/Go-Learning/1_golang_do_zero/17_testes_automatizados/enderecos/enderecos.go:30.29,33.3 1 1
github.com/yuricapella/Go-Learning/1_golang_do_zero/17_testes_automatizados/enderecos/enderecos.go:35.2,35.25 1 1

podemos utilizar o comando
go tool cover --func=cobertura.txt
go tool cover --func=cobertura.txt
github.com/yuricapella/Go-Learning/1_golang_do_zero/17_testes_automatizados/enderecos/enderecos.go:11:  TipoDeEndereco  90.0%
total:                                                                                                  (statements)    90.0%

porem ainda nao mostra oque queremos saber, onde esta a falta de cobertura e para isso usamos:
go tool cover --html=cobertura.txt
com isso ele abre um arquivo html que mostra o arquivo onde falta cobertura e a linha com as palavras em vermelho mostra oque falta
*/
func main() {
	fmt.Println(enderecos.TipoDeEndereco("Rua dos Bobos"))

}
