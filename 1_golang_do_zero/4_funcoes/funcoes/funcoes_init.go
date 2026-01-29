package funcoes

import "fmt"

var contadorInit int

// init() é executada automaticamente antes de main()
// Pode haver múltiplas funções init no mesmo pacote
// Ordem de execução: imports → init (em ordem alfabética dos arquivos) → main
func init() {
	contadorInit++
	fmt.Printf("init() executado (arquivo funcoes_init.go) - contador: %d\n", contadorInit)
}

func DemonstrarFuncaoInit() {
	fmt.Println("--- FUNÇÃO INIT ---")
	fmt.Println("A função init() é executada automaticamente antes de main().")
	fmt.Println("Pode haver múltiplas funções init no mesmo pacote.")
	fmt.Println("Ordem de execução: imports → init (em ordem alfabética dos arquivos) → main\n")

	fmt.Println("Exemplo 1: init() executado automaticamente")
	fmt.Printf("Contador de init: %d (incrementado pela função init acima)\n", contadorInit)
	fmt.Println()

	fmt.Println("Exemplo 2: init() pode inicializar variáveis")
	fmt.Printf("Variável inicializada por init: %d\n", variavelInicializada)
	fmt.Println()

	fmt.Println("Exemplo 3: init() pode executar código de inicialização")
	fmt.Println("A função init() acima já executou e incrementou o contador")
	fmt.Println()

	fmt.Println("⚠️  IMPORTANTE:")
	fmt.Println("  - init() não recebe parâmetros e não retorna valores")
	fmt.Println("  - init() é executada automaticamente, não precisa ser chamada")
	fmt.Println("  - Múltiplas funções init são executadas em ordem alfabética dos arquivos")
	fmt.Println("  - init() é útil para inicialização de variáveis, configurações, etc.")
	fmt.Println()
}

var variavelInicializada int

func init() {
	variavelInicializada = 100
	fmt.Println("Segunda função init() executada (inicializando variável)")
}
