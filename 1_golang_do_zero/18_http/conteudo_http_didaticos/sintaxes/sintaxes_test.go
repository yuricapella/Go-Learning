package sintaxeHttp_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestSintaxeHandlerFunc testa criação de handler básico
func TestSintaxeHandlerFunc(t *testing.T) {
	requisicao := httptest.NewRequest(http.MethodGet, "/rota", nil)
	gravadorResposta := httptest.NewRecorder()

	handler := func(escritorResposta http.ResponseWriter, requisicao *http.Request) {
		escritorResposta.Write([]byte("Resposta da rota"))
	}

	handler(gravadorResposta, requisicao)

	if gravadorResposta.Code != http.StatusOK {
		t.Errorf("Status code recebido: %d, Esperado: %d", gravadorResposta.Code, http.StatusOK)
	}

	corpoEsperado := "Resposta da rota"
	if gravadorResposta.Body.String() != corpoEsperado {
		t.Errorf("Corpo recebido: %s, Esperado: %s", gravadorResposta.Body.String(), corpoEsperado)
	}
}

// TestSintaxeMetodoGET testa handler para método GET
func TestSintaxeMetodoGET(t *testing.T) {
	requisicao := httptest.NewRequest(http.MethodGet, "/dados", nil)
	gravadorResposta := httptest.NewRecorder()

	handler := func(escritorResposta http.ResponseWriter, requisicao *http.Request) {
		if requisicao.Method != http.MethodGet {
			http.Error(escritorResposta, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}
		escritorResposta.Write([]byte("Dados retornados"))
	}

	handler(gravadorResposta, requisicao)

	if gravadorResposta.Code != http.StatusOK {
		t.Errorf("Status code recebido: %d, Esperado: %d", gravadorResposta.Code, http.StatusOK)
	}

	if !strings.Contains(gravadorResposta.Body.String(), "Dados retornados") {
		t.Errorf("Corpo não contém texto esperado: %s", gravadorResposta.Body.String())
	}
}

// TestSintaxeMetodoPOST testa handler para método POST
func TestSintaxeMetodoPOST(t *testing.T) {
	corpoRequisicao := bytes.NewBufferString("Dados de teste")
	requisicao := httptest.NewRequest(http.MethodPost, "/criar", corpoRequisicao)
	gravadorResposta := httptest.NewRecorder()

	handler := func(escritorResposta http.ResponseWriter, requisicao *http.Request) {
		if requisicao.Method != http.MethodPost {
			http.Error(escritorResposta, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}
		escritorResposta.Write([]byte("Recebido: Dados de teste"))
	}

	handler(gravadorResposta, requisicao)

	if gravadorResposta.Code != http.StatusOK {
		t.Errorf("Status code recebido: %d, Esperado: %d", gravadorResposta.Code, http.StatusOK)
	}

	if !strings.Contains(gravadorResposta.Body.String(), "Recebido") {
		t.Errorf("Corpo não contém texto esperado: %s", gravadorResposta.Body.String())
	}
}

// TestSintaxeQueryParams testa leitura de query parameters
func TestSintaxeQueryParams(t *testing.T) {
	requisicao := httptest.NewRequest(http.MethodGet, "/buscar?nome=João&idade=30", nil)
	gravadorResposta := httptest.NewRecorder()

	handler := func(escritorResposta http.ResponseWriter, requisicao *http.Request) {
		nome := requisicao.URL.Query().Get("nome")
		idade := requisicao.URL.Query().Get("idade")
		escritorResposta.Write([]byte("Nome: " + nome + ", Idade: " + idade))
	}

	handler(gravadorResposta, requisicao)

	if !strings.Contains(gravadorResposta.Body.String(), "João") {
		t.Errorf("Corpo não contém nome esperado: %s", gravadorResposta.Body.String())
	}

	if !strings.Contains(gravadorResposta.Body.String(), "30") {
		t.Errorf("Corpo não contém idade esperada: %s", gravadorResposta.Body.String())
	}
}

// TestSintaxeRespostaJSON testa retorno de JSON
func TestSintaxeRespostaJSON(t *testing.T) {
	requisicao := httptest.NewRequest(http.MethodGet, "/json", nil)
	gravadorResposta := httptest.NewRecorder()

	handler := func(escritorResposta http.ResponseWriter, requisicao *http.Request) {
		dados := map[string]string{
			"mensagem": "Olá",
			"status":   "sucesso",
		}
		escritorResposta.Header().Set("Content-Type", "application/json")
		json.NewEncoder(escritorResposta).Encode(dados)
	}

	handler(gravadorResposta, requisicao)

	if gravadorResposta.Header().Get("Content-Type") != "application/json" {
		t.Errorf("Content-Type recebido: %s, Esperado: application/json", gravadorResposta.Header().Get("Content-Type"))
	}

	var dadosRecebidos map[string]string
	erro := json.NewDecoder(gravadorResposta.Body).Decode(&dadosRecebidos)
	if erro != nil {
		t.Fatalf("Erro ao decodificar JSON: %v", erro)
	}

	if dadosRecebidos["mensagem"] != "Olá" {
		t.Errorf("Mensagem recebida: %s, Esperado: %s", dadosRecebidos["mensagem"], "Olá")
	}
}

// TestSintaxeStatusCode testa definição de status code
func TestSintaxeStatusCode(t *testing.T) {
	requisicao := httptest.NewRequest(http.MethodGet, "/erro", nil)
	gravadorResposta := httptest.NewRecorder()

	handler := func(escritorResposta http.ResponseWriter, requisicao *http.Request) {
		escritorResposta.WriteHeader(http.StatusNotFound)
		escritorResposta.Write([]byte("Página não encontrada"))
	}

	handler(gravadorResposta, requisicao)

	if gravadorResposta.Code != http.StatusNotFound {
		t.Errorf("Status code recebido: %d, Esperado: %d", gravadorResposta.Code, http.StatusNotFound)
	}
}

// TestSintaxeLerBodyJSON testa leitura de body JSON
func TestSintaxeLerBodyJSON(t *testing.T) {
	dadosJSON := `{"nome":"João","idade":30}`
	corpoRequisicao := bytes.NewBufferString(dadosJSON)
	requisicao := httptest.NewRequest(http.MethodPost, "/receber", corpoRequisicao)
	gravadorResposta := httptest.NewRecorder()

	handler := func(escritorResposta http.ResponseWriter, requisicao *http.Request) {
		var dados map[string]interface{}
		json.NewDecoder(requisicao.Body).Decode(&dados)
		defer requisicao.Body.Close()

		if dados["nome"] != "João" {
			t.Errorf("Nome recebido: %v, Esperado: João", dados["nome"])
		}
	}

	handler(gravadorResposta, requisicao)
}

// TestSintaxeServidorComMux testa uso de ServeMux
func TestSintaxeServidorComMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(escritorResposta http.ResponseWriter, requisicao *http.Request) {
		escritorResposta.Write([]byte("Rota raiz"))
	})
	mux.HandleFunc("/outra", func(escritorResposta http.ResponseWriter, requisicao *http.Request) {
		escritorResposta.Write([]byte("Outra rota"))
	})

	requisicao := httptest.NewRequest(http.MethodGet, "/outra", nil)
	gravadorResposta := httptest.NewRecorder()

	mux.ServeHTTP(gravadorResposta, requisicao)

	if !strings.Contains(gravadorResposta.Body.String(), "Outra rota") {
		t.Errorf("Corpo não contém texto esperado: %s", gravadorResposta.Body.String())
	}
}
