package didatica

import "fmt"

func DemonstrarHTTPHandler() {
	fmt.Println("=== AULA 06: HTTP Handler Function ===")
	fmt.Println()

	fmt.Println("--- O que e um handler HTTP? ---")
	fmt.Println("E uma funcao que recebe `http.ResponseWriter` e `*http.Request`.")
	fmt.Println("O `ResponseWriter` representa a resposta que voce vai escrever.")
	fmt.Println("O `Request` representa a requisicao recebida: metodo, path, headers, body e outros dados.")
	fmt.Println()

	Exemplo1HealthCheck()
	Exemplo2Routing()
	Exemplo3RouterPadrao()
	Exemplo4ListenAndServe()
	PontosImportantes()
}

func Exemplo1HealthCheck() {
	fmt.Println("--- Exemplo 1: health check ---")
	fmt.Println("Um health handler responde algo simples, como `OK`, para sistemas de monitoramento.")
	fmt.Println("Em uma aplicacao real, ele poderia tambem verificar banco, cache ou dependencia externa.")
	fmt.Println()
}

func Exemplo2Routing() {
	fmt.Println("--- Exemplo 2: routing ---")
	fmt.Println("Routing e o mapa entre caminho e handler.")
	fmt.Println("Exemplo: quando alguem acessa `/health`, o servidor chama `HealthHandler`.")
	fmt.Println()
}

func Exemplo3RouterPadrao() {
	fmt.Println("--- Exemplo 3: router padrao do Go ---")
	fmt.Println("O `http.ServeMux` e simples e suficiente para muitas APIs pequenas.")
	fmt.Println("Ele nao tenta resolver todos os casos de frameworks maiores.")
	fmt.Println("Quando precisar de rotas muito sofisticadas, voce pode avaliar um router externo.")
	fmt.Println()
}

func Exemplo4ListenAndServe() {
	fmt.Println("--- Exemplo 4: ListenAndServe ---")
	fmt.Println("`http.ListenAndServe(\":8080\", mux)` inicia o servidor na porta 8080 usando o mux informado.")
	fmt.Println("Se o endereco comeca com `:`, o servidor escuta em todas as interfaces disponiveis.")
	fmt.Println()
}

func PontosImportantes() {
	fmt.Println("--- Pontos importantes ---")
	fmt.Println("- Handler HTTP tem a assinatura `func(w http.ResponseWriter, r *http.Request)`.")
	fmt.Println("- Escreva a resposta usando `w`.")
	fmt.Println("- Leia dados da requisicao usando `r`.")
	fmt.Println("- Use `http.NewServeMux()` para criar um roteador simples e explicito.")
}
