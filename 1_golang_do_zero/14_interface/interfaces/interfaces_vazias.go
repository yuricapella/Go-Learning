package interfaces

import "fmt"

func DemonstrarInterfacesVazias() {
	fmt.Println("--- INTERFACES VAZIAS ---")
	fmt.Println("A interface vazia `interface{}` (ou `any` no Go 1.18+) funciona como um tipo genérico.")
	fmt.Println("Um tipo genérico é capaz de aceitar qualquer tipo, permitindo flexibilidade no código.")
	fmt.Println("É útil quando precisamos trabalhar com tipos desconhecidos ou heterogêneos.\n")

	fmt.Println("Exemplo 1: interface{} como tipo genérico aceita qualquer tipo")
	fmt.Println("Como tipo genérico, interface{} pode armazenar valores de qualquer tipo:\n")

	var valorQualquer interface{}
	valorQualquer = "texto"
	fmt.Printf("valorQualquer (tipo genérico interface{}) como string: %s\n", valorQualquer)

	valorQualquer = 42
	fmt.Printf("valorQualquer (tipo genérico interface{}) como int: %d\n", valorQualquer)

	valorQualquer = true
	fmt.Printf("valorQualquer (tipo genérico interface{}) como bool: %t\n", valorQualquer)
	fmt.Println()

	fmt.Println("Exemplo 2: Slice heterogêneo usando tipo genérico interface{}")
	fmt.Println("Como interface{} é um tipo genérico, podemos criar slices que contêm diferentes tipos:\n")

	listaValores := []interface{}{
		"texto",
		42,
		true,
		3.14,
		[]string{"a", "b", "c"},
	}

	fmt.Println("Lista heterogênea (usando tipo genérico interface{}):")
	for indice, valor := range listaValores {
		fmt.Printf("  [%d] %v (tipo: %T)\n", indice, valor, valor)
	}
	fmt.Println()

	fmt.Println("Exemplo 3: Função genérica usando interface{}")
	fmt.Println("Funções podem receber interface{} (tipo genérico) para aceitar qualquer tipo:\n")

	exibirTipo("texto")
	exibirTipo(42)
	exibirTipo(true)
	exibirTipo(3.14)
	fmt.Println()

	fmt.Println("⚠️  IMPORTANTE:")
	fmt.Println("  - interface{} ou any funciona como tipo genérico, aceitando qualquer tipo")
	fmt.Println("  - Tipo genérico permite flexibilidade ao trabalhar com tipos desconhecidos")
	fmt.Println("  - Útil para criar funções genéricas antes do Go 1.18")
	fmt.Println("  - Requer type assertion para recuperar o tipo original")
	fmt.Println()
}

func exibirTipo(valorQualquer interface{}) {
	fmt.Printf("Valor: %v, Tipo: %T\n", valorQualquer, valorQualquer)
}
