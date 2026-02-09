package enderecos_test

import (
	"testing"

	. "github.com/yuricapella/Go-Learning/1_golang_do_zero/16_testes_automatizados/introducao/enderecos"
)

/* Teste unitário
arquivos de testes sao a unica exceção onde podemos ter dois pacotes diferentes na mesma pasta, colocamos pacote original + _test
enderecos = enderecos_test, além disso deve-se importar o pacote original pois agora nao temos mais as funções para testar
Colocando o alias . no pacote original ele entende como pacote principal e nao precisa mais colocar pacote.funcao, só funcao

a função sempre tem que começar com Test e o arquivo sempre com _test, teste unitario sempre fica no lugar onde o arquivo original esta, endereco.go e endereco_test.go
Precisa estar dentro do pacote que o teste esta para usar o comando go test e rodar os testes da pasta
*/

func TestTipoDeEndereco(test *testing.T) {
	enderecoParaTeste := "Rua dos Bobos"
	tipoDeEnderecoEsperado := "Rua"

	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		test.Errorf("Tipo de endereço recebido: %s, Tipo de endereço esperado: %s", tipoDeEnderecoRecebido, tipoDeEnderecoEsperado)
	}
}

/* se passar ele coloca PASS e ok no terminal
 se falhar ele coloca FAIL, o nome do teste, o arquivo e a linha que deu erro, tambem informa o erro que foi parametrizado dessa linha
--- FAIL: TestTipoDeEndereco (0.00s)
    enderecos_test.go:19: Tipo de endereço recebido: Rua, Tipo de endereço esperado: rua
FAIL
exit status 1
FAIL    github.com/yuricapella/Go-Learning/1_golang_do_zero/17_testes_automatizados/enderecos   0.638s
*/

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEnderecoComStruct(test *testing.T) {
	cenariosDeTeste := []cenarioDeTeste{
		{enderecoInserido: "Rua dos Bobos", retornoEsperado: "Rua"},
		{enderecoInserido: "rua dos bobos", retornoEsperado: "Rua"},
		{enderecoInserido: "AVENIDA DOS BOBOS", retornoEsperado: "Avenida"},
		{enderecoInserido: "Avenida dos Bobos", retornoEsperado: "Avenida"},
		{enderecoInserido: "Estrada dos Bobos", retornoEsperado: "Estrada"},
		{enderecoInserido: "Rodovia dos Bobos", retornoEsperado: "Rodovia"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			test.Errorf("Tipo de endereço recebido: %s, Tipo de endereço esperado: %s", tipoDeEnderecoRecebido, cenario.retornoEsperado)
		}
	}
}
