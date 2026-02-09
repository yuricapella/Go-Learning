package funcoes

import "fmt"

func DemonstrarFuncoesVariadicas() {
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
	fmt.Println()
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
