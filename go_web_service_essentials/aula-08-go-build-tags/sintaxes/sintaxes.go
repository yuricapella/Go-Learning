package sintaxes

import "fmt"

func DemonstrarBuildTags() {
	fmt.Println("=== AULA 08: Go Build Tags ===")
	fmt.Println()
	fmt.Println("--- CONCEITO ---")
	fmt.Println("Build tags sao regras que dizem ao compilador quais arquivos entram ou nao no build.")
	fmt.Println("Em outras palavras: elas servem para montar versoes diferentes do mesmo programa a partir do mesmo projeto.")
	fmt.Println("Elas resolvem um problema de compilacao, nao de logica em tempo de execucao.")
	fmt.Println("Nao sao `if`s normais. Um `if` decide algo com o programa rodando; a build tag decide se um arquivo inteiro participa do programa antes mesmo de rodar.")
	fmt.Println()

	fmt.Println("--- Por que usar? ---")
	fmt.Println("Sem build tags, voce acaba levando para producao codigo que so queria usar em debug, profiling ou em um sistema operacional especifico.")
	fmt.Println("Com build tags, o binario final muda antes mesmo de rodar: certos arquivos entram e outros ficam de fora.")
	fmt.Println()

	Exemplo1TagExplicita()
	Exemplo2TagImplicitaPorSistemaOperacional()
	Exemplo3ComoONomeDaTagFunciona()
	Exemplo4OQueEPprof()
	Exemplo5PorQueGoRunSemTagsFalhaNoExemploDoVideo()
	Exemplo6ComoTestarNaPratica()
	Exemplo7QuandoFazSentidoUsar()
}

func Exemplo1TagExplicita() {
	fmt.Println("--- Exemplo 1: Tag explicita ---")
	fmt.Println("Use quando voce quer habilitar algo opcionalmente.")
	fmt.Println("No exemplo desta pasta, o arquivo `profile_enabled.go` so entra no build com `-tags profile`.")
	fmt.Println()
	fmt.Println("Sintaxe basica:")
	fmt.Println(SintaxeTagExplicita())
	fmt.Println()
}

func Exemplo2TagImplicitaPorSistemaOperacional() {
	fmt.Println("--- Exemplo 2: Tag implicita por sistema operacional ---")
	fmt.Println("Arquivos com sufixos como `_darwin.go`, `_linux.go` e `_windows.go` sao escolhidos automaticamente pelo Go.")
	fmt.Println("Voce nao passa `-tags` aqui: o proprio compilador decide com base em `GOOS` e `GOARCH`.")
	fmt.Println()
	fmt.Println("Sintaxe basica:")
	fmt.Println(SintaxeArquivoPorSistemaOperacional())
	fmt.Println()
}

func Exemplo3ComoONomeDaTagFunciona() {
	fmt.Println("--- Exemplo 3: Como o nome da tag funciona ---")
	fmt.Println("Se um arquivo tem `//go:build profile`, ele so entra quando voce compila com `-tags profile`.")
	fmt.Println("Se voce rodar `go run -tags oi .`, isso nao ativa `profile`.")
	fmt.Println("Para `-tags oi` funcionar, precisa existir arquivo com `//go:build oi`.")
	fmt.Println("Ou seja: o nome da tag e voce quem escolhe, e o comando precisa bater exatamente com esse nome.")
	fmt.Println()
}

func Exemplo4OQueEPprof() {
	fmt.Println("--- Exemplo 4: O que e `pprof` ---")
	fmt.Println("`pprof` e uma ferramenta de profiling do Go.")
	fmt.Println("Profiling significa medir como seu programa usa recursos, por exemplo:")
	fmt.Println("- CPU")
	fmt.Println("- memoria")
	fmt.Println("- goroutines")
	fmt.Println("- tempo gasto em funcoes")
	fmt.Println()
	fmt.Println("No caso de servidor HTTP, o import `_ \"net/http/pprof\"` registra endpoints como `/debug/pprof/`.")
	fmt.Println("Esses endpoints ajudam a investigar lentidao, consumo alto de memoria e gargalos.")
	fmt.Println("O motivo de esconder isso atras de uma build tag e simples: isso costuma ser util em diagnostico, mas nao e algo que voce quer expor por padrao em producao.")
	fmt.Println()
}

func Exemplo5PorQueGoRunSemTagsFalhaNoExemploDoVideo() {
	fmt.Println("--- Exemplo 5: Por que `go run .` nao ativa profiling? ---")
	fmt.Println("Porque o arquivo com `//go:build profile` fica fora da compilacao quando voce nao passa `-tags profile`.")
	fmt.Println("Entao o import de `net/http/pprof` simplesmente nao existe no binario padrao.")
	fmt.Println("Ou seja: nao e erro de runtime. E o compilador montando outro programa.")
	fmt.Println()
}

func Exemplo6ComoTestarNaPratica() {
	fmt.Println("--- Exemplo 6: Como testar nesta pasta ---")
	fmt.Println("Entre em `exampleapp` e rode:")
	fmt.Println("1. `go run .`")
	fmt.Println("2. `curl localhost:8080/feature-flags`")
	fmt.Println("3. `curl localhost:8080/by`")
	fmt.Println("4. `curl -i localhost:8080/debug/pprof/`")
	fmt.Println()
	fmt.Println("Depois rode:")
	fmt.Println("1. `go run -tags profile .`")
	fmt.Println("2. `curl localhost:8080/feature-flags`")
	fmt.Println("3. `curl -i localhost:8080/debug/pprof/`")
	fmt.Println()
	fmt.Println("Resultado esperado:")
	fmt.Println("- Sem tag: `/debug/pprof/` deve responder 404 e `profile_enabled=false`")
	fmt.Println("- Com tag: `/debug/pprof/` deve existir e `profile_enabled=true`")
	fmt.Println()
}

func Exemplo7QuandoFazSentidoUsar() {
	fmt.Println("--- Exemplo 7: Quando faz sentido usar ---")
	fmt.Println("- Debug, profiling ou observabilidade opcional")
	fmt.Println("- Implementacoes diferentes por sistema operacional")
	fmt.Println("- Integracoes nativas que so existem em certas arquiteturas")
	fmt.Println()
	fmt.Println("⚠️ IMPORTANTE:")
	fmt.Println("- Build tag nao substitui `if` normal quando a decisao e de negocio ou configuracao")
	fmt.Println("- Use build tag quando o codigo precisa entrar ou sair do binario")
	fmt.Println("- O maior valor aqui e seguranca, portabilidade e binarios mais limpos")
}

func SintaxeTagExplicita() string {
	return `//go:build profile

package main

import _ "net/http/pprof"`
}

func SintaxeArquivoPorSistemaOperacional() string {
	return "founder_darwin.go | founder_linux.go | founder_windows.go"
}
