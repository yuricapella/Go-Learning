package json

import (
	"encoding/json"
	"fmt"
)

// DemonstrarMarshalIndent demonstra JSON formatado para leitura humana
func DemonstrarMarshalIndent() {
	fmt.Println("--- MARSHAL INDENT (JSON Formatado) ---")
	fmt.Println("MarshalIndent gera JSON com indentação para facilitar leitura.")
	fmt.Println("Útil para logs, debug e arquivos de configuração.")
	fmt.Println()

	fmt.Println("Por que usar MarshalIndent?")
	fmt.Println("  Sem indentação: JSON fica em uma linha só, difícil de ler e debugar.")
	fmt.Println()
	fmt.Println("  Com indentação: JSON formatado com quebras de linha e espaços.")
	fmt.Println("  Facilita leitura humana e depuração, mas ocupa mais espaço.")
	fmt.Println()

	fmt.Println("Função:")
	fmt.Println("  json.MarshalIndent(estrutura, prefixo, indentacao) retorna ([]byte, error)")
	fmt.Println()

	ExemploMarshalIndent()
	ExemploMarshalIndentCustomizado()

	fmt.Println("⚠️  IMPORTANTE:")
	fmt.Println("  - MarshalIndent gera JSON legível mas ocupa mais espaço")
	fmt.Println("  - Use para debug, logs e arquivos de configuração")
	fmt.Println("  - Para APIs HTTP, prefira Marshal() (menor tamanho)")
	fmt.Println("  - Prefixo é adicionado no início de cada linha")
	fmt.Println("  - Indentação controla espaçamento (geralmente \"  \" ou \"\\t\")")
	fmt.Println()
}

func ExemploMarshalIndent() {
	fmt.Println("Exemplo 1: MarshalIndent")
	fmt.Println("JSON formatado com indentação padrão:")
	fmt.Println()

	type Endereco struct {
		Rua    string
		Cidade string
		CEP    string
	}

	type Usuario struct {
		Nome     string
		Idade    int
		Endereco Endereco
	}

	usuario := Usuario{
		Nome:  "Carlos",
		Idade: 35,
		Endereco: Endereco{
			Rua:    "Av. Principal",
			Cidade: "São Paulo",
			CEP:    "01234-567",
		},
	}

	jsonBytes, _ := json.MarshalIndent(usuario, "", "  ")
	fmt.Println("  JSON formatado:")
	fmt.Printf("  %s\n", string(jsonBytes))
	fmt.Println()
}

func ExemploMarshalIndentCustomizado() {
	fmt.Println("Exemplo 2: MarshalIndent com prefixo customizado")
	fmt.Println("Usando prefixo para comentários ou formatação especial:")
	fmt.Println()

	type Config struct {
		Host string
		Port int
		SSL  bool
	}

	config := Config{
		Host: "api.example.com",
		Port: 443,
		SSL:  true,
	}

	jsonBytes, _ := json.MarshalIndent(config, "// ", "  ")
	fmt.Println("  JSON com prefixo:")
	fmt.Printf("  %s\n", string(jsonBytes))
	fmt.Println()
}
