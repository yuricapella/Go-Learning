package goroutines

import (
	"fmt"
	"math/rand"
	"time"
)

// DemonstrarGenerator demonstra o padrão Generator
func DemonstrarGenerator() {
	fmt.Println("--- PADRÃO GENERATOR ---")
	fmt.Println("Generator é um padrão onde uma função retorna um canal que produz")
	fmt.Println("uma sequência de valores. É útil para gerar dados de forma lazy\n")

	fmt.Println("Por que usar Generator?")
	fmt.Println("  Sem generator: Se você precisa gerar milhões de números, precisa criar")
	fmt.Println("  todos na memória antes de processar. Isso consome muita memória.\n")
	fmt.Println("  Com generator: Produz valores sob demanda (lazy). Só gera o próximo")
	fmt.Println("  valor quando o consumidor pedir. Isso economiza memória e permite")
	fmt.Println("  trabalhar com sequências infinitas ou muito grandes.\n")

	fmt.Println("Características do Generator:")
	fmt.Println("  - Retorna um canal de leitura (<-chan)")
	fmt.Println("  - Produz valores em uma goroutine separada")
	fmt.Println("  - Fecha o canal quando não há mais valores")
	fmt.Println("  - Permite consumo lazy de dados\n")

	ExemploGeneratorBasico()
	ExemploGeneratorAula()

	fmt.Println("Vantagens do padrão Generator:")
	fmt.Println("  ✓ Produção lazy de dados (só gera quando necessário)")
	fmt.Println("  ✓ Separação de responsabilidades (produção vs consumo)")
	fmt.Println("  ✓ Permite composição com outros padrões")
	fmt.Println("  ✓ Facilita testes e mock de dados")
	fmt.Println("  ✓ Pode ser usado com pipelines\n")

	fmt.Println("Casos de uso:")
	fmt.Println("  - Geração de sequências infinitas ou grandes")
	fmt.Println("  - Leitura de arquivos grandes linha por linha")
	fmt.Println("  - Processamento de streams de dados")
	fmt.Println("  - Geração de IDs ou tokens")
	fmt.Println("  - Iteração sobre estruturas de dados grandes\n")

	fmt.Println("⚠️  IMPORTANTE:")
	fmt.Println("  - Sempre feche o canal quando não houver mais valores")
	fmt.Println("  - Use defer close() para garantir fechamento")
	fmt.Println("  - Canais retornados devem ser de leitura (<-chan)")
	fmt.Println("  - Generators podem ser compostos em pipelines")
	fmt.Println("  - Considere usar context.Context para cancelamento")
	fmt.Println()
}

func ExemploGeneratorBasico() {
	fmt.Println("Exemplo 1: Generator básico de números")
	fmt.Println("Gerando uma sequência de números:\n")

	geradorNumeros := func(limite int) <-chan int {
		canalSaida := make(chan int)
		go func() {
			defer close(canalSaida)
			for numero := 1; numero <= limite; numero++ {
				canalSaida <- numero
			}
		}()
		return canalSaida
	}

	fmt.Println("  Gerando números de 1 a 5:")
	canalGerador := geradorNumeros(5)
	for numero := range canalGerador {
		fmt.Printf("  Recebido: %d\n", numero)
	}
	fmt.Println()
}

func ExemploGeneratorAula() {
	fmt.Println("Exemplo 2: Generator encapsulando goroutine")
	fmt.Println("Generator esconde complexidade: encapsula a criação da goroutine e")
	fmt.Println("retorna apenas o canal. O consumidor não precisa saber sobre goroutines:\n")

	canal := escrever("Hello World")

	fmt.Println("  Consumindo valores do generator:")
	for i := 0; i < 10; i++ {
		fmt.Printf("  %s\n", <-canal)
	}
	fmt.Println()
}

// escrever é uma função auxiliar que retorna um generator
// Encapsula a criação da goroutine, retornando apenas o canal de leitura
func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
		}
	}()

	return canal
}
