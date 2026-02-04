package goroutines

import (
	"fmt"
)

// DemonstrarWorkerPools demonstra o padrão Worker Pool
func DemonstrarWorkerPools() {
	fmt.Println("--- PADRÃO WORKER POOLS ---")
	fmt.Println("Worker Pool é um padrão onde um número fixo de goroutines (workers)")
	fmt.Println("processa trabalhos de uma fila. É útil para limitar concorrência e")
	fmt.Println("controlar o uso de recursos.\n")

	fmt.Println("Por que usar Worker Pool?")
	fmt.Println("  Sem limite: Criar uma goroutine para cada trabalho pode sobrecarregar")
	fmt.Println("  o sistema (memória, CPU, conexões de rede/banco). Exemplo: 10.000")
	fmt.Println("  requisições HTTP = 10.000 goroutines simultâneas = sistema travado.\n")
	fmt.Println("  Com Worker Pool: Limita concorrência a um número controlado (ex: 10")
	fmt.Println("  workers). Trabalhos esperam na fila e são processados conforme workers")
	fmt.Println("  ficam disponíveis. Isso controla recursos e mantém o sistema estável.\n")

	fmt.Println("Componentes do Worker Pool:")
	fmt.Println("  1. Fila de trabalhos (canal de entrada)")
	fmt.Println("  2. Pool de workers (goroutines que processam trabalhos)")
	fmt.Println("  3. Fila de resultados (canal de saída, opcional)\n")

	ExemploWorkerPoolAula()

	fmt.Println("Vantagens do Worker Pool:")
	fmt.Println("  ✓ Controla o número máximo de goroutines simultâneas")
	fmt.Println("  ✓ Evita sobrecarga do sistema")
	fmt.Println("  ✓ Permite processar grandes volumes de dados")
	fmt.Println("  ✓ Facilita controle de recursos (memória, CPU)")
	fmt.Println("  ✓ Permite balanceamento de carga entre workers\n")

	fmt.Println("Casos de uso:")
	fmt.Println("  - Processamento de requisições HTTP")
	fmt.Println("  - Processamento de arquivos em lote")
	fmt.Println("  - Scraping de websites")
	fmt.Println("  - Processamento de imagens")
	fmt.Println("  - Operações de banco de dados\n")

	fmt.Println("⚠️  IMPORTANTE:")
	fmt.Println("  - Escolha o número de workers baseado na carga e recursos disponíveis")
	fmt.Println("  - Sempre feche o canal de trabalhos quando não houver mais trabalhos")
	fmt.Println("  - Use WaitGroup para garantir que todos os workers terminem")
	fmt.Println("  - Considere usar context.Context para cancelamento")
	fmt.Println("  - Monitore o desempenho e ajuste o número de workers conforme necessário")
	fmt.Println()
}

func ExemploWorkerPoolAula() {
	fmt.Println("Exemplo 1: Worker Pool básico")
	fmt.Println("Usando worker pool para calcular Fibonacci de múltiplos números:")
	fmt.Println("Worker processa tarefas da fila e envia resultados:\n")

	tarefas := make(chan int, 10)
	resultados := make(chan int, 10)

	go worker(tarefas, resultados)

	fmt.Println("  Enviando tarefas para cálculo:")
	for i := 0; i < 10; i++ {
		tarefas <- i
		fmt.Printf("  Tarefa %d enviada\n", i)
	}
	close(tarefas)

	fmt.Println("\n  Resultados:")
	for i := 0; i < 10; i++ {
		resultado := <-resultados
		fmt.Printf("  Fibonacci de %d é: %d\n", i, resultado)
	}
	fmt.Println()
}

// worker processa tarefas do canal de entrada e envia resultados para o canal de saída
func worker(tarefas <-chan int, resultados chan<- int) {
	for tarefa := range tarefas {
		resultados <- calcularFibonacci(tarefa)
	}
}

// calcularFibonacci calcula o número de Fibonacci de forma recursiva
func calcularFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return calcularFibonacci(n-1) + calcularFibonacci(n-2)
}
