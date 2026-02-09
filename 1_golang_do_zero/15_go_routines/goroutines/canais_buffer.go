package goroutines

import (
	"fmt"
	"time"
)

// DemonstrarCanaisComBuffer demonstra o uso de canais com buffer
func DemonstrarCanaisComBuffer() {
	fmt.Println("--- CANAIS COM BUFFER ---")
	fmt.Println("Canais com buffer permitem armazenar valores antes de serem recebidos.")
	fmt.Println("Eles não bloqueiam o envio até que o buffer esteja cheio.\n")

	fmt.Println("Por que usar buffer?")
	fmt.Println("  Sem buffer: Produtor e consumidor devem estar sincronizados. Se o")
	fmt.Println("  produtor for rápido e o consumidor lento, o produtor fica bloqueado.\n")
	fmt.Println("  Com buffer: Desacopla produtor e consumidor. Produtor pode continuar")
	fmt.Println("  produzindo enquanto consumidor processa, melhorando throughput quando")
	fmt.Println("  há variação na velocidade de produção/consumo.\n")

	fmt.Println("Diferença entre canais sem buffer e com buffer:")
	fmt.Println("  - Sem buffer (unbuffered): Bloqueia até haver receptor")
	fmt.Println("  - Com buffer: Bloqueia apenas quando buffer está cheio\n")

	fmt.Println("Criação:")
	fmt.Println("  canalSemBuffer := make(chan int)        // Sem buffer")
	fmt.Println("  canalComBuffer := make(chan int, 3)     // Buffer de tamanho 3\n")

	ExemploCanalSemBuffer()
	ExemploCanalComBuffer()
	ExemploProdutorRapidoConsumidorLento()

	fmt.Println("⚠️  IMPORTANTE:")
	fmt.Println("  - Buffer não aumenta capacidade total, apenas desacopla envio/recebimento")
	fmt.Println("  - Buffer muito grande pode consumir muita memória")
	fmt.Println("  - Buffer muito pequeno pode causar bloqueios frequentes")
	fmt.Println("  - Use buffer quando souber o padrão de produção/consumo")
	fmt.Println("  - Canais sem buffer são mais seguros para sincronização")
	fmt.Println("  - Canais com buffer são melhores para throughput")
	fmt.Println()
}

func ExemploCanalSemBuffer() {
	fmt.Println("Exemplo 1: Canal sem buffer (bloqueia no envio)")
	fmt.Println("O envio bloqueia até que haja um receptor:\n")

	canalSemBuffer := make(chan string)

	go func() {
		fmt.Println("  Tentando enviar para canal sem buffer...")
		canalSemBuffer <- "Mensagem"
		fmt.Println("  Mensagem enviada (após receptor estar pronto)")
	}()

	time.Sleep(100 * time.Millisecond)
	fmt.Println("  Recebendo mensagem...")
	mensagem := <-canalSemBuffer
	fmt.Printf("  Mensagem recebida: %s\n\n", mensagem)
}

func ExemploCanalComBuffer() {
	fmt.Println("Exemplo 2: Canal com buffer (não bloqueia até buffer cheio)")
	fmt.Println("Você pode enviar múltiplos valores sem bloqueio:\n")

	canalComBuffer := make(chan string, 3)

	fmt.Println("  Enviando 3 mensagens para canal com buffer de tamanho 3:")
	canalComBuffer <- "Primeira mensagem"
	fmt.Println("  ✓ Primeira mensagem enviada")
	canalComBuffer <- "Segunda mensagem"
	fmt.Println("  ✓ Segunda mensagem enviada")
	canalComBuffer <- "Terceira mensagem"
	fmt.Println("  ✓ Terceira mensagem enviada")
	fmt.Println("  Buffer cheio, mas não bloqueou!\n")

	fmt.Println("  Recebendo mensagens:")
	for contador := 0; contador < 3; contador++ {
		mensagem := <-canalComBuffer
		fmt.Printf("  Recebido: %s\n", mensagem)
	}
	fmt.Println()
}

func ExemploProdutorRapidoConsumidorLento() {
	fmt.Println("Exemplo 3: Produtor rápido, consumidor lento")
	fmt.Println("Buffer ajuda quando o produtor é mais rápido:\n")

	canalProdutorConsumidor := make(chan int, 5)

	// Produtor rápido
	go func() {
		for numero := 1; numero <= 10; numero++ {
			canalProdutorConsumidor <- numero
			fmt.Printf("  Produtor enviou: %d\n", numero)
		}
		close(canalProdutorConsumidor)
	}()

	// Consumidor lento
	time.Sleep(200 * time.Millisecond)
	fmt.Println("  Consumidor começando a receber...")
	for valor := range canalProdutorConsumidor {
		fmt.Printf("  Consumidor recebeu: %d\n", valor)
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Println()
}
