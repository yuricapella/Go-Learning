package goroutines

import (
	"fmt"
	"sync"
	"time"
)

// DemonstrarWaitGroup demonstra o uso de sync.WaitGroup para sincronização
func DemonstrarWaitGroup() {
	fmt.Println("--- WAITGROUP ---")
	fmt.Println("WaitGroup é uma estrutura do pacote sync que permite esperar")
	fmt.Println("até que um grupo de goroutines termine sua execução.\n")

	fmt.Println("Por que usar WaitGroup?")
	fmt.Println("  Sem WaitGroup: Se main() terminar antes das goroutines, elas são")
	fmt.Println("  encerradas abruptamente. Você não sabe quando todas terminaram.\n")
	fmt.Println("  Com WaitGroup: Garante que main() espere todas as goroutines")
	fmt.Println("  terminarem antes de continuar. Ideal quando você só precisa")
	fmt.Println("  sincronizar fim de execução, sem trocar dados entre goroutines.\n")

	fmt.Println("Métodos principais:")
	fmt.Println("  - Add(quantidade): Adiciona goroutines ao grupo")
	fmt.Println("  - Done(): Indica que uma goroutine terminou")
	fmt.Println("  - Wait(): Bloqueia até que todas as goroutines terminem\n")

	ExemploWaitGroupComDefer()
	ExemploWaitGroupEmLoop()

	fmt.Println("⚠️  IMPORTANTE:")
	fmt.Println("  - Sempre chame Done() quando a goroutine terminar")
	fmt.Println("  - Use defer para garantir que Done() seja chamado mesmo em caso de erro")
	fmt.Println("  - Add() deve ser chamado antes de iniciar a goroutine")
	fmt.Println("  - Wait() bloqueia até que todas as goroutines terminem")
	fmt.Println("  - WaitGroup é útil quando você só precisa saber quando goroutines terminam")
	fmt.Println("  - Para comunicação entre goroutines, use canais")
	fmt.Println()
}

func ExemploWaitGroupComDefer() {
	fmt.Println("Exemplo 1: Uso básico de WaitGroup")
	fmt.Println("Construção básica usando defer para garantir que Done() seja chamado:\n")

	var grupoEsperaComDefer sync.WaitGroup
	grupoEsperaComDefer.Add(2)

	go func() {
		defer grupoEsperaComDefer.Done()
		fmt.Println("  Tarefa 1: Processando dados...")
		time.Sleep(150 * time.Millisecond)
		fmt.Println("  Tarefa 1: Concluída")
	}()

	go func() {
		defer grupoEsperaComDefer.Done()
		fmt.Println("  Tarefa 2: Processando dados...")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("  Tarefa 2: Concluída")
	}()

	grupoEsperaComDefer.Wait()
	fmt.Println("  Todas as tarefas concluídas!\n")
}

func ExemploWaitGroupEmLoop() {
	fmt.Println("Exemplo 2: WaitGroup em loop")
	fmt.Println("Adicionando goroutines dinamicamente:\n")

	var grupoEsperaLoop sync.WaitGroup
	tarefas := []string{"Tarefa A", "Tarefa B", "Tarefa C"}

	for _, tarefa := range tarefas {
		grupoEsperaLoop.Add(1)
		go func(nomeTarefa string) {
			defer grupoEsperaLoop.Done()
			fmt.Printf("  Executando: %s\n", nomeTarefa)
			time.Sleep(50 * time.Millisecond)
		}(tarefa)
	}

	grupoEsperaLoop.Wait()
	fmt.Println("  Todas as tarefas do loop concluídas!\n")
}
