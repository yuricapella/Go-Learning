package goroutines

import (
	"fmt"
	"time"
)

// DemonstrarSelect demonstra o uso do select statement para multiplexação de canais
func DemonstrarSelect() {
	fmt.Println("--- SELECT STATEMENT ---")
	fmt.Println("Select permite aguardar múltiplas operações de canal simultaneamente.")
	fmt.Println("É similar ao switch, mas para operações de canal.\n")

	fmt.Println("Por que usar select?")
	fmt.Println("  Sem select: Se você tentar receber de múltiplos canais sequencialmente,")
	fmt.Println("  uma goroutine lenta pode bloquear outras rápidas. Exemplo:")
	fmt.Println("    valor1 := <-canal1  // Bloqueia aqui se canal1 estiver vazio")
	fmt.Println("    valor2 := <-canal2  // Nunca chega aqui se canal1 demorar\n")
	fmt.Println("  Com select: Processa o primeiro canal que estiver pronto, evitando")
	fmt.Println("  que goroutines rápidas sejam bloqueadas por lentas. Isso melhora")
	fmt.Println("  a responsividade e eficiência do sistema.\n")

	fmt.Println("Sintaxe básica:")
	fmt.Println("  select {")
	fmt.Println("  case valor := <-canal1:")
	fmt.Println("      // Processar valor do canal1")
	fmt.Println("  case valor := <-canal2:")
	fmt.Println("      // Processar valor do canal2")
	fmt.Println("  default:")
	fmt.Println("      // Executar se nenhum canal estiver pronto")
	fmt.Println("  }\n")

	ExemploSelectMultiplosCanais()
	ExemploSelectComDefault()
	ExemploSelectComTimeout()
}

func ExemploSelectMultiplosCanais() {
	fmt.Println("Exemplo 1: Select com múltiplos canais")
	fmt.Println("Evita bloqueio: se canal1 demorar 150ms e canal2 100ms, select processa")
	fmt.Println("canal2 primeiro (mais rápido) em vez de esperar canal1 bloquear tudo:\n")

	canalUm := make(chan string)
	canalDois := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		canalUm <- "Mensagem do canal um"
	}()

	go func() {
		time.Sleep(150 * time.Millisecond)
		canalDois <- "Mensagem do canal dois"
	}()

	select {
	case mensagemUm := <-canalUm:
		fmt.Printf("  Recebido do canal um: %s\n", mensagemUm)
	case mensagemDois := <-canalDois:
		fmt.Printf("  Recebido do canal dois: %s\n", mensagemDois)
	}
	fmt.Println()
}

func ExemploSelectComDefault() {
	fmt.Println("Exemplo 2: Select com case default (não bloqueante)")
	fmt.Println("Útil para operações não-bloqueantes: permite fazer outras tarefas quando")
	fmt.Println("nenhum canal está pronto, evitando que o código trave esperando:\n")

	canalVazio := make(chan string)

	select {
	case mensagem := <-canalVazio:
		fmt.Printf("  Recebido: %s\n", mensagem)
	default:
		fmt.Println("  Nenhum canal está pronto, executando default")
	}
	fmt.Println()
}

func ExemploSelectComTimeout() {
	fmt.Println("Exemplo 3: Select com timeout")
	fmt.Println("Evita espera infinita: se uma operação demorar muito, timeout cancela")
	fmt.Println("e permite continuar. Essencial para sistemas que precisam de resposta rápida:\n")

	canalComTimeout := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		canalComTimeout <- "Mensagem recebida"
	}()

	select {
	case mensagem := <-canalComTimeout:
		fmt.Printf("  %s\n", mensagem)
	case <-time.After(1 * time.Second):
		fmt.Println("  Timeout: nenhuma mensagem recebida em 1 segundo")
	}
	fmt.Println()
}
