package main

import "fmt"

func main() {
	fmt.Println("=== PONTEIROS EM GO ===\n")

	// ============================================
	// PASSAGEM POR VALOR (CÓPIA)
	// ============================================
	fmt.Println("--- PASSAGEM POR VALOR (CÓPIA) ---")
	fmt.Println("Quando atribuímos uma variável a outra sem usar ponteiros,")
	fmt.Println("Go cria uma CÓPIA do valor, não uma referência.\n")

	variavel1 := 10
	variavel2 := variavel1 // Cópia do valor

	fmt.Printf("variavel1 inicial: %d\n", variavel1)
	fmt.Printf("variavel2 (cópia de variavel1): %d\n", variavel2)
	fmt.Printf("Endereço de variavel1: %p\n", &variavel1)
	fmt.Printf("Endereço de variavel2: %p (diferente!)\n\n", &variavel2)

	variavel1++
	fmt.Println("Após incrementar variavel1:")
	fmt.Printf("variavel1: %d (modificada)\n", variavel1)
	fmt.Printf("variavel2: %d (não mudou - é uma cópia independente)\n", variavel2)
	fmt.Println("⚠️  variavel2 não foi afetada porque é apenas uma cópia do valor!\n")

	// ============================================
	// PONTEIROS - OPERADOR & (ENDEREÇO)
	// ============================================
	fmt.Println("--- PONTEIROS - OPERADOR & (ENDEREÇO) ---")
	fmt.Println("O operador & retorna o endereço de memória de uma variável.\n")

	variavel3 := &variavel1 // Ponteiro para variavel1
	fmt.Printf("variavel3 = &variavel1 (ponteiro para variavel1)\n")
	fmt.Printf("Endereço armazenado em variavel3: %p\n", variavel3)
	fmt.Printf("Endereço de variavel1: %p\n", &variavel1)
	fmt.Printf("São iguais? %t\n\n", variavel3 == &variavel1)

	// ============================================
	// PONTEIROS - OPERADOR * (DESREFERENCIAÇÃO)
	// ============================================
	fmt.Println("--- PONTEIROS - OPERADOR * (DESREFERENCIAÇÃO) ---")
	fmt.Println("O operador * acessa o valor armazenado no endereço de memória.\n")

	fmt.Printf("Valor de variavel1: %d\n", variavel1)
	fmt.Printf("Valor através do ponteiro (*variavel3): %d\n", *variavel3)
	fmt.Printf("São iguais? %t\n\n", variavel1 == *variavel3)

	// ============================================
	// MODIFICAÇÃO ATRAVÉS DE PONTEIROS
	// ============================================
	fmt.Println("--- MODIFICAÇÃO ATRAVÉS DE PONTEIROS ---")
	fmt.Println("Quando modificamos uma variável através de um ponteiro,")
	fmt.Println("a variável original também é modificada.\n")

	fmt.Printf("Antes: variavel1 = %d, *variavel3 = %d\n", variavel1, *variavel3)
	variavel1++
	fmt.Printf("Após variavel1++: variavel1 = %d, *variavel3 = %d\n", variavel1, *variavel3)
	fmt.Println("✓ variavel3 reflete a mudança porque aponta para o mesmo endereço!\n")

	// Modificando através do ponteiro
	*variavel3 = 100
	fmt.Printf("Após *variavel3 = 100: variavel1 = %d, *variavel3 = %d\n", variavel1, *variavel3)
	fmt.Println("✓ Modificar através do ponteiro também modifica a variável original!\n")

	// ============================================
	// EXEMPLO PRÁTICO - FUNÇÃO COM PONTEIRO
	// ============================================
	fmt.Println("--- EXEMPLO PRÁTICO - FUNÇÃO COM PONTEIRO ---")

	valor := 5
	fmt.Printf("Valor antes da função: %d\n", valor)

	// Função que modifica através de ponteiro
	incrementar(&valor)
	fmt.Printf("Valor depois da função: %d\n", valor)
	fmt.Println("✓ A função modificou a variável original através do ponteiro!\n")

	// ============================================
	// RESUMO DOS OPERADORES
	// ============================================
	fmt.Println("--- RESUMO DOS OPERADORES ---")
	fmt.Println("& (operador de endereço):")
	fmt.Println("  - Retorna o endereço de memória de uma variável")
	fmt.Println("  - Exemplo: ponteiro := &variavel")
	fmt.Println()
	fmt.Println("* (operador de desreferenciação):")
	fmt.Println("  - Acessa o valor armazenado no endereço")
	fmt.Println("  - Exemplo: valor := *ponteiro")
	fmt.Println()
	fmt.Println("Diferença:")
	fmt.Println("  - Sem ponteiro: cópia do valor (independente)")
	fmt.Println("  - Com ponteiro: referência ao mesmo endereço (compartilhado)")
}

// Função auxiliar para demonstrar uso de ponteiros em funções
func incrementar(ponteiro *int) {
	*ponteiro++
	fmt.Printf("  (Dentro da função: incrementando valor através do ponteiro)\n")
}
