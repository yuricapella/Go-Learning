package funcoes

import "fmt"

func DemonstrarFuncoesRecursivas() {
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
	fmt.Println()
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
