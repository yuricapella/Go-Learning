package interfaces

import "fmt"

func DemonstrarTypeAssertions() {
	fmt.Println("--- TYPE ASSERTIONS E TYPE SWITCHES ---")
	fmt.Println("Type assertions permitem verificar e converter tipos de interfaces.")
	fmt.Println("Type switches permitem verificar múltiplos tipos de uma vez.\n")

	fmt.Println("Exemplo 1: Type assertion básica")
	fmt.Println("Sintaxe: valor.(Tipo) - retorna o valor convertido ou causa panic se falhar\n")

	var valorQualquer interface{} = "texto"
	textoConvertido := valorQualquer.(string)
	fmt.Printf("valorQualquer como string: %s\n", textoConvertido)
	fmt.Println()

	fmt.Println("Exemplo 2: Type assertion segura")
	fmt.Println("Sintaxe: valor, ok := valor.(Tipo) - retorna ok=false se falhar\n")

	valorQualquer = 42
	numeroConvertido, conversaoSucedida := valorQualquer.(int)
	if conversaoSucedida {
		fmt.Printf("Conversão bem-sucedida: %d\n", numeroConvertido)
	} else {
		fmt.Println("Conversão falhou")
	}
	fmt.Println()

	fmt.Println("Exemplo 3: Type assertion que falha (com verificação)")
	valorQualquer = "texto"
	numeroConvertido, conversaoSucedida = valorQualquer.(int)
	if conversaoSucedida {
		fmt.Printf("Número: %d\n", numeroConvertido)
	} else {
		fmt.Printf("Não é possível converter '%v' para int\n", valorQualquer)
	}
	fmt.Println()

	fmt.Println("Exemplo 4: Type switch")
	fmt.Println("Permite verificar múltiplos tipos de uma vez:\n")

	processarTipoDiferente("texto")
	processarTipoDiferente(42)
	processarTipoDiferente(true)
	processarTipoDiferente(3.14)
	fmt.Println()

	fmt.Println("⚠️  IMPORTANTE:")
	fmt.Println("  - Type assertion: valor.(Tipo) - causa panic se falhar")
	fmt.Println("  - Type assertion segura: valor, ok := valor.(Tipo) - retorna ok=false se falhar")
	fmt.Println("  - Type switch: switch v := valor.(type) { case Tipo: ... }")
	fmt.Println("  - Type switch é mais seguro e legível para múltiplos tipos")
	fmt.Println()
}

func processarTipoDiferente(valorQualquer interface{}) {
	switch valorVerificado := valorQualquer.(type) {
	case string:
		fmt.Printf("É string: '%s'\n", valorVerificado)
	case int:
		fmt.Printf("É int: %d\n", valorVerificado)
	case bool:
		fmt.Printf("É bool: %t\n", valorVerificado)
	case float64:
		fmt.Printf("É float64: %.2f\n", valorVerificado)
	default:
		fmt.Printf("Tipo desconhecido: %T\n", valorVerificado)
	}
}
