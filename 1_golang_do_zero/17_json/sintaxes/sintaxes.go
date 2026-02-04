package sintaxeJson

import (
	"encoding/json"
)

// SintaxeMarshal demonstra marshal
func SintaxeMarshal() {
	type Usuario struct {
		Nome  string
		Idade int
	}

	usuario := Usuario{Nome: "João", Idade: 30}
	jsonBytes, erro := json.Marshal(usuario)
	if erro != nil {
		// Tratar erro
		_ = erro
		return
	}
	_ = jsonBytes
}

// SintaxeUnmarshal demonstra unmarshal
func SintaxeUnmarshal() {
	type Usuario struct {
		Nome  string
		Idade int
	}

	jsonString := `{"Nome":"Maria","Idade":25}`
	var usuario Usuario
	erro := json.Unmarshal([]byte(jsonString), &usuario)
	if erro != nil {
		// Tratar erro
		_ = erro
		return
	}
	_ = usuario
}

// SintaxeMarshalIndent demonstra marshal com indentação
func SintaxeMarshalIndent() {
	type Usuario struct {
		Nome string
	}

	usuario := Usuario{Nome: "João"}
	jsonBytes, _ := json.MarshalIndent(usuario, "", "  ")
	_ = jsonBytes
}

// SintaxeTagJSON demonstra uso de tags JSON
func SintaxeTagJSON() {
	type Usuario struct {
		Nome  string `json:"nome"`
		Idade int    `json:"idade,omitempty"`
		Senha string `json:"-"`
	}

	usuario := Usuario{Nome: "João"}
	jsonBytes, _ := json.Marshal(usuario)
	_ = jsonBytes
}
