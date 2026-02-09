package main

import "fmt"

func main() {
	numero1 := 10
	numero2 := 20

	// ============================================
	// OPERADORES ARITMÉTICOS (Binários)
	// ============================================
	fmt.Println("--- OPERADORES ARITMÉTICOS ---")
	fmt.Printf("Soma: %d + %d = %d\n", numero1, numero2, numero1+numero2)
	fmt.Printf("Subtração: %d - %d = %d\n", numero1, numero2, numero1-numero2)
	fmt.Printf("Multiplicação: %d * %d = %d\n", numero1, numero2, numero1*numero2)
	fmt.Printf("Divisão: %d / %d = %d\n", numero1, numero2, numero1/numero2)
	fmt.Printf("Resto (módulo): %d %% %d = %d\n", numero1, numero2, numero1%numero2)
	fmt.Println()

	// ============================================
	// OPERADORES DE ATRIBUIÇÃO COMPOSTA
	// ============================================
	fmt.Println("--- OPERADORES DE ATRIBUIÇÃO COMPOSTA ---")
	valor := 10
	fmt.Printf("Valor inicial: %d\n", valor)

	valor += 5
	fmt.Printf("Atribuição com soma: valor += 5 → valor = %d\n", valor)

	valor -= 3
	fmt.Printf("Atribuição com subtração: valor -= 3 → valor = %d\n", valor)

	valor *= 2
	fmt.Printf("Atribuição com multiplicação: valor *= 2 → valor = %d\n", valor)

	valor /= 4
	fmt.Printf("Atribuição com divisão: valor /= 4 → valor = %d\n", valor)

	valor %= 3
	fmt.Printf("Atribuição com resto: valor %%= 3 → valor = %d\n", valor)
	fmt.Println()

	// ============================================
	// OPERADORES UNÁRIOS
	// ============================================
	fmt.Println("--- OPERADORES UNÁRIOS ---")
	num := 10
	fmt.Printf("Valor inicial: %d\n", num)

	numAntes := num
	num++
	fmt.Printf("Incremento: %d++ = %d\n", numAntes, num)

	numAntes = num
	num--
	fmt.Printf("Decremento: %d-- = %d\n", numAntes, num)

	positivo := +5
	negativo := -5
	fmt.Printf("Operador unário positivo: +5 = %d\n", positivo)
	fmt.Printf("Operador unário negativo: -5 = %d\n", negativo)
	fmt.Println()

	// ============================================
	// OPERADORES DE COMPARAÇÃO (Relacionais)
	// ============================================
	fmt.Println("--- OPERADORES DE COMPARAÇÃO ---")
	a := 10
	b := 20
	fmt.Printf("Igual: %d == %d = %t\n", a, b, a == b)
	fmt.Printf("Diferente: %d != %d = %t\n", a, b, a != b)
	fmt.Printf("Maior: %d > %d = %t\n", a, b, a > b)
	fmt.Printf("Menor: %d < %d = %t\n", a, b, a < b)
	fmt.Printf("Maior ou igual: %d >= %d = %t\n", a, b, a >= b)
	fmt.Printf("Menor ou igual: %d <= %d = %t\n", a, b, a <= b)
	fmt.Println()

	// ============================================
	// OPERADORES LÓGICOS
	// ============================================
	fmt.Println("--- OPERADORES LÓGICOS ---")
	x := 10
	y := 20
	fmt.Printf("E (AND): %d == %d && %d < %d = %t\n", x, y, x, y, x == y && x < y)
	fmt.Printf("Ou (OR): %d == %d || %d < %d = %t\n", x, y, x, y, x == y || x < y)

	booleano := false
	fmt.Printf("Negação (NOT): !%t = %t\n", booleano, !booleano)
	fmt.Println()

	// ============================================
	// CONVERSÃO DE TIPOS
	// ============================================
	fmt.Println("--- CONVERSÃO DE TIPOS ---")
	var numero3 int16 = 10
	var numero4 int32 = 20
	fmt.Printf("numero3 (int16): %d\n", numero3)
	fmt.Printf("numero4 (int32): %d\n", numero4)

	// Não pode fazer isso pois são tipos diferentes
	// soma := numero3 + numero4

	// Assim podemos fazer pois convertemos de um tipo para outro
	var soma2 int32 = int32(numero3) + numero4
	fmt.Printf("Conversão: int32(%d) + %d = %d\n", numero3, numero4, soma2)

	//Nao tem operador ternario, tem que fazer if/else
}
