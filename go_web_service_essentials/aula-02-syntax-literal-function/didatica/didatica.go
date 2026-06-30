package didatica

import "fmt"

func DemonstrarLiteralFunction() {
	fmt.Println("=== AULA 02: Literal Function ===")
	fmt.Println()

	fmt.Println("--- O que e uma literal function? ---")
	fmt.Println("E uma funcao criada como valor, sem precisar declarar um nome com `func nome() {}`.")
	fmt.Println("Voce pode executar essa funcao na hora ou guardar em uma variavel para chamar depois.")
	fmt.Println()

	fmt.Println("--- Por que isso existe? ---")
	fmt.Println("Porque em Go funcoes tambem sao valores.")
	fmt.Println("Isso ajuda quando voce quer criar uma pequena funcao local, adiar uma execucao com `defer`, ou passar comportamento para outra funcao.")
	fmt.Println()

	Exemplo1ExecutarDireto()
	Exemplo2GuardarEmVariavel()
	Exemplo3DeferECapturaDeVariavel()
	PontosImportantes()
}

func Exemplo1ExecutarDireto() {
	fmt.Println("--- Exemplo 1: declarar e chamar na hora ---")

	n := 0
	func() {
		fmt.Println("Direct:", n)
	}()

	fmt.Println()
}

func Exemplo2GuardarEmVariavel() {
	fmt.Println("--- Exemplo 2: guardar a funcao em uma variavel ---")

	n := 0
	f := func() {
		fmt.Println("Variable:", n)
	}

	f()
	n = 3
	f()

	fmt.Println()
}

func Exemplo3DeferECapturaDeVariavel() {
	fmt.Println("--- Exemplo 3: defer executa no final e enxerga o valor atual ---")

	n := 0
	defer func() {
		fmt.Println("defer 1:", n)
	}()

	n = 3

	defer func() {
		fmt.Println("defer 2:", n)
	}()

	fmt.Println("antes dos defers:", n)
	fmt.Println("Quando esta funcao terminar, os defers rodam em ordem LIFO: o ultimo defer roda primeiro.")
}

func PontosImportantes() {
	fmt.Println()
	fmt.Println("--- Pontos importantes ---")
	fmt.Println("- `func() { ... }` cria uma funcao anonima.")
	fmt.Println("- `func() { ... }()` cria e executa na hora.")
	fmt.Println("- `f := func() { ... }` guarda a funcao para chamar depois.")
	fmt.Println("- `defer` adia a chamada ate a funcao atual terminar.")
	fmt.Println("- Funcoes anonimas capturam variaveis do escopo externo.")
}
