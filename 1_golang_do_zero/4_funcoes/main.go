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

	// ============================================
	// FUNÇÕES ANÔNIMAS
	// ============================================
	fmt.Println("--- FUNÇÕES ANÔNIMAS ---")
	fmt.Println("Existem três formas de usar funções anônimas em Go:\n")

	fmt.Println("1. Função anônima atribuída a variável (sem parâmetros):")
	var f = func() {
		fmt.Println("Função f")
	}
	f()
	fmt.Println("Tipo de f:", reflect.TypeOf(f))
	fmt.Println()

	fmt.Println("2. Função anônima atribuída a variável (com parâmetros):")
	var f2 = func(txt string) string {
		return txt
	}
	resultado2 := f2("Função f2")
	fmt.Println("Resultado da função f2:", resultado2)
	fmt.Println("Tipo de f2:", reflect.TypeOf(f2))
	fmt.Println()

	fmt.Println("3. Função anônima executada imediatamente (IIFE - Immediately Invoked Function Expression):")
	fmt.Println("A função é declarada e executada na mesma linha, passando parâmetros após a definição.")
	retorno := func(texto string) string {
		return fmt.Sprintf("recebido -> %s", texto)
	}("Passando parametro")
	fmt.Println(retorno)
	fmt.Println()

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

	// ============================================
	// FUNÇÃO VARIÁDICA
	// ============================================
	fmt.Println("--- FUNÇÃO VARIÁDICA ---")
	fmt.Println("Funções variádicas aceitam um número variável de argumentos.")
	fmt.Println("Usamos '...' antes do tipo do último parâmetro.")
	fmt.Println("O parâmetro variádico deve ser o último na lista de parâmetros.\n")

	somaVariadica := somarNumeros(1, 2, 3)
	fmt.Printf("somarNumeros(1, 2, 3) = %d\n", somaVariadica)

	somaVariadica2 := somarNumeros(10, 20, 30, 40, 50)
	fmt.Printf("somarNumeros(10, 20, 30, 40, 50) = %d\n", somaVariadica2)

	numeros := []int{5, 10, 15, 20}
	somaVariadica3 := somarNumeros(numeros...)
	fmt.Printf("somarNumeros([]int{5, 10, 15, 20}...) = %d\n", somaVariadica3)
	fmt.Println("⚠️  Para passar um slice como argumentos variádicos, usamos '...' após o slice\n")

	fmt.Println("Exemplo com strings:")
	concatenar := concatenarStrings("Olá", " ", "Mundo", "!")
	fmt.Printf("concatenarStrings(\"Olá\", \" \", \"Mundo\", \"!\") = %s\n", concatenar)

	// ============================================
	// FUNÇÃO RECURSIVA
	// ============================================
	fmt.Println("--- FUNÇÃO RECURSIVA ---")
	fmt.Println("Função recursiva é uma função que chama a si mesma.")
	fmt.Println("É importante ter uma condição de parada (caso base) para evitar loop infinito.\n")

	fmt.Println("Exemplo 1: Fatorial")
	fmt.Println("Fatorial de n (n!) = n * (n-1) * (n-2) * ... * 1")
	fmt.Println("Caso base: 0! = 1 e 1! = 1\n")

	fatorial5 := calcularFatorial(5)
	fmt.Printf("calcularFatorial(5) = %d\n", fatorial5)

	fatorial0 := calcularFatorial(0)
	fmt.Printf("calcularFatorial(0) = %d\n", fatorial0)

	fatorial7 := calcularFatorial(7)
	fmt.Printf("calcularFatorial(7) = %d\n", fatorial7)
	fmt.Println()

	fmt.Println("Exemplo 2: Sequência de Fibonacci")
	fmt.Println("Fibonacci: cada número é a soma dos dois anteriores")
	fmt.Println("F(0) = 0, F(1) = 1, F(n) = F(n-1) + F(n-2)\n")

	fibonacci10 := calcularFibonacci(10)
	fmt.Printf("calcularFibonacci(10) = %d\n", fibonacci10)

	fibonacci7 := calcularFibonacci(7)
	fmt.Printf("calcularFibonacci(7) = %d\n", fibonacci7)
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

// Função variádica: aceita um número variável de argumentos
// O parâmetro variádico deve ser o último e usa '...' antes do tipo
// Dentro da função, o parâmetro variádico é tratado como um slice
func somarNumeros(numeros ...int) int {
	soma := 0
	for _, num := range numeros {
		soma += num
	}
	return soma
}

// Função variádica com strings
func concatenarStrings(strings ...string) string {
	resultado := ""
	for _, str := range strings {
		resultado += str
	}
	return resultado
}

// Função recursiva: chama a si mesma até atingir o caso base
// Caso base: fatorial de 0 ou 1 é 1
// Caso recursivo: n! = n * (n-1)!
func calcularFatorial(n int) int {
	if n <= 1 {
		return 1 // Caso base: para a recursão
	}
	return n * calcularFatorial(n-1) // Chamada recursiva
}

// Função recursiva: Sequência de Fibonacci
// Caso base: F(0) = 0, F(1) = 1
// Caso recursivo: F(n) = F(n-1) + F(n-2)
func calcularFibonacci(n int) int {
	if n <= 1 {
		return n // Caso base: F(0) = 0, F(1) = 1
	}
	return calcularFibonacci(n-1) + calcularFibonacci(n-2) // Chamada recursiva
}
