package funcoes

import (
	"fmt"
)

func DemonstrarFuncoesBasicas() {
	fmt.Println("--- FUNÇÕES BÁSICAS ---")
	fmt.Println("Funções são blocos de código reutilizáveis que podem receber parâmetros e retornar valores.\n")

	fmt.Println("Exemplo 1: Função simples")
	resultado := soma(10, 20)
	fmt.Printf("soma(10, 20) = %d\n", resultado)
	fmt.Println()

	fmt.Println("Exemplo 2: Função com múltiplos retornos")
	soma, subtracao, multiplicacao, divisao := calculosMatematicos(10, 20)
	fmt.Printf("calculosMatematicos(10, 20):\n")
	fmt.Printf("  Soma: %d\n", soma)
	fmt.Printf("  Subtração: %d\n", subtracao)
	fmt.Printf("  Multiplicação: %d\n", multiplicacao)
	fmt.Printf("  Divisão: %d\n", divisao)
	fmt.Println()

	fmt.Println("Exemplo 3: Ignorando valores de retorno com _")
	soma2, _, _, _ := calculosMatematicos(10, 20)
	fmt.Printf("Apenas soma (ignorando outros retornos): %d\n", soma2)
	fmt.Println()

	fmt.Println("Exemplo 4: Função com retorno de erro")
	nome, idade, erro := mostrarNomeEIdade("João", 20)
	if erro != nil {
		fmt.Printf("Erro: %v\n", erro)
	} else {
		fmt.Printf("Nome: %s, Idade: %d\n", nome, idade)
	}
	fmt.Println()
}

func soma(a int, b int) int {
	return a + b
}

func calculosMatematicos(n1, n2 int) (int, int, int, int) {
	soma := n1 + n2
	subtracao := n1 - n2
	multiplicacao := n1 * n2
	divisao := n1 / n2
	return soma, subtracao, multiplicacao, divisao
}

func mostrarNomeEIdade(nome string, idade int) (string, int, error) {
	return nome, idade, nil
}
