package testes

import (
	"fmt"
)

// DemonstrarTestesBasicos demonstra a estrutura básica de testes em Go
func DemonstrarTestesBasicos() {
	fmt.Println("--- TESTES BÁSICOS ---")
	fmt.Println("Testes automatizados são funções que verificam se o código funciona corretamente.")
	fmt.Println("Em Go, os testes são escritos em arquivos que terminam com '_test.go'.\n")

	fmt.Println("Por que usar testes automatizados?")
	fmt.Println("  - Garantem que o código funciona como esperado")
	fmt.Println("  - Previnem regressões quando o código é modificado")
	fmt.Println("  - Servem como documentação do comportamento esperado")
	fmt.Println("  - Facilitam refatoração com confiança")
	fmt.Println("  - Melhoram a qualidade do código\n")

	fmt.Println("Estrutura básica de um teste:")
	fmt.Println("  1. Arquivo: nome_test.go (ex: calculadora_test.go)")
	fmt.Println("  2. Função: TestNomeDaFuncao (ex: TestSoma)")
	fmt.Println("  3. Parâmetro: *testing.T")
	fmt.Println("  4. Verificação: comparar resultado com esperado")
	fmt.Println("  5. Erro: usar t.Error() ou t.Fatal() se falhar\n")

	fmt.Println("Exemplo 1: Teste básico simples")
	fmt.Println("Estrutura mínima de um teste:\n")

	fmt.Println("  func TestSoma(t *testing.T) {")
	fmt.Println("      resultado := soma(2, 3)")
	fmt.Println("      esperado := 5")
	fmt.Println("      if resultado != esperado {")
	fmt.Println("          t.Errorf(\"Resultado: %d, Esperado: %d\", resultado, esperado)")
	fmt.Println("      }")
	fmt.Println("  }\n")

	fmt.Println("Exemplo 2: Teste com t.Error()")
	fmt.Println("t.Error() registra um erro mas continua executando o teste:\n")

	fmt.Println("  func TestMultiplicacao(t *testing.T) {")
	fmt.Println("      resultado := multiplicacao(2, 3)")
	fmt.Println("      esperado := 6")
	fmt.Println("      if resultado != esperado {")
	fmt.Println("          t.Error(\"Multiplicação falhou\")")
	fmt.Println("      }")
	fmt.Println("      // Código continua executando mesmo se houver erro")
	fmt.Println("  }\n")

	fmt.Println("Exemplo 3: Teste com t.Fatal()")
	fmt.Println("t.Fatal() registra um erro e para a execução do teste imediatamente:\n")

	fmt.Println("  func TestDivisao(t *testing.T) {")
	fmt.Println("      resultado := divisao(10, 0)")
	fmt.Println("      if resultado == 0 {")
	fmt.Println("          t.Fatal(\"Divisão por zero não tratada\")")
	fmt.Println("      }")
	fmt.Println("      // Este código não será executado se Fatal for chamado")
	fmt.Println("  }\n")

	fmt.Println("Diferença entre t.Error() e t.Fatal():")
	fmt.Println("  - t.Error(): Registra erro mas continua executando")
	fmt.Println("    ✓ Use quando quiser testar múltiplas condições")
	fmt.Println("    ✓ Use quando o erro não impede outros testes")
	fmt.Println("  - t.Fatal(): Registra erro e para imediatamente")
	fmt.Println("    ✓ Use quando o erro impede continuar o teste")
	fmt.Println("    ✓ Use para setup que deve funcionar\n")

	fmt.Println("Exemplo 4: Teste prático completo")
	fmt.Println("Testando uma função de validação:\n")

	// Exemplo prático
	validarEmail := func(email string) bool {
		return len(email) > 0 && email != ""
	}

	func() {
		emailValido := "teste@exemplo.com"
		resultado := validarEmail(emailValido)
		if !resultado {
			fmt.Printf("  ❌ Email válido foi rejeitado: %s\n", emailValido)
		} else {
			fmt.Printf("  ✓ Email válido aceito: %s\n", emailValido)
		}

		emailInvalido := ""
		resultado = validarEmail(emailInvalido)
		if resultado {
			fmt.Printf("  ❌ Email inválido foi aceito: %s\n", emailInvalido)
		} else {
			fmt.Printf("  ✓ Email inválido rejeitado: %s\n", emailInvalido)
		}
	}()
	fmt.Println()

	fmt.Println("Exemplo 5: Teste com mensagem descritiva")
	fmt.Println("Sempre inclua mensagens claras nos erros:\n")

	fmt.Println("  func TestSubtracao(t *testing.T) {")
	fmt.Println("      resultado := subtracao(10, 3)")
	fmt.Println("      esperado := 7")
	fmt.Println("      if resultado != esperado {")
	fmt.Println("          t.Errorf(\"Subtracao(10, 3) = %d, esperado %d\", resultado, esperado)")
	fmt.Println("      }")
	fmt.Println("  }\n")

	fmt.Println("⚠️  IMPORTANTE:")
	fmt.Println("  - Nome do arquivo deve terminar com '_test.go'")
	fmt.Println("  - Arquivos '_test.go' são IGNORADOS pelo compilador Go normal")
	fmt.Println("    Eles só são incluídos quando você executa 'go test'")
	fmt.Println("    Por isso, código didático/demonstrativo NÃO deve usar '_test.go'")
	fmt.Println("  - Nome da função deve começar com 'Test'")
	fmt.Println("  - Função deve receber *testing.T como parâmetro")
	fmt.Println("  - Use t.Error() para continuar após erro")
	fmt.Println("  - Use t.Fatal() para parar imediatamente")
	fmt.Println("  - Mensagens de erro devem ser descritivas")
	fmt.Println("  - Testes devem ser independentes e isolados")
	fmt.Println("  - Um teste não deve depender de outro")
	fmt.Println()
}
