package goroutines

import (
	"fmt"
	"time"
)

// DemonstrarGoroutinesBasicas demonstra o uso básico de goroutines
func DemonstrarGoroutinesBasicas() {
	fmt.Println("--- GOROUTINES BÁSICAS ---")
	fmt.Println("Goroutines são funções que executam de forma concorrente com outras funções.")
	fmt.Println("Elas são leves e gerenciadas pelo runtime do Go.\n")

	fmt.Println("Concorrência vs Paralelismo:")
	fmt.Println("  - Concorrência: Múltiplas tarefas sendo executadas em períodos alternados")
	fmt.Println("  - Paralelismo: Múltiplas tarefas sendo executadas simultaneamente")
	fmt.Println("  - Go permite concorrência através de goroutines")
	fmt.Println("  - O paralelismo depende do número de CPUs disponíveis\n")

	ExemploGoroutineBasica()
	ExemploMultiplasGoroutines()
	ExemploGoroutineComFuncaoAnonima()

	fmt.Println("⚠️  IMPORTANTE:")
	fmt.Println("  - Goroutines são leves: você pode criar milhares delas")
	fmt.Println("  - A função main() não espera goroutines automaticamente")
	fmt.Println("  - Se main() terminar, todas as goroutines são encerradas")
	fmt.Println("  - Use WaitGroup ou canais para sincronizar goroutines")
	fmt.Println("  - Goroutines compartilham o mesmo espaço de memória")
	fmt.Println()
}

func ExemploGoroutineBasica() {
	fmt.Println("Exemplo 1: Executando função em goroutine")
	fmt.Println("Para executar uma função em goroutine, use a palavra-chave 'go' antes da chamada:\n")

	fmt.Println("  go imprimirNumero(1)")
	go imprimirNumero(1)
	time.Sleep(100 * time.Millisecond)
	fmt.Println()
}

func ExemploMultiplasGoroutines() {
	fmt.Println("Exemplo 2: Múltiplas goroutines")
	fmt.Println("Você pode iniciar quantas goroutines quiser:\n")

	for numero := 1; numero <= 3; numero++ {
		fmt.Printf("  Iniciando goroutine %d\n", numero)
		go imprimirNumero(numero)
	}
	time.Sleep(200 * time.Millisecond)
	fmt.Println()
}

func ExemploGoroutineComFuncaoAnonima() {
	fmt.Println("Exemplo 3: Goroutine com função anônima")
	fmt.Println("Você pode usar funções anônimas diretamente:\n")

	go func(mensagem string) {
		fmt.Printf("  Mensagem da goroutine: %s\n", mensagem)
	}("Olá do mundo concorrente")
	time.Sleep(100 * time.Millisecond)
	fmt.Println()
}

// imprimirNumero é uma função auxiliar para demonstração
func imprimirNumero(numero int) {
	fmt.Printf("    Goroutine executando: número %d\n", numero)
}
