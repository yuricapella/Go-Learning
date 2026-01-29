package funcoes

import (
	"fmt"
	"reflect"
)

func DemonstrarFuncoesAnonimas() {
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
	fmt.Println("⚠️  Este padrão é conhecido como IIFE (Immediately Invoked Function Expression)")
	fmt.Println()
}
