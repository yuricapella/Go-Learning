package testes

import (
	"fmt"
)

// DemonstrarCobertura demonstra análise de cobertura de código
func DemonstrarCobertura() {
	fmt.Println("--- COBERTURA DE CÓDIGO ---")
	fmt.Println("Cobertura de código mede quantas linhas do código foram executadas pelos testes.")
	fmt.Println("É uma métrica importante para garantir qualidade dos testes.\n")

	fmt.Println("Por que medir cobertura?")
	fmt.Println("  - Identifica código não testado")
	fmt.Println("  - Ajuda a encontrar gaps nos testes")
	fmt.Println("  - Garante que mudanças são testadas")
	fmt.Println("  - Melhora confiança no código")
	fmt.Println("  - Facilita manutenção futura\n")

	fmt.Println("Comando básico de cobertura:")
	fmt.Println("  go test --cover")
	fmt.Println("  - Mostra porcentagem geral de cobertura")
	fmt.Println("  - Indica quantas linhas foram executadas")
	fmt.Println("  - Saída: coverage: 85.7%% of statements\n")

	fmt.Println("Exemplo 1: Verificar cobertura básica")
	fmt.Println("Executando testes com flag --cover:\n")

	fmt.Println("  $ go test --cover")
	fmt.Println("  PASS")
	fmt.Println("  coverage: 90.0%% of statements")
	fmt.Println("  ok      github.com/projeto/pacote   1.279s\n")

	fmt.Println("Comando 2: Gerar perfil de cobertura")
	fmt.Println("  go test --coverprofile=cobertura.txt")
	fmt.Println("  - Cria arquivo com detalhes da cobertura")
	fmt.Println("  - Formato não é legível diretamente")
	fmt.Println("  - Precisa ser analisado com go tool cover\n")

	fmt.Println("Exemplo 2: Gerar e visualizar perfil")
	fmt.Println("Passos para análise detalhada:\n")

	fmt.Println("  Passo 1: Gerar perfil")
	fmt.Println("  $ go test --coverprofile=cobertura.txt\n")

	fmt.Println("  Passo 2: Ver cobertura por função")
	fmt.Println("  $ go tool cover --func=cobertura.txt")
	fmt.Println("  github.com/projeto/pacote/funcao.go:5:   Soma           100.0%%")
	fmt.Println("  github.com/projeto/pacote/funcao.go:10:  Subtracao      90.0%%")
	fmt.Println("  total:                                              (statements)    95.0%%\n")

	fmt.Println("  Passo 3: Visualizar em HTML (recomendado)")
	fmt.Println("  $ go tool cover --html=cobertura.txt")
	fmt.Println("  - Abre arquivo HTML no navegador")
	fmt.Println("  - Mostra código com cores:")
	fmt.Println("    Verde: código coberto pelos testes")
	fmt.Println("    Vermelho: código não coberto")
	fmt.Println("    Cinza: código não executável\n")

	fmt.Println("Interpretando relatórios de cobertura:")
	fmt.Println("  - 100%%: Todas as linhas executáveis foram testadas")
	fmt.Println("  - 80-99%%: Boa cobertura, alguns gaps menores")
	fmt.Println("  - 50-79%%: Cobertura moderada, precisa melhorar")
	fmt.Println("  - <50%%: Cobertura baixa, muitos gaps\n")

	fmt.Println("Exemplo 3: Análise de cobertura por função")
	fmt.Println("Identificando funções com baixa cobertura:\n")

	fmt.Println("  $ go tool cover --func=cobertura.txt")
	fmt.Println("  github.com/projeto/pacote/validacao.go:5:   ValidarEmail    100.0%%")
	fmt.Println("  github.com/projeto/pacote/validacao.go:15:  ValidarTelefone 50.0%%")
	fmt.Println("  github.com/projeto/pacote/validacao.go:25:  ValidarCPF      0.0%%")
	fmt.Println("  total:                                              (statements)    50.0%%\n")

	fmt.Println("  Análise:")
	fmt.Println("  - ValidarEmail: totalmente coberta ✓")
	fmt.Println("  - ValidarTelefone: parcialmente coberta (precisa mais testes)")
	fmt.Println("  - ValidarCPF: não coberta (precisa criar testes)\n")

	fmt.Println("Exemplo 4: Visualização HTML detalhada")
	fmt.Println("O HTML mostra exatamente quais linhas não foram cobertas:\n")

	fmt.Println("  Código no HTML:")
	fmt.Println("  func ValidarTelefone(telefone string) bool {")
	fmt.Println("      if len(telefone) == 0 {  // Verde: coberto")
	fmt.Println("          return false")
	fmt.Println("      }")
	fmt.Println("      if len(telefone) < 10 {  // Vermelho: não coberto")
	fmt.Println("          return false")
	fmt.Println("      }")
	fmt.Println("      return true  // Verde: coberto")
	fmt.Println("  }\n")

	fmt.Println("Boas práticas de cobertura:")
	fmt.Println("  ✓ Aponte para alta cobertura (80-100%%)")
	fmt.Println("  ✓ Foque em código crítico primeiro")
	fmt.Println("  ✓ Teste casos de erro e edge cases")
	fmt.Println("  ✓ Não sacrifique qualidade por quantidade")
	fmt.Println("  ✓ Use cobertura como guia, não como meta absoluta\n")

	fmt.Println("Limitações da cobertura:")
	fmt.Println("  - Cobertura não garante qualidade dos testes")
	fmt.Println("  - Código coberto pode ainda ter bugs")
	fmt.Println("  - Não testa lógica de negócio complexa")
	fmt.Println("  - Foque em testes significativos, não apenas cobertura\n")

	fmt.Println("Comandos úteis combinados:")
	fmt.Println("  go test --cover --coverprofile=cobertura.txt")
	fmt.Println("  - Gera cobertura e perfil em um comando\n")

	fmt.Println("  go test -v --cover ./...")
	fmt.Println("  - Verbose + cobertura em todos os pacotes\n")

	fmt.Println("  go test --coverprofile=cobertura.txt && go tool cover --html=cobertura.txt")
	fmt.Println("  - Gera perfil e abre HTML automaticamente\n")

	fmt.Println("⚠️  IMPORTANTE:")
	fmt.Println("  - Cobertura é uma métrica, não um objetivo")
	fmt.Println("  - Foque em testes significativos, não apenas números")
	fmt.Println("  - Use go tool cover --html para visualização detalhada")
	fmt.Println("  - Analise funções com baixa cobertura")
	fmt.Println("  - Teste casos de erro e edge cases")
	fmt.Println("  - Não ignore código não coberto sem análise")
	fmt.Println("  - Cobertura alta não garante código sem bugs")
	fmt.Println("  - Combine cobertura com outros tipos de teste")
	fmt.Println()
}
