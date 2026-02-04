package goroutines

import (
	"fmt"
	"time"
)

// DemonstrarCanais demonstra o uso de canais para comunicação entre goroutines
func DemonstrarCanais() {
	fmt.Println("--- CANAIS ---")
	fmt.Println("Canais são o mecanismo primário de comunicação entre goroutines em Go.")
	fmt.Println("Eles permitem enviar e receber valores de forma segura e sincronizada.\n")

	fmt.Println("Criação de canais:")
	fmt.Println("  canal := make(chan Tipo)        // Canal sem buffer (unbuffered)")
	fmt.Println("  canal := make(chan Tipo, tamanho) // Canal com buffer\n")

	fmt.Println("Por que usar canais?")
	fmt.Println("  Sem canais: Para comunicar entre goroutines, você precisaria usar")
	fmt.Println("  variáveis compartilhadas com mutexes, o que é propenso a erros")
	fmt.Println("  (race conditions, deadlocks).\n")
	fmt.Println("  Com canais: Comunicação segura e sincronizada. Segue a filosofia")
	fmt.Println("  Go: 'Don't communicate by sharing memory; share memory by communicating'.")
	fmt.Println("  Canais garantem que apenas uma goroutine acesse o dado por vez.\n")

	fmt.Println("Operações básicas:")
	fmt.Println("  canal <- valor  // Enviar valor para o canal")
	fmt.Println("  valor := <-canal  // Receber valor do canal")
	fmt.Println("  close(canal)  // Fechar o canal\n")

	ExemploCanalBasico()
	ExemploCanalMultiplosValores()
	ExemploCanalBidirecional()

	fmt.Println("--- COMPARAÇÃO: CANAIS vs WAITGROUP ---\n")

	fmt.Println("WaitGroup:")
	fmt.Println("  ✓ Ideal quando você só precisa saber quando goroutines terminam")
	fmt.Println("  ✓ Não permite comunicação de dados")
	fmt.Println("  ✓ Simples para sincronização básica")
	fmt.Println("  ✗ Não permite troca de informações entre goroutines\n")

	fmt.Println("Canais:")
	fmt.Println("  ✓ Permitem comunicação de dados entre goroutines")
	fmt.Println("  ✓ Sincronização implícita (bloqueio automático)")
	fmt.Println("  ✓ Composable (podem ser combinados em padrões complexos)")
	fmt.Println("  ✓ Suportam múltiplos leitores/escritores")
	fmt.Println("  ✓ Podem ser fechados para sinalizar fim")
	fmt.Println("  ✓ Funcionam bem com select para multiplexação\n")

	fmt.Println("Por que canais são mais usados que WaitGroup?")
	fmt.Println("  1. Comunicação: Canais permitem trocar dados, não apenas sincronizar")
	fmt.Println("  2. Composição: Canais podem ser combinados em padrões complexos")
	fmt.Println("  3. Filosofia Go: 'Don't communicate by sharing memory; share memory by communicating'")
	fmt.Println("  4. Flexibilidade: Canais podem fazer o papel de WaitGroup e muito mais")
	fmt.Println("  5. Padrões: Worker pools, generators, pipelines são mais fáceis com canais\n")

	ExemploCanalComoWaitGroup()

	fmt.Println("⚠️  IMPORTANTE:")
	fmt.Println("  - Canais sem buffer bloqueiam até que haja um receptor")
	fmt.Println("  - Canais com buffer bloqueiam apenas quando o buffer está cheio")
	fmt.Println("  - Sempre feche canais quando não precisar mais deles")
	fmt.Println("  - Receber de um canal fechado retorna zero value e ok=false")
	fmt.Println("  - Enviar para um canal fechado causa panic")
	fmt.Println("  - Use range para receber valores até o canal ser fechado")
	fmt.Println("  - Canais são a forma idiomática de comunicação em Go")
	fmt.Println()
}

func ExemploCanalBasico() {
	fmt.Println("Exemplo 1: Canal básico para comunicação")
	fmt.Println("Criando um canal e enviando/recebendo valores:\n")

	canalMensagens := make(chan string)

	go func() {
		canalMensagens <- "Mensagem da goroutine"
		fmt.Println("  Mensagem enviada")
	}()

	mensagemRecebida := <-canalMensagens
	fmt.Printf("  Mensagem recebida: %s\n\n", mensagemRecebida)
}

func ExemploCanalMultiplosValores() {
	fmt.Println("Exemplo 2: Canal com múltiplos valores")
	fmt.Println("Enviando e recebendo múltiplos valores:\n")

	canalNumeros := make(chan int)

	go func() {
		defer close(canalNumeros)
		for numero := 1; numero <= 3; numero++ {
			canalNumeros <- numero
			fmt.Printf("  Enviado: %d\n", numero)
		}
	}()

	fmt.Println("  Recebendo valores:")
	for numero := range canalNumeros {
		fmt.Printf("  Recebido: %d\n", numero)
	}
	fmt.Println()
}

func ExemploCanalBidirecional() {
	fmt.Println("Exemplo 3: Canal bidirecional")
	fmt.Println("Canais podem ser usados para comunicação em ambas as direções:\n")

	canalBidirecional := make(chan string)

	go func() {
		canalBidirecional <- "Olá"
		resposta := <-canalBidirecional
		fmt.Printf("  Goroutine recebeu: %s\n", resposta)
	}()

	mensagem := <-canalBidirecional
	fmt.Printf("  Main recebeu: %s\n", mensagem)
	canalBidirecional <- "Tchau"
	time.Sleep(100 * time.Millisecond)
	fmt.Println()
}

func ExemploCanalComoWaitGroup() {
	fmt.Println("Exemplo 4: Usando canal como WaitGroup")
	fmt.Println("Você pode usar um canal para sincronizar goroutines:\n")

	canalSincronizacao := make(chan bool)
	numeroGoroutines := 3

	for identificador := 1; identificador <= numeroGoroutines; identificador++ {
		go func(id int) {
			fmt.Printf("  Goroutine %d executando\n", id)
			time.Sleep(100 * time.Millisecond)
			canalSincronizacao <- true
		}(identificador)
	}

	for contador := 0; contador < numeroGoroutines; contador++ {
		<-canalSincronizacao
	}
	fmt.Println("  Todas as goroutines terminaram (usando canal como WaitGroup)\n")
}
