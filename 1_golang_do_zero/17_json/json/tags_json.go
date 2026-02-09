package json

import (
	"encoding/json"
	"fmt"
)

// DemonstrarTagsJSON demonstra o uso de tags JSON para customização
func DemonstrarTagsJSON() {
	fmt.Println("--- TAGS JSON ---")
	fmt.Println("Tags JSON permitem customizar como campos são serializados/deserializados.")
	fmt.Println("Elas controlam nomes dos campos, omissão de valores zero, etc.")
	fmt.Println()

	fmt.Println("Por que usar tags JSON?")
	fmt.Println("  Sem tags: Nomes dos campos na struct devem ser exatamente como no JSON.")
	fmt.Println("  Isso força nomes em inglês ou quebra quando APIs mudam nomes.")
	fmt.Println()
	fmt.Println("  Com tags: Você pode usar nomes em português na struct e mapear para")
	fmt.Println("  nomes diferentes no JSON. Também permite omitir campos vazios.")
	fmt.Println()

	fmt.Println("Sintaxe básica:")
	fmt.Println("  type Struct struct {")
	fmt.Println("      Campo string `json:\"nome_no_json\"`")
	fmt.Println("  }")
	fmt.Println()

	ExemploTagNomeCustomizado()
	ExemploTagOmitempty()
	ExemploTagIgnorarCampo()

	fmt.Println("⚠️  IMPORTANTE:")
	fmt.Println("  - Tag json:\"-\" ignora o campo completamente")
	fmt.Println("  - Tag json:\"nome,omitempty\" omite campo se for zero value")
	fmt.Println("  - Tags são case-sensitive")
	fmt.Println("  - Use tags para compatibilidade com APIs externas")
	fmt.Println("  - Tags funcionam tanto em Marshal quanto Unmarshal")
	fmt.Println()
}

func ExemploTagNomeCustomizado() {
	fmt.Println("Exemplo 1: Tag para nome customizado")
	fmt.Println("Mapeando campo em português para JSON em inglês:")
	fmt.Println()

	type Usuario struct {
		NomeCompleto string `json:"full_name"`
		Idade        int    `json:"age"`
		Email        string `json:"email"`
	}

	usuario := Usuario{
		NomeCompleto: "Pedro Silva",
		Idade:        28,
		Email:        "pedro@example.com",
	}

	jsonBytes, _ := json.Marshal(usuario)
	fmt.Printf("  Struct: %+v\n", usuario)
	fmt.Printf("  JSON: %s\n", string(jsonBytes))
	fmt.Println()
}

func ExemploTagOmitempty() {
	fmt.Println("Exemplo 2: Tag omitempty")
	fmt.Println("Omitindo campos com zero value do JSON:")
	fmt.Println()

	type Config struct {
		Nome      string `json:"nome"`
		Descricao string `json:"descricao,omitempty"`
		Ativo     bool   `json:"ativo,omitempty"`
	}

	config := Config{
		Nome: "Sistema",
		// Descricao e Ativo são zero value
	}

	jsonBytes, _ := json.Marshal(config)
	fmt.Printf("  Struct: %+v\n", config)
	fmt.Printf("  JSON (sem campos vazios): %s\n", string(jsonBytes))
	fmt.Println()
}

func ExemploTagIgnorarCampo() {
	fmt.Println("Exemplo 3: Tag para ignorar campo")
	fmt.Println("Campo não aparece no JSON:")
	fmt.Println()

	type Usuario struct {
		Nome  string `json:"nome"`
		Senha string `json:"-"` // Ignorado no JSON
		Email string `json:"email"`
	}

	usuario := Usuario{
		Nome:  "Ana",
		Senha: "senha123",
		Email: "ana@example.com",
	}

	jsonBytes, _ := json.Marshal(usuario)
	fmt.Printf("  Struct: %+v\n", usuario)
	fmt.Printf("  JSON (sem senha): %s\n", string(jsonBytes))
	fmt.Println()
}
