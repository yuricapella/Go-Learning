package sintaxeJson_test

import (
	"encoding/json"
	"strings"
	"testing"
)

// TestSintaxeMarshal testa marshal
func TestSintaxeMarshal(t *testing.T) {
	type Usuario struct {
		Nome  string
		Idade int
	}

	usuario := Usuario{Nome: "João", Idade: 30}
	jsonBytes, erro := json.Marshal(usuario)

	if erro != nil {
		t.Fatalf("Erro ao fazer marshal: %v", erro)
	}

	esperado := `{"Nome":"João","Idade":30}`
	if string(jsonBytes) != esperado {
		t.Errorf("JSON recebido: %s, Esperado: %s", string(jsonBytes), esperado)
	}
}

// TestSintaxeUnmarshal testa unmarshal
func TestSintaxeUnmarshal(t *testing.T) {
	type Usuario struct {
		Nome  string
		Idade int
	}

	jsonString := `{"Nome":"Maria","Idade":25}`
	var usuario Usuario

	erro := json.Unmarshal([]byte(jsonString), &usuario)
	if erro != nil {
		t.Fatalf("Erro ao fazer unmarshal: %v", erro)
	}

	if usuario.Nome != "Maria" {
		t.Errorf("Nome recebido: %s, Esperado: Maria", usuario.Nome)
	}

	if usuario.Idade != 25 {
		t.Errorf("Idade recebida: %d, Esperado: 25", usuario.Idade)
	}
}

// TestSintaxeMarshalIndent testa marshal com indentação
func TestSintaxeMarshalIndent(t *testing.T) {
	type Usuario struct {
		Nome string
	}

	usuario := Usuario{Nome: "João"}
	jsonBytes, erro := json.MarshalIndent(usuario, "", "  ")

	if erro != nil {
		t.Fatalf("Erro ao fazer marshal indent: %v", erro)
	}

	// Verifica se contém o nome
	if !strings.Contains(string(jsonBytes), "João") {
		t.Errorf("JSON não contém o nome esperado: %s", string(jsonBytes))
	}
}

// TestSintaxeTagJSON testa tags JSON
func TestSintaxeTagJSON(t *testing.T) {
	type Usuario struct {
		Nome  string `json:"nome"`
		Idade int    `json:"idade,omitempty"`
		Senha string `json:"-"`
	}

	usuario := Usuario{Nome: "João", Senha: "senha123"}
	jsonBytes, erro := json.Marshal(usuario)

	if erro != nil {
		t.Fatalf("Erro ao fazer marshal: %v", erro)
	}

	// Verifica se contém "nome" mas não "Senha"
	jsonString := string(jsonBytes)
	if !strings.Contains(jsonString, "nome") {
		t.Errorf("JSON não contém campo 'nome': %s", jsonString)
	}

	if strings.Contains(jsonString, "Senha") {
		t.Errorf("JSON não deveria conter campo 'Senha': %s", jsonString)
	}
}

// TestSintaxeTagOmitempty testa tag omitempty
func TestSintaxeTagOmitempty(t *testing.T) {
	type Config struct {
		Nome      string `json:"nome"`
		Descricao string `json:"descricao,omitempty"`
	}

	config := Config{Nome: "Sistema"}
	jsonBytes, _ := json.Marshal(config)
	jsonString := string(jsonBytes)

	// Verifica se não contém "descricao" (zero value omitido)
	if strings.Contains(jsonString, "descricao") {
		t.Errorf("JSON não deveria conter 'descricao' quando omitempty: %s", jsonString)
	}
}
