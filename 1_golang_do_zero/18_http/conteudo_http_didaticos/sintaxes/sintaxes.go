package sintaxeHttp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// SintaxeServidorHTTPBasico demonstra criação básica de servidor HTTP
func SintaxeServidorHTTPBasico() {
	handler := func(escritorResposta http.ResponseWriter, requisicao *http.Request) {
		fmt.Fprintf(escritorResposta, "Olá, mundo!")
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

// SintaxeHandlerFunc demonstra uso de HandleFunc
func SintaxeHandlerFunc() {
	http.HandleFunc("/rota", func(escritorResposta http.ResponseWriter, requisicao *http.Request) {
		fmt.Fprintf(escritorResposta, "Resposta da rota")
	})
}

// SintaxeMetodoGET demonstra handler para método GET
func SintaxeMetodoGET() {
	http.HandleFunc("/dados", func(escritorResposta http.ResponseWriter, requisicao *http.Request) {
		if requisicao.Method != http.MethodGet {
			http.Error(escritorResposta, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}
		fmt.Fprintf(escritorResposta, "Dados retornados")
	})
}

// SintaxeMetodoPOST demonstra handler para método POST
func SintaxeMetodoPOST() {
	http.HandleFunc("/criar", func(escritorResposta http.ResponseWriter, requisicao *http.Request) {
		if requisicao.Method != http.MethodPost {
			http.Error(escritorResposta, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}
		corpoRequisicao, _ := io.ReadAll(requisicao.Body)
		defer requisicao.Body.Close()
		fmt.Fprintf(escritorResposta, "Recebido: %s", string(corpoRequisicao))
	})
}

// SintaxeQueryParams demonstra leitura de query parameters
func SintaxeQueryParams() {
	http.HandleFunc("/buscar", func(escritorResposta http.ResponseWriter, requisicao *http.Request) {
		nome := requisicao.URL.Query().Get("nome")
		idade := requisicao.URL.Query().Get("idade")
		fmt.Fprintf(escritorResposta, "Nome: %s, Idade: %s", nome, idade)
	})
}

// SintaxeRespostaJSON demonstra retorno de JSON
func SintaxeRespostaJSON() {
	http.HandleFunc("/json", func(escritorResposta http.ResponseWriter, requisicao *http.Request) {
		dados := map[string]string{
			"mensagem": "Olá",
			"status":   "sucesso",
		}
		escritorResposta.Header().Set("Content-Type", "application/json")
		json.NewEncoder(escritorResposta).Encode(dados)
	})
}

// SintaxeStatusCode demonstra definição de status code
func SintaxeStatusCode() {
	http.HandleFunc("/erro", func(escritorResposta http.ResponseWriter, requisicao *http.Request) {
		escritorResposta.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(escritorResposta, "Página não encontrada")
	})
}

// SintaxeLerBodyJSON demonstra leitura de body JSON
func SintaxeLerBodyJSON() {
	http.HandleFunc("/receber", func(escritorResposta http.ResponseWriter, requisicao *http.Request) {
		var dados map[string]interface{}
		json.NewDecoder(requisicao.Body).Decode(&dados)
		defer requisicao.Body.Close()
		fmt.Fprintf(escritorResposta, "Recebido: %v", dados)
	})
}

// SintaxeServidorComPorta demonstra servidor em porta específica
func SintaxeServidorComPorta() {
	porta := ":8080"
	http.ListenAndServe(porta, nil)
}

// SintaxeServidorComMux demonstra uso de ServeMux
func SintaxeServidorComMux() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(escritorResposta http.ResponseWriter, requisicao *http.Request) {
		fmt.Fprintf(escritorResposta, "Rota raiz")
	})
	mux.HandleFunc("/outra", func(escritorResposta http.ResponseWriter, requisicao *http.Request) {
		fmt.Fprintf(escritorResposta, "Outra rota")
	})
	http.ListenAndServe(":8080", mux)
}
