package main

import "fmt"

func main() {
	// ============================================
	// ARRAYS - CONCEITOS BÁSICOS
	// ============================================
	fmt.Println("--- ARRAYS ---")
	fmt.Println("Arrays são fixos: tamanho definido na declaração")
	fmt.Println("Não podemos adicionar ou remover elementos após criação\n")

	// Array de strings com valores zero
	fmt.Println("1. Array de strings inicializado com valores zero:")
	var arrayString [5]string
	fmt.Printf("   arrayString inicial: %q\n", arrayString)
	fmt.Printf("   Tamanho: %d\n", len(arrayString))

	arrayString[0] = "Posição 1"
	fmt.Printf("   arrayString após atribuição: %q\n", arrayString)
	fmt.Println()

	// Array de inteiros com valores zero
	fmt.Println("2. Array de inteiros inicializado com valores zero:")
	var arrayInt [5]int
	fmt.Printf("   arrayInt inicial: %v\n", arrayInt)
	fmt.Printf("   Tamanho: %d\n", len(arrayInt))

	arrayInt[0] = 10
	fmt.Printf("   arrayInt após atribuição: %v\n", arrayInt)
	fmt.Println()

	// ============================================
	// ARRAYS - INICIALIZAÇÃO COM VALORES
	// ============================================
	fmt.Println("--- ARRAYS - INICIALIZAÇÃO COM VALORES ---")

	fmt.Println("3. Array inicializado com valores literais:")
	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Printf("   array2: %q\n", array2)
	fmt.Printf("   Tamanho: %d\n", len(array2))
	fmt.Println()

	// ============================================
	// ARRAYS - TAMANHO AUTOMÁTICO
	// ============================================
	fmt.Println("--- ARRAYS - TAMANHO AUTOMÁTICO ---")

	fmt.Println("4. Array com tamanho inferido automaticamente ([...]):")
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("   array3: %v\n", array3)
	fmt.Printf("   Tamanho: %d (inferido automaticamente)\n", len(array3))
	fmt.Println()

	// ============================================
	// SLICES - CONCEITOS BÁSICOS
	// ============================================
	fmt.Println("--- SLICES ---")
	fmt.Println("Slices são arrays dinâmicos:")
	fmt.Println("- Não precisamos informar o tamanho")
	fmt.Println("- Podem ter elementos adicionados ou removidos")
	fmt.Println("- São mais flexíveis que arrays\n")

	fmt.Println("5. Slice inicializado com valores literais:")
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("   slice inicial: %v\n", slice)
	fmt.Printf("   Tamanho: %d, Capacidade: %d\n", len(slice), cap(slice))
	fmt.Println()

	// ============================================
	// SLICES - ADICIONANDO ELEMENTOS (APPEND)
	// ============================================
	fmt.Println("--- SLICES - ADICIONANDO ELEMENTOS (APPEND) ---")

	fmt.Println("6. Adicionando elementos com append:")
	slice = append(slice, 6)
	fmt.Printf("   slice após append(6): %v\n", slice)
	fmt.Printf("   Tamanho: %d, Capacidade: %d\n", len(slice), cap(slice))

	slice = append(slice, 7, 8, 9)
	fmt.Printf("   slice após append(7, 8, 9): %v\n", slice)
	fmt.Printf("   Tamanho: %d, Capacidade: %d\n", len(slice), cap(slice))
	fmt.Println()

	// ============================================
	// SLICES - CRIAÇÃO COM MAKE
	// ============================================
	fmt.Println("--- SLICES - CRIAÇÃO COM MAKE ---")

	fmt.Println("7. Criando slice com make (tamanho e capacidade):")
	slice2 := make([]int, 5) // tamanho 5, capacidade 5
	fmt.Printf("   slice2 (make([]int, 5)): %v\n", slice2)
	fmt.Printf("   Tamanho: %d, Capacidade: %d\n", len(slice2), cap(slice2))

	slice2ComCap := make([]int, 3, 10) // tamanho 3, capacidade 10
	fmt.Printf("   slice2ComCap (make([]int, 3, 10)): %v\n", slice2ComCap)
	fmt.Printf("   Tamanho: %d, Capacidade: %d\n", len(slice2ComCap), cap(slice2ComCap))
	fmt.Println()

	// ============================================
	// SLICES - SLICING DE ARRAYS (REFERÊNCIA)
	// ============================================
	fmt.Println("--- SLICES - SLICING DE ARRAYS (REFERÊNCIA) ---")
	fmt.Println("Slices criados a partir de arrays compartilham a mesma memória:")
	fmt.Println()

	fmt.Printf("   array3 original: %v\n", array3)
	slice3 := array3[1:3] // elementos do índice 1 até 3 (exclusive)
	fmt.Printf("   slice3 = array3[1:3]: %v\n", slice3)
	fmt.Printf("   slice3 - Tamanho: %d, Capacidade: %d\n", len(slice3), cap(slice3))
	fmt.Println()

	fmt.Println("8. Modificando o array original:")
	fmt.Printf("   Antes: array3[1] = %d\n", array3[1])
	array3[1] = 20
	fmt.Printf("   Depois: array3[1] = %d\n", array3[1])
	fmt.Printf("   slice3 (referência compartilhada): %v\n", slice3)
	fmt.Println("   ⚠️  slice3 também foi modificado porque compartilha a memória!")
	fmt.Println()

	// ============================================
	// RESUMO DAS DIFERENÇAS
	// ============================================
	fmt.Println("--- RESUMO DAS DIFERENÇAS ---")
	fmt.Println("Arrays:")
	fmt.Println("  ✓ Tamanho fixo definido na declaração")
	fmt.Println("  ✓ Não podem crescer ou diminuir")
	fmt.Println("  ✓ Passados por valor (cópia)")
	fmt.Println("  ✓ Sintaxe: [tamanho]tipo")
	fmt.Println()
	fmt.Println("Slices:")
	fmt.Println("  ✓ Tamanho dinâmico")
	fmt.Println("  ✓ Podem crescer com append()")
	fmt.Println("  ✓ Passados por referência")
	fmt.Println("  ✓ Sintaxe: []tipo")
	fmt.Println("  ✓ Podem ser criados a partir de arrays")
}
