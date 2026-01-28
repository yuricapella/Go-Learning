package main

import (
	"fmt"
	"reflect"
)

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

func main() {
	resultado := soma(10, 20)
	fmt.Println("Resultado da soma:", resultado)

	var f = func() {
		fmt.Println("Função f")
	}

	f()
	fmt.Println("Tipo de f:", reflect.TypeOf(f))

	// o tipo func ele mostra tanto os parametros quanto o retorno como seu tipo
	var f2 = func(txt string) string {
		return txt
	}

	resultado2 := f2("Função f2")
	fmt.Println("Resultado da função f2:", resultado2)
	fmt.Println("Tipo de f2:", reflect.TypeOf(f2))

	// podemos ter mais de um retorno no go
	soma, subtracao, multiplicacao, divisao := calculosMatematicos(10, 20)
	fmt.Println("Resultado da soma:", soma)
	fmt.Println("Resultado da subtração:", subtracao)
	fmt.Println("Resultado da multiplicação:", multiplicacao)
	fmt.Println("Resultado da divisão:", divisao)

	// podemos ignorar os valores que não precisamos, usando o _
	soma2, _, _, _ := calculosMatematicos(10, 20)
	fmt.Println("Resultado da soma:", soma2)

	// funções multi retorno servem normalmente para retornar erro junto da função caso der alguma inconcistencia
	nome, idade, erro := mostrarNomeEIdade("João", 20)
	if erro != nil {
		fmt.Println("Erro:", erro)
	}

	fmt.Println("Nome:", nome)
	fmt.Println("Idade:", idade)
	fmt.Println("Erro:", erro)

	// ============================================
	// RETORNO NOMEADO
	// ============================================
	fmt.Println("--- RETORNO NOMEADO ---")
	fmt.Println("Podemos nomear os valores de retorno na assinatura da função.")
	fmt.Println("Isso permite usar 'return' sem especificar os valores explicitamente.\n")

	resultado3, erro2 := calcularComRetornoNomeado(10, 5)
	if erro2 != nil {
		fmt.Println("Erro:", erro2)
	} else {
		fmt.Printf("Resultado: %d e Erro: %v\n", resultado3, erro2)
	}

	resultado4, erro3 := calcularComRetornoNomeado(10, 0)
	if erro3 != nil {
		fmt.Printf("Resultado: %d e Erro: %v\n", resultado4, erro3)
	} else {
		fmt.Println("Não vai entrar aqui")
	}
}

func mostrarNomeEIdade(nome string, idade int) (string, int, error) {
	return nome, idade, nil
}

// Retorno nomeado: os nomes das variáveis de retorno são declarados na assinatura
// Podemos atribuir valores a essas variáveis e usar apenas 'return' sem especificar valores
func calcularComRetornoNomeado(a, b int) (resultado int, erro error) {
	if b == 0 {
		resultado = 0
		erro = fmt.Errorf("divisão por zero não permitida")
		return // Retorna os valores atuais de 'resultado' (0) e 'erro'
	}
	resultado = a / b
	erro = nil
	return // Retorna os valores atuais de 'resultado' e 'erro' (nil)
}
