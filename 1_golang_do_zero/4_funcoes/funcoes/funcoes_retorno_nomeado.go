package funcoes

import "fmt"

func DemonstrarRetornoNomeado() {
	fmt.Println("--- RETORNO NOMEADO ---")
	fmt.Println("Podemos nomear os valores de retorno na assinatura da função.")
	fmt.Println("Isso permite usar 'return' sem especificar os valores explicitamente.\n")

	resultado3, erro2 := calcularComRetornoNomeado(10, 5)
	if erro2 != nil {
		fmt.Println("Erro:", erro2)
	} else {
		fmt.Printf("Resultado: %d e Erro: %v\n", resultado3, erro2)
	}

	resultado4, erro3 := calcularComRetornoNomeado(10, 0)
	if erro3 != nil {
		fmt.Printf("Resultado: %d e Erro: %v\n", resultado4, erro3)
	} else {
		fmt.Println("Não vai entrar aqui")
	}
	fmt.Println()
}

// Retorno nomeado: os nomes das variáveis de retorno são declarados na assinatura
// Podemos atribuir valores a essas variáveis e usar apenas 'return' sem especificar valores
func calcularComRetornoNomeado(a, b int) (resultado int, erro error) {
	if b == 0 {
		resultado = 0
		erro = fmt.Errorf("divisão por zero não permitida")
		return // Retorna os valores atuais de 'resultado' (0) e 'erro'
	}
	resultado = a / b
	erro = nil
	return // Retorna os valores atuais de 'resultado' e 'erro' (nil)
}
