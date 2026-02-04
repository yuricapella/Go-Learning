package json

import (
	"encoding/json"
	"fmt"
)

// DemonstrarUnmarshal demonstra como converter JSON para estruturas Go
func DemonstrarUnmarshal() {
	fmt.Println("--- UNMARSHAL (JSON → Go) ---")
	fmt.Println("Unmarshal converte dados JSON em estruturas Go.")
	fmt.Println("É usado para deserializar dados recebidos via HTTP, ler de arquivos, etc.")
	fmt.Println()

	fmt.Println("Por que usar Unmarshal?")
	fmt.Println("  Sem Unmarshal: Você precisaria fazer parsing manual do JSON,")
	fmt.Println("  extraindo valores campo por campo, o que é trabalhoso e propenso a erros.")
	fmt.Println()
	fmt.Println("  Com Unmarshal: Conversão automática e type-safe. O pacote encoding/json")
	fmt.Println("  mapeia campos JSON para campos da struct automaticamente.")
	fmt.Println()

	fmt.Println("Função principal:")
	fmt.Println("  json.Unmarshal([]byte, &estrutura) retorna error")
	fmt.Println()

	ExemploUnmarshal()
	ExemploUnmarshalCamposParciais()

	fmt.Println("⚠️  IMPORTANTE:")
	fmt.Println("  - Sempre passe um ponteiro para Unmarshal (&estrutura)")
	fmt.Println("  - Campos JSON que não existem na struct são ignorados")
	fmt.Println("  - Campos da struct que não existem no JSON recebem zero value")
	fmt.Println("  - Tipos incompatíveis causam erro (ex: string em campo int)")
	fmt.Println("  - Use tags JSON para mapear nomes diferentes")
	fmt.Println("  - Unmarshal não retorna erro se JSON tiver campos extras")
	fmt.Println()
}

func ExemploUnmarshal() {
	fmt.Println("Exemplo 1: Unmarshal")
	fmt.Println("Convertendo JSON para struct:")
	fmt.Println()

	type Usuario struct {
		Nome  string
		Idade int
		Email string
	}

	jsonString := `{"Nome":"Maria","Idade":25,"Email":"maria@example.com"}`
	jsonBytes := []byte(jsonString)

	var usuario Usuario
	erro := json.Unmarshal(jsonBytes, &usuario)
	if erro != nil {
		fmt.Printf("  Erro ao fazer unmarshal: %v\n", erro)
		return
	}

	fmt.Printf("  JSON recebido: %s\n", jsonString)
	fmt.Printf("  Struct Go: %+v\n", usuario)
	fmt.Println()
}

func ExemploUnmarshalCamposParciais() {
	fmt.Println("Exemplo 2: Unmarshal com campos parciais")
	fmt.Println("Campos faltantes recebem zero value:")
	fmt.Println()

	type Produto struct {
		Nome       string
		Preco      float64
		Disponivel bool
	}

	jsonParcial := `{"Nome":"Mouse"}`
	var produto Produto

	json.Unmarshal([]byte(jsonParcial), &produto)
	fmt.Printf("  JSON: %s\n", jsonParcial)
	fmt.Printf("  Produto: %+v\n", produto)
	fmt.Println("  Preco e Disponivel receberam zero value")
	fmt.Println()
}
