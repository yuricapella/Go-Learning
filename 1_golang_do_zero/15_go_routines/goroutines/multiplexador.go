package goroutines

import (
	"fmt"
)

// DemonstrarMultiplexador demonstra o padrão Multiplexador (Fan-in)
func DemonstrarMultiplexador() {
	fmt.Println("--- PADRÃO MULTIPLEXADOR (FAN-IN) ---")
	fmt.Println("Multiplexador (ou Fan-in) é um padrão que combina múltiplos canais")
	fmt.Println("de entrada em um único canal de saída. É útil para consolidar")
	fmt.Println("resultados de múltiplas fontes.\n")

	fmt.Println("Por que usar Multiplexador?")
	fmt.Println("  Sem multiplexador: Se você tem 3 APIs diferentes retornando dados em")
	fmt.Println("  canais separados, precisa ler de cada um sequencialmente ou criar")
	fmt.Println("  lógica complexa para gerenciar múltiplos canais.\n")
	fmt.Println("  Com multiplexador: Combina múltiplos canais em um único canal de")
	fmt.Println("  saída usando select. Processa valores assim que chegam de qualquer")
	fmt.Println("  fonte, simplificando o código consumidor e melhorando eficiência.\n")

	fmt.Println("Características do Multiplexador:")
	fmt.Println("  - Recebe dados de múltiplos canais")
	fmt.Println("  - Combina em um único canal de saída")
	fmt.Println("  - Usa select para ler de múltiplos canais")
	fmt.Println("  - Fecha o canal de saída quando todos os canais de entrada fecham\n")

	ExemploMultiplexadorAula()

	fmt.Println("Vantagens do padrão Multiplexador:")
	fmt.Println("  ✓ Consolida resultados de múltiplas fontes")
	fmt.Println("  ✓ Permite processamento paralelo com resultado único")
	fmt.Println("  ✓ Facilita agregação de dados")
	fmt.Println("  ✓ Útil em padrões pipeline complexos")
	fmt.Println("  ✓ Permite combinar worker pools\n")

	fmt.Println("Casos de uso:")
	fmt.Println("  - Agregar resultados de múltiplas APIs")
	fmt.Println("  - Combinar resultados de worker pools")
	fmt.Println("  - Consolidar logs de múltiplas fontes")
	fmt.Println("  - Agregar métricas de diferentes serviços")
	fmt.Println("  - Combinar streams de dados\n")

	fmt.Println("⚠️  IMPORTANTE:")
	fmt.Println("  - Sempre verifique se o canal está fechado (ok == false)")
	fmt.Println("  - Desabilite cases de canais fechados (canal = nil)")
	fmt.Println("  - Feche o canal de saída quando todos os canais de entrada fecharem")
	fmt.Println("  - Use select para ler de múltiplos canais simultaneamente")
	fmt.Println("  - Considere usar context.Context para cancelamento")
	fmt.Println("  - O padrão Fan-in é complementar ao Fan-out")
	fmt.Println()
}

func ExemploMultiplexadorAula() {
	fmt.Println("Exemplo 1: Multiplexador reutilizando generators")
	fmt.Println("Reutiliza função escrever() de generator para criar múltiplos canais")
	fmt.Println("e combina em um único canal usando multiplexador:\n")

	canalMultiplexado := multiplexar(escrever("Hello World"), escrever("Hello World 2"))

	fmt.Println("  Valores multiplexados de múltiplos generators:")
	for i := 0; i < 10; i++ {
		fmt.Printf("  %s\n", <-canalMultiplexado)
	}
	fmt.Println()
}

// multiplexar combina múltiplos canais de entrada em um único canal de saída
// Usa select para ler do primeiro canal que tiver dados disponíveis
func multiplexar(canalDeEntrada1, canalDeEntrada2 <-chan string) <-chan string {
	canalSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canalDeEntrada1:
				canalSaida <- mensagem
			case mensagem := <-canalDeEntrada2:
				canalSaida <- mensagem
			}
		}
	}()

	return canalSaida
}
