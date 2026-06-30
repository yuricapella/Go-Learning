package didatica

import "fmt"

func DemonstrarVulnerabilidades() {
	fmt.Println("=== AULA 07: Identificar Vulnerabilidades ===")
	fmt.Println()

	fmt.Println("--- O que e identificar vulnerabilidade? ---")
	fmt.Println("E olhar para codigo, dependencias, configuracao e comportamento da aplicacao procurando formas de abuso.")
	fmt.Println("Nao e so procurar bug de compilacao. E perguntar: que entrada maliciosa, permissao errada ou exposicao de dado pode quebrar a seguranca?")
	fmt.Println()

	fmt.Println("--- Por que isso existe? ---")
	fmt.Println("Aplicacoes web recebem dados de fora o tempo todo.")
	fmt.Println("Se voce confia demais na entrada, expoe dados demais ou usa dependencias vulneraveis, o problema pode aparecer em producao.")
	fmt.Println()

	Exemplo1FontesDeInformacao()
	Exemplo2OWASPTop10()
	Exemplo3AgrupandoRiscos()
	Exemplo4ChecklistPratico()
	PontosImportantes()
}

func Exemplo1FontesDeInformacao() {
	fmt.Println("--- Exemplo 1: onde acompanhar vulnerabilidades ---")
	fmt.Println("- CVEs relacionados ao Go")
	fmt.Println("- bancos de vulnerabilidade de dependencias")
	fmt.Println("- anuncios oficiais em `golang-announce`")
	fmt.Println("- changelogs de bibliotecas que voce usa")
	fmt.Println()
}

func Exemplo2OWASPTop10() {
	fmt.Println("--- Exemplo 2: OWASP Top 10 ---")
	fmt.Println("OWASP Top 10 e uma lista de categorias comuns de falhas em aplicacoes web.")
	fmt.Println("Ela ajuda como checklist inicial para revisar entrada, saida, autenticacao, autorizacao e infraestrutura.")
	fmt.Println()
}

func Exemplo3AgrupandoRiscos() {
	fmt.Println("--- Exemplo 3: agrupando riscos ---")
	fmt.Println("Entrada: injection, XML external entities, insecure deserialization.")
	fmt.Println("Saida: XSS e exposicao de dados sensiveis.")
	fmt.Println("Acesso: autenticacao quebrada e controle de acesso quebrado.")
	fmt.Println("Infraestrutura: configuracao insegura, dependencias vulneraveis, logging e monitoramento insuficientes.")
	fmt.Println()
}

func Exemplo4ChecklistPratico() {
	fmt.Println("--- Exemplo 4: checklist pratico para olhar uma app Go ---")
	fmt.Println("1. Entradas externas sao validadas?")
	fmt.Println("2. SQL e comandos sao montados com parametros, nao concatenacao?")
	fmt.Println("3. Erros retornados ao cliente vazam detalhe interno?")
	fmt.Println("4. Rotas protegidas verificam autenticacao e autorizacao?")
	fmt.Println("5. Dependencias estao atualizadas?")
	fmt.Println("6. Logs ajudam a investigar problema sem vazar segredo?")
	fmt.Println()
}

func PontosImportantes() {
	fmt.Println("--- Pontos importantes ---")
	fmt.Println("- Seguranca nao e uma etapa unica; e um habito de revisao.")
	fmt.Println("- O OWASP Top 10 nao e codigo, e um mapa de riscos comuns.")
	fmt.Println("- Em Go, use a standard library a seu favor: parametros, escaping, limites de leitura e tratamento explicito de erros.")
	fmt.Println("- Nunca coloque segredo, token ou dado sensivel em resposta HTTP ou log sem necessidade.")
}
