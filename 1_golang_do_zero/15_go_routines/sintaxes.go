package main

import (
	"fmt"
	"sync"
	"time"
)

// sintaxeGoroutineBasica demonstra como iniciar uma goroutine básica
func sintaxeGoroutineBasica() {
	go func() {
		fmt.Println("Executando em goroutine")
	}()
	time.Sleep(100 * time.Millisecond)
}

// sintaxeGoroutineComFuncao demonstra goroutine com função nomeada
func sintaxeGoroutineComFuncao() {
	go imprimirMensagem("Olá da goroutine")
	time.Sleep(100 * time.Millisecond)
}

func imprimirMensagem(mensagem string) {
	fmt.Println(mensagem)
}

// sintaxeWaitGroupBasico demonstra uso básico de WaitGroup
func sintaxeWaitGroupBasico() {
	var grupoEspera sync.WaitGroup

	grupoEspera.Add(2)

	go func() {
		defer grupoEspera.Done()
		fmt.Println("Goroutine 1 executando")
	}()

	go func() {
		defer grupoEspera.Done()
		fmt.Println("Goroutine 2 executando")
	}()

	grupoEspera.Wait()
	fmt.Println("Todas as goroutines terminaram")
}

// sintaxeCanalBasico demonstra criação e uso básico de canal
func sintaxeCanalBasico() {
	canal := make(chan string)

	go func() {
		canal <- "Mensagem enviada"
	}()

	mensagem := <-canal
	fmt.Println(mensagem)
}

// sintaxeCanalComBuffer demonstra canal com buffer
func sintaxeCanalComBuffer() {
	canalComBuffer := make(chan string, 2)

	canalComBuffer <- "Primeira mensagem"
	canalComBuffer <- "Segunda mensagem"

	fmt.Println(<-canalComBuffer)
	fmt.Println(<-canalComBuffer)
}

// sintaxeCanalFechar demonstra como fechar um canal
func sintaxeCanalFechar() {
	canal := make(chan int)

	go func() {
		defer close(canal)
		for i := 0; i < 3; i++ {
			canal <- i
		}
	}()

	for valor := range canal {
		fmt.Println(valor)
	}
}

// sintaxeSelectBasico demonstra uso básico do select
func sintaxeSelectBasico() {
	canalUm := make(chan string)
	canalDois := make(chan string)

	go func() {
		canalUm <- "Mensagem do canal um"
	}()

	go func() {
		canalDois <- "Mensagem do canal dois"
	}()

	select {
	case mensagemUm := <-canalUm:
		fmt.Println(mensagemUm)
	case mensagemDois := <-canalDois:
		fmt.Println(mensagemDois)
	}
}

// sintaxeSelectComDefault demonstra select com case default
func sintaxeSelectComDefault() {
	canal := make(chan string)

	select {
	case mensagem := <-canal:
		fmt.Println(mensagem)
	default:
		fmt.Println("Nenhuma mensagem disponível")
	}
}

// sintaxeSelectComTimeout demonstra select com timeout
func sintaxeSelectComTimeout() {
	canal := make(chan string)

	select {
	case mensagem := <-canal:
		fmt.Println(mensagem)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: nenhuma mensagem recebida")
	}
}

// sintaxeGeneratorBasico demonstra padrão generator básico
func sintaxeGeneratorBasico() {
	gerador := func() <-chan int {
		canalSaida := make(chan int)
		go func() {
			defer close(canalSaida)
			for i := 1; i <= 5; i++ {
				canalSaida <- i
			}
		}()
		return canalSaida
	}

	canalGerador := gerador()
	for valor := range canalGerador {
		fmt.Println(valor)
	}
}

// sintaxeGeneratorAula demonstra generator encapsulando goroutine (padrão da aula)
func sintaxeGeneratorAula() {
	canal := escreverSintaxe("Hello World")

	for i := 0; i < 5; i++ {
		fmt.Println(<-canal)
	}
}

// escreverSintaxe é função auxiliar que retorna generator (encapsula goroutine)
func escreverSintaxe(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	return canal
}

// sintaxeMultiplexadorAula demonstra multiplexador reutilizando generators
func sintaxeMultiplexadorAula() {
	canalMultiplexado := multiplexarSintaxe(
		escreverSintaxe("Hello World"),
		escreverSintaxe("Hello World 2"),
	)

	for i := 0; i < 10; i++ {
		fmt.Println(<-canalMultiplexado)
	}
}

// multiplexarSintaxe combina múltiplos canais em um único canal
func multiplexarSintaxe(canal1, canal2 <-chan string) <-chan string {
	canalSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canal1:
				canalSaida <- mensagem
			case mensagem := <-canal2:
				canalSaida <- mensagem
			}
		}
	}()

	return canalSaida
}

// sintaxeWorkerPoolAula demonstra worker pool básico
func sintaxeWorkerPoolAula() {
	tarefas := make(chan int, 10)
	resultados := make(chan int, 10)

	go workerSintaxe(tarefas, resultados)

	for i := 0; i < 10; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 10; i++ {
		resultado := <-resultados
		fmt.Printf("Fibonacci de %d é: %d\n", i, resultado)
	}
}

// workerSintaxe processa tarefas e envia resultados
func workerSintaxe(tarefas <-chan int, resultados chan<- int) {
	for tarefa := range tarefas {
		resultados <- calcularFibonacciSintaxe(tarefa)
	}
}

// calcularFibonacciSintaxe calcula Fibonacci recursivamente
func calcularFibonacciSintaxe(n int) int {
	if n <= 1 {
		return n
	}
	return calcularFibonacciSintaxe(n-1) + calcularFibonacciSintaxe(n-2)
}
