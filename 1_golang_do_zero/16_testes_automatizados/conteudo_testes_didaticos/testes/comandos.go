package testes

import (
	"fmt"
)

// DemonstrarComandosTest demonstra os comandos e flags do go test
func DemonstrarComandosTest() {
	fmt.Println("--- COMANDOS E FLAGS DO GO TEST ---")
	fmt.Println("O comando 'go test' executa testes em pacotes Go.")
	fmt.Println("Ele possui várias flags úteis para controlar a execução.\n")

	fmt.Println("Comando básico:")
	fmt.Println("  go test")
	fmt.Println("  - Executa todos os testes no pacote atual")
	fmt.Println("  - Mostra apenas se passou (PASS) ou falhou (FAIL)")
	fmt.Println("  - Usa cache quando possível\n")

	fmt.Println("Comando 1: Modo verbose (-v)")
	fmt.Println("  go test -v")
	fmt.Println("  - Mostra o nome de cada teste executado")
	fmt.Println("  - Exibe mensagens de log durante a execução")
	fmt.Println("  - Útil para depuração e ver progresso\n")

	fmt.Println("  Saída exemplo:")
	fmt.Println("  === RUN   TestSoma")
	fmt.Println("  --- PASS: TestSoma (0.00s)")
	fmt.Println("  === RUN   TestSubtracao")
	fmt.Println("  --- PASS: TestSubtracao (0.00s)")
	fmt.Println("  PASS\n")

	fmt.Println("Comando 2: Executar em todos os pacotes (./...)")
	fmt.Println("  go test ./...")
	fmt.Println("  - Executa testes em todos os pacotes do projeto")
	fmt.Println("  - Entra recursivamente em todas as subpastas")
	fmt.Println("  - Útil para executar toda a suíte de testes\n")

	fmt.Println("  Saída exemplo:")
	fmt.Println("  ok      github.com/projeto/pacote1   0.564s")
	fmt.Println("  ok      github.com/projeto/pacote2   0.234s")
	fmt.Println("  ok      github.com/projeto/pacote3   0.123s\n")

	fmt.Println("Comando 3: Executar teste específico (-run)")
	fmt.Println("  go test -run TestNomeDoTeste")
	fmt.Println("  - Executa apenas testes que correspondem ao padrão")
	fmt.Println("  - Suporta expressões regulares")
	fmt.Println("  - Útil para testar função específica\n")

	fmt.Println("  Exemplos:")
	fmt.Println("  go test -run TestSoma              // Apenas TestSoma")
	fmt.Println("  go test -run TestSoma/Subteste     // Subteste específico")
	fmt.Println("  go test -run '^TestSoma$'          // Regex exata")
	fmt.Println("  go test -run 'Test.*'              // Todos que começam com Test\n")

	fmt.Println("Comando 4: Verificar cobertura (--cover)")
	fmt.Println("  go test --cover")
	fmt.Println("  - Mostra porcentagem de cobertura de código")
	fmt.Println("  - Indica quantas linhas foram executadas pelos testes")
	fmt.Println("  - Útil para identificar código não testado\n")

	fmt.Println("  Saída exemplo:")
	fmt.Println("  PASS")
	fmt.Println("  coverage: 85.7%% of statements")
	fmt.Println("  ok      github.com/projeto/pacote   1.279s\n")

	fmt.Println("Comando 5: Gerar perfil de cobertura (--coverprofile)")
	fmt.Println("  go test --coverprofile=cobertura.txt")
	fmt.Println("  - Gera arquivo com detalhes da cobertura")
	fmt.Println("  - Pode ser analisado com go tool cover")
	fmt.Println("  - Útil para relatórios detalhados\n")

	fmt.Println("Comando 6: Executar testes em paralelo (-parallel)")
	fmt.Println("  go test -parallel 4")
	fmt.Println("  - Define número máximo de testes paralelos")
	fmt.Println("  - Padrão é GOMAXPROCS")
	fmt.Println("  - Útil para acelerar execução\n")

	fmt.Println("Comando 7: Mostrar informações de benchmark (-bench)")
	fmt.Println("  go test -bench=.")
	fmt.Println("  - Executa benchmarks (funções Benchmark*)")
	fmt.Println("  - Mostra tempo de execução")
	fmt.Println("  - Útil para medir performance\n")

	fmt.Println("Comando 8: Pular cache (-count=1)")
	fmt.Println("  go test -count=1")
	fmt.Println("  - Força execução sem usar cache")
	fmt.Println("  - Útil quando quer garantir execução completa")
	fmt.Println("  - Útil para depuração\n")

	fmt.Println("Comando 9: Timeout para testes (-timeout)")
	fmt.Println("  go test -timeout 30s")
	fmt.Println("  - Define tempo máximo para execução")
	fmt.Println("  - Padrão é 10 minutos")
	fmt.Println("  - Útil para detectar testes travados\n")

	fmt.Println("Comando 10: Mostrar output de testes (-args)")
	fmt.Println("  go test -v -args -debug")
	fmt.Println("  - Passa argumentos para os testes")
	fmt.Println("  - Útil para flags customizadas\n")

	fmt.Println("Cache de testes:")
	fmt.Println("  - Go armazena resultados de testes em cache")
	fmt.Println("  - Testes não modificados são marcados como (cached)")
	fmt.Println("  - Cache é invalidado quando código muda")
	fmt.Println("  - Acelera execução de testes repetidos\n")

	fmt.Println("  Saída com cache:")
	fmt.Println("  ok      github.com/projeto/pacote   (cached)")
	fmt.Println("  ok      github.com/projeto/pacote2   0.234s\n")

	fmt.Println("Flags úteis combinadas:")
	fmt.Println("  go test -v -cover ./...")
	fmt.Println("  - Verbose + cobertura em todos os pacotes\n")

	fmt.Println("  go test -v -run TestSoma -cover")
	fmt.Println("  - Verbose + teste específico + cobertura\n")

	fmt.Println("  go test -v -count=1 -timeout 30s")
	fmt.Println("  - Verbose + sem cache + timeout\n")

	fmt.Println("⚠️  IMPORTANTE:")
	fmt.Println("  - Use -v para depuração e ver progresso")
	fmt.Println("  - Use ./... para executar todos os testes do projeto")
	fmt.Println("  - Use -run para executar testes específicos")
	fmt.Println("  - Use --cover para verificar cobertura básica")
	fmt.Println("  - Use --coverprofile para análise detalhada")
	fmt.Println("  - Cache acelera execução mas pode ocultar problemas")
	fmt.Println("  - Use -count=1 para garantir execução completa")
	fmt.Println("  - Combine flags conforme necessário")
	fmt.Println()
}
