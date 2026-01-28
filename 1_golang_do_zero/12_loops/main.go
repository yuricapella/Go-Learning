package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== LOOPS EM GO ===")
	fmt.Println()

	// ============================================
	// FOR BÁSICO (WHILE-LIKE)
	// ============================================
	fmt.Println("--- FOR BÁSICO (SIMILAR AO WHILE) ---")
	forBasico()
	fmt.Println()

	// ============================================
	// FOR TRADICIONAL (INICIALIZAÇÃO, CONDIÇÃO, INCREMENTO)
	// ============================================
	fmt.Println("--- FOR TRADICIONAL (INICIALIZAÇÃO, CONDIÇÃO, INCREMENTO) ---")
	forTradicional()
	fmt.Println()

	// ============================================
	// FOR RANGE - SLICES/ARRAYS
	// ============================================
	fmt.Println("--- FOR RANGE - SLICES/ARRAYS ---")
	forRange()
	fmt.Println()

	// ============================================
	// FOR RANGE - IGNORANDO VALORES COM _
	// ============================================
	fmt.Println("--- FOR RANGE - IGNORANDO VALORES COM _ ---")
	forRangeComValorOculto()
	fmt.Println()

	// ============================================
	// FOR RANGE - STRINGS
	// ============================================
	fmt.Println("--- FOR RANGE - STRINGS ---")
	forRangeString()
	fmt.Println()

	// ============================================
	// FOR RANGE - MAPS
	// ============================================
	fmt.Println("--- FOR RANGE - MAPS ---")
	forRangeMap()
	fmt.Println()

	// ============================================
	// BREAK E CONTINUE
	// ============================================
	fmt.Println("--- BREAK E CONTINUE ---")
	exemploBreakContinue()
	fmt.Println()

	// ============================================
	// FOR INFINITO
	// ============================================
	fmt.Println("--- FOR INFINITO ---")
	forInfinito()
	fmt.Println()

	// ============================================
	// RESUMO
	// ============================================
	fmt.Println("--- RESUMO ---")
	fmt.Println("Tipos de loops em Go:")
	fmt.Println("  ✓ for condição { } - Similar ao while")
	fmt.Println("  ✓ for init; condição; incremento { } - For tradicional")
	fmt.Println("  ✓ for range coleção { } - Itera sobre coleções")
	fmt.Println("  ✓ for { } - Loop infinito")
	fmt.Println()
	fmt.Println("Controle de fluxo:")
	fmt.Println("  ✓ break - Sai do loop")
	fmt.Println("  ✓ continue - Pula para próxima iteração")
	fmt.Println()
	fmt.Println("Range funciona com:")
	fmt.Println("  ✓ Slices e Arrays")
	fmt.Println("  ✓ Maps")
	fmt.Println("  ✓ Strings (retorna rune e índice)")
	fmt.Println("  ✗ Structs (não são iteráveis)")
}

// ============================================
// FOR BÁSICO (SIMILAR AO WHILE)
// ============================================
func forBasico() {
	fmt.Println("Loop com condição simples (similar ao while):")
	fmt.Println("Condição: i < 10 (vai rodar enquanto i for menor que 10)")
	fmt.Println()

	i := 0
	fmt.Printf("Valor inicial de i: %d\n", i)
	fmt.Println()

	for i < 10 {
		i++
		fmt.Printf("  Iteração: i++ executado, agora i = %d (condição i < 10: %t)\n", i, i < 10)
		time.Sleep(time.Millisecond * 10)
	}
	fmt.Printf("Loop finalizado! Valor final de i: %d (condição i < 10 agora é %t)\n", i, i < 10)
}

// ============================================
// FOR TRADICIONAL
// ============================================
func forTradicional() {
	fmt.Println("For tradicional com inicialização, condição e incremento:")
	fmt.Println("Sintaxe: for inicialização; condição; incremento { }")
	fmt.Println()
	fmt.Println("⚠️  IMPORTANTE: O incremento acontece APÓS cada iteração!")
	fmt.Println("   Fluxo: inicialização → verifica condição → executa bloco → incremento → verifica condição...")
	fmt.Println()

	fmt.Println("Exemplo 1: Usando variável local")
	fmt.Println("Condição: j < 5 (vai rodar 5 vezes: j = 0, 1, 2, 3, 4)")
	for j := 0; j < 5; j++ {
		fmt.Printf("  Dentro do loop: j = %d\n", j)
		time.Sleep(time.Millisecond * 10)
		// Após executar este bloco, j++ é executado automaticamente
		// Depois verifica se j < 5, se sim, continua; se não, sai do loop
	}
	fmt.Println("Após o loop, j não existe mais (escopo local)")
	fmt.Println()

	fmt.Println("Exemplo 2: Usando ponteiro para variável externa")
	fmt.Println("Utilizando ponteiro para modificar variável de fora do escopo do for")
	i := 0
	for j := &i; *j < 5; *j++ {
		fmt.Printf("  Dentro do loop: valor de i = %d\n", *j)
		time.Sleep(time.Millisecond * 10)
	}
	fmt.Printf("Valor final de i (modificado pelo ponteiro): %d\n", i)
	fmt.Println()
	fmt.Println("Resumo: O incremento (j++ ou *j++) sempre acontece após cada iteração, assim que acaba o loop, j nao sabe do ultimo incremento, apenas a variavel externa sabe que 4 virou 5")
}

// ============================================
// FOR RANGE - SLICES/ARRAYS
// ============================================
func forRange() {
	fmt.Println("Range é usado para iterar sobre slices, arrays, maps, strings, etc.")
	fmt.Println()

	letras := []string{"a", "b", "c", "d", "e"}
	fmt.Printf("Slice de letras: %v\n", letras)
	fmt.Printf("Tamanho do slice: %d elementos\n", len(letras))
	fmt.Println()

	fmt.Println("Iterando sobre slice com índice e valor:")
	fmt.Println("Sintaxe: for indice, valor := range coleção { }")
	for indice, valor := range letras {
		fmt.Printf("  Índice: %d, Valor: %s\n", indice, valor)
	}
	fmt.Printf("Loop finalizado após iterar sobre %d elementos\n", len(letras))
}

// ============================================
// FOR RANGE - IGNORANDO VALORES
// ============================================
func forRangeComValorOculto() {
	fmt.Println("Usando _ para ignorar valores que não precisamos:")
	fmt.Println("O _ (underscore) é usado quando não precisamos de um dos valores retornados pelo range")
	fmt.Println()

	letras := []string{"a", "b", "c", "d", "e"}
	fmt.Printf("Slice de letras: %v\n", letras)
	fmt.Println()

	fmt.Println("Exemplo 1: Ignorando o índice (apenas valores):")
	fmt.Println("Sintaxe: for _, valor := range coleção { }")
	for _, valor := range letras {
		fmt.Printf("  Valor: %s\n", valor)
	}
	fmt.Println()

	fmt.Println("Exemplo 2: Ignorando o valor (apenas índices):")
	fmt.Println("Sintaxe: for indice := range coleção { }")
	for indice := range letras {
		fmt.Printf("  Índice: %d\n", indice)
	}
}

// ============================================
// FOR RANGE - STRINGS
// ============================================
func forRangeString() {
	fmt.Println("Iterando sobre strings:")
	fmt.Println("⚠️  Range em string retorna o índice e a RUNE (código Unicode), não o byte!")
	fmt.Println("⚠️  O valor retornado é o código Unicode (rune), não o caractere diretamente!")
	fmt.Println()

	texto := "Hello"
	fmt.Printf("Texto: %s\n", texto)
	fmt.Printf("Tamanho em bytes: %d\n", len(texto))
	fmt.Println()

	fmt.Println("Iterando sobre cada caractere:")
	fmt.Println("Sintaxe: for indice, rune := range string { }")
	for indice, rune := range texto {
		fmt.Printf("  Índice: %d\n", indice)
		fmt.Printf("    Rune (código Unicode): %d\n", rune)
		fmt.Printf("    Caractere convertido: string(%d) = %s\n", rune, string(rune))
		fmt.Println()
	}

	fmt.Println("Exemplo com caracteres especiais (Unicode):")
	texto2 := "Olá, 世界"
	fmt.Printf("Texto: %s\n", texto2)
	fmt.Printf("Tamanho em bytes: %d (maior que número de caracteres devido ao Unicode)\n", len(texto2))
	fmt.Println("Iterando sobre cada caractere Unicode:")
	for indice, rune := range texto2 {
		fmt.Printf("  Índice: %d, Rune: %d, Caractere: %s\n", indice, rune, string(rune))
	}
}

// ============================================
// FOR RANGE - MAPS
// ============================================
func forRangeMap() {
	fmt.Println("Iterando sobre maps:")
	fmt.Println("⚠️  A ordem de iteração em maps é ALEATÓRIA! Não confie na ordem!")
	fmt.Println()

	mapa := map[string]string{
		"nome":  "João",
		"idade": "20",
		"email": "joao@example.com",
	}
	fmt.Printf("Map criado com %d elementos: %v\n", len(mapa), mapa)
	fmt.Println()

	fmt.Println("Exemplo 1: Iterando com chave e valor:")
	fmt.Println("Sintaxe: for chave, valor := range mapa { }")
	for chave, valor := range mapa {
		fmt.Printf("  Chave: %s, Valor: %s\n", chave, valor)
	}
	fmt.Println("⚠️  A ordem pode ser diferente a cada execução!")
	fmt.Println()

	fmt.Println("Exemplo 2: Iterando apenas com chaves:")
	fmt.Println("Sintaxe: for chave := range mapa { }")
	for chave := range mapa {
		fmt.Printf("  Chave: %s\n", chave)
	}
}

// ============================================
// BREAK E CONTINUE
// ============================================
func exemploBreakContinue() {
	fmt.Println("BREAK - Sai do loop completamente:")
	fmt.Println("Loop: for i := 0; i < 10; i++")
	fmt.Println("Condição do loop: i < 10 (vai rodar de 0 até 9)")
	fmt.Println("Condição do break: if i == 5 (sai quando i for igual a 5)")
	fmt.Println()

	for i := 0; i < 10; i++ {
		fmt.Printf("  Iteração: i = %d (condição i < 10: %t)\n", i, i < 10)
		if i == 5 {
			fmt.Printf("  → Condição i == 5 é verdadeira! Executando break...\n")
			fmt.Printf("  → BREAK: Saindo do loop completamente (não executa mais iterações)\n")
			break
		}
		fmt.Printf("  → Continuando loop...\n")
	}
	fmt.Printf("Loop finalizado após break! (não chegou até i = 9)\n")
	fmt.Println()

	fmt.Println("CONTINUE - Pula para a próxima iteração:")
	fmt.Println("Loop: for i := 0; i < 10; i++")
	fmt.Println("Condição do loop: i < 10 (vai rodar de 0 até 9)")
	fmt.Println("Condição do continue: if i%%2 == 0 (pula números pares)")
	fmt.Println("Números pares que serão pulados: 0, 2, 4, 6, 8")
	fmt.Println("Números ímpares que serão impressos: 1, 3, 5, 7, 9")
	fmt.Println()

	for i := 0; i < 10; i++ {
		fmt.Printf("  Iteração: i = %d\n", i)
		if i%2 == 0 {
			fmt.Printf("  → Condição i%%2 == 0 é verdadeira! Executando continue...\n")
			fmt.Printf("  → CONTINUE: Pulando para próxima iteração (não executa o resto do código)\n")
			continue
		}
		fmt.Printf("  → Número ímpar encontrado: i = %d\n", i)
	}
	fmt.Printf("Loop finalizado! (executou todas as iterações, mas pulou os pares)\n")
}

// ============================================
// FOR INFINITO
// ============================================
func forInfinito() {
	fmt.Println("Loop infinito com for { }")
	fmt.Println("Sintaxe: for { } - não tem condição, roda infinitamente")
	fmt.Println("⚠️  Use break para sair, senão rodará para sempre!")
	fmt.Println()

	contador := 0
	fmt.Println("Iniciando loop infinito...")
	for {
		contador++
		fmt.Printf("  Iteração %d: Loop infinito (sem condição de parada)\n", contador)
		time.Sleep(time.Millisecond * 100)

		if contador >= 3 {
			fmt.Printf("  → Condição contador >= 3 é verdadeira! (contador = %d)\n", contador)
			fmt.Println("  → Executando break para sair do loop infinito")
			break
		}
		fmt.Printf("  → Continuando loop... (contador = %d, condição contador >= 3: %t)\n", contador, contador >= 3)
	}
	fmt.Printf("Loop finalizado após break! Total de iterações: %d\n", contador)
}

// ============================================
// NOTA SOBRE STRUCTS
// ============================================
// Structs não são iteráveis com range.
// Structs são tipos de dados compostos, não são coleções.
// Para iterar sobre campos de uma struct, você precisaria usar
// reflection (reflect package) ou acessar os campos manualmente.
