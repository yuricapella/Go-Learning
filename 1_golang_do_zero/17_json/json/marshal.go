package json

import (
	"encoding/json"
	"fmt"
)

// DemonstrarMarshal demonstra como converter estruturas Go para JSON
func DemonstrarMarshal() {
	fmt.Println("--- MARSHAL (Go → JSON) ---")
	fmt.Println("Marshal converte estruturas Go em formato JSON.")
	fmt.Println("É usado para serializar dados antes de enviar via HTTP, salvar em arquivo, etc.")
	fmt.Println()

	fmt.Println("Por que usar Marshal?")
	fmt.Println("  Sem Marshal: Você precisaria construir strings JSON manualmente,")
	fmt.Println("  o que é propenso a erros (aspas faltando, vírgulas erradas, etc.).")
	fmt.Println()
	fmt.Println("  Com Marshal: Conversão automática e segura. O pacote encoding/json")
	fmt.Println("  garante que o JSON gerado seja válido e bem formatado.")
	fmt.Println()

	fmt.Println("Função principal:")
	fmt.Println("  json.Marshal(estrutura) retorna ([]byte, error)")
	fmt.Println()

	ExemploMarshal()
	ExemploMarshalTiposDiferentes()

	fmt.Println("⚠️  IMPORTANTE:")
	fmt.Println("  - Marshal retorna []byte, não string (use string() para converter)")
	fmt.Println("  - Campos não exportados (minúsculos) são ignorados no JSON")
	fmt.Println("  - Tipos não suportados causam erro (ex: channels, functions)")
	fmt.Println("  - Use tags JSON para customizar nomes dos campos")
	fmt.Println("  - Marshal sempre retorna JSON válido quando bem-sucedido")
	fmt.Println("  - Use json.MarshalIndent() para JSON formatado (legível)")
	fmt.Println()
}

func ExemploMarshal() {
	fmt.Println("Exemplo 1: Marshal")
	fmt.Println("Convertendo uma struct simples para JSON:")
	fmt.Println()

	type Usuario struct {
		Nome  string
		Idade int
		Email string
	}

	usuario := Usuario{
		Nome:  "João",
		Idade: 30,
		Email: "joao@example.com",
	}

	jsonBytes, erro := json.Marshal(usuario)
	if erro != nil {
		fmt.Printf("  Erro ao fazer marshal: %v\n", erro)
		return
	}

	fmt.Printf("  Struct Go: %+v\n", usuario)
	fmt.Printf("  JSON gerado: %s\n", string(jsonBytes))
	fmt.Println()
}

func ExemploMarshalTiposDiferentes() {
	fmt.Println("Exemplo 2: Marshal com diferentes tipos")
	fmt.Println("JSON suporta strings, números, booleanos, arrays e objetos:")
	fmt.Println()

	type Produto struct {
		Nome       string
		Preco      float64
		Disponivel bool
		Tags       []string
	}

	produto := Produto{
		Nome:       "Notebook",
		Preco:      2999.99,
		Disponivel: true,
		Tags:       []string{"eletrônicos", "computadores"},
	}

	jsonBytes, _ := json.Marshal(produto)
	fmt.Printf("  JSON: %s\n", string(jsonBytes))
	fmt.Println()
}
