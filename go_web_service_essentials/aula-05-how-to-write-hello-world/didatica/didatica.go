package didatica

import "fmt"

func DemonstrarHelloWorld() {
	fmt.Println("=== AULA 05: Hello World em Go ===")
	fmt.Println()

	fmt.Println("--- O que e um programa Go executavel? ---")
	fmt.Println("E um pacote chamado `main` com uma funcao chamada `main`.")
	fmt.Println("Quando voce roda `go run .`, o runtime procura `package main` e executa `func main()`.")
	fmt.Println()

	Exemplo1PackageMain()
	Exemplo2Imports()
	Exemplo3FuncMain()
	Exemplo4FmtPrintln()
	Exemplo5ExportadoENaoExportado()
	PontosImportantes()
}

func Exemplo1PackageMain() {
	fmt.Println("--- Exemplo 1: package main ---")
	fmt.Println("Todo arquivo Go comeca declarando um pacote.")
	fmt.Println("O pacote `main` e especial: ele indica que o projeto pode gerar um executavel.")
	fmt.Println()
}

func Exemplo2Imports() {
	fmt.Println("--- Exemplo 2: imports ---")
	fmt.Println("Go tem poucas coisas embutidas. Para imprimir texto, usamos o pacote `fmt`.")
	fmt.Println("O import deixa claro de onde vem cada funcionalidade usada no arquivo.")
	fmt.Println()
}

func Exemplo3FuncMain() {
	fmt.Println("--- Exemplo 3: func main ---")
	fmt.Println("`func` declara uma funcao em Go.")
	fmt.Println("`main` nao recebe parametros aqui, mas os parenteses ainda sao obrigatorios.")
	fmt.Println()
}

func Exemplo4FmtPrintln() {
	fmt.Println("--- Exemplo 4: fmt.Println ---")
	fmt.Println("Em Go, voce chama funcoes de outros pacotes usando `pacote.Funcao`.")
	fmt.Println("`fmt.Println` imprime na saida padrao e adiciona uma quebra de linha.")
	fmt.Println()
}

func Exemplo5ExportadoENaoExportado() {
	fmt.Println("--- Exemplo 5: exportado vs nao exportado ---")
	fmt.Println("Nomes que comecam com letra maiuscula sao exportados para outros pacotes.")
	fmt.Println("Por isso usamos `fmt.Println`, com P maiusculo.")
	fmt.Println("Nomes com letra minuscula ficam internos ao pacote.")
	fmt.Println()
}

func PontosImportantes() {
	fmt.Println("--- Pontos importantes ---")
	fmt.Println("- `package main` + `func main()` formam o ponto de entrada de um executavel.")
	fmt.Println("- `import` traz codigo de outros pacotes.")
	fmt.Println("- `fmt.Println` imprime com quebra de linha.")
	fmt.Println("- `gofmt` padroniza a formatacao do codigo Go.")
}
