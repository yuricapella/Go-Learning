package didatica

import "fmt"

func DemonstrarTestesHTTPHandler() {
	fmt.Println("=== AULA 09: Testing HTTP Handler ===")
	fmt.Println()

	fmt.Println("--- O que e testar um handler HTTP? ---")
	fmt.Println("Um handler HTTP em Go e uma funcao que recebe `http.ResponseWriter` e `*http.Request`.")
	fmt.Println("Testar um handler significa chamar essa funcao diretamente, criando esses dois objetos em memoria.")
	fmt.Println("Assim voce valida status code, headers e body sem subir um servidor real.")
	fmt.Println()

	fmt.Println("--- Por que isso existe? ---")
	fmt.Println("Subir servidor, abrir porta e fazer curl e util em teste manual, mas e pesado para teste automatizado.")
	fmt.Println("O pacote `net/http/httptest` cria uma request falsa e um recorder de resposta para testar o handler de forma rapida.")
	fmt.Println()

	Exemplo1ObjetosNecessarios()
	Exemplo2FluxoDoTeste()
	Exemplo3OQueValidar()
	Exemplo4ComoRodar()
	PontosImportantes()
}

func Exemplo1ObjetosNecessarios() {
	fmt.Println("--- Exemplo 1: objetos necessarios ---")
	fmt.Println("Para testar um handler voce precisa de:")
	fmt.Println("- `httptest.NewRequest`: cria uma request parecida com a real")
	fmt.Println("- `httptest.NewRecorder`: captura o que o handler escreveria na resposta")
	fmt.Println("- chamada direta do handler: `MeuHandler(w, r)`")
	fmt.Println()
}

func Exemplo2FluxoDoTeste() {
	fmt.Println("--- Exemplo 2: fluxo do teste ---")
	fmt.Println("1. Monte o body da request com `strings.NewReader`.")
	fmt.Println("2. Crie a request com metodo, path e body.")
	fmt.Println("3. Crie o recorder.")
	fmt.Println("4. Chame o handler diretamente.")
	fmt.Println("5. Pegue a resposta com `w.Result()`.")
	fmt.Println("6. Compare status code e decode o JSON retornado.")
	fmt.Println()
}

func Exemplo3OQueValidar() {
	fmt.Println("--- Exemplo 3: o que validar ---")
	fmt.Println("Valide comportamento observavel, nao detalhe interno.")
	fmt.Println("Bons alvos: status code, content type, body JSON, mensagem de erro e regra de metodo HTTP.")
	fmt.Println("Evite testar se o handler chamou uma funcao interna especifica quando a saida ja prova o comportamento.")
	fmt.Println()
}

func Exemplo4ComoRodar() {
	fmt.Println("--- Exemplo 4: como rodar nesta aula ---")
	fmt.Println("Use `go run .` para ver esta explicacao.")
	fmt.Println("Use `go test ./...` para rodar o teste real do handler em `exampleapp`.")
	fmt.Println()
}

func PontosImportantes() {
	fmt.Println("--- Pontos importantes ---")
	fmt.Println("- `httptest` testa o handler sem rede, sem porta e sem servidor real.")
	fmt.Println("- `ResponseRecorder` guarda status, headers e body escritos pelo handler.")
	fmt.Println("- `http.StatusOK` e mais legivel que usar o numero `200` diretamente.")
	fmt.Println("- Feche `res.Body` nos testes quando usar `w.Result()`.")
}
