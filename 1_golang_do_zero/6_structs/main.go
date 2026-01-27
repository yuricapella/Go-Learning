package main

import "fmt"

// ============================================
// DEFINIÇÃO DE STRUCT
// ============================================
// Structs são uma coleção de campos nomeados
// que agrupam dados relacionados
type usuario struct {
	nome  string
	idade int
	email string
	ativo bool
}

func main() {
	// ============================================
	// INICIALIZAÇÃO POSICIONAL
	// ============================================
	fmt.Println("--- INICIALIZAÇÃO POSICIONAL ---")
	fmt.Println("Criando struct com valores na ordem dos campos:")
	usuario1 := usuario{"João", 20, "joao@gmail.com", true}
	fmt.Printf("usuario1: %+v\n", usuario1)
	fmt.Printf("Nome: %s, Idade: %d, Email: %s, Ativo: %t\n\n",
		usuario1.nome, usuario1.idade, usuario1.email, usuario1.ativo)

	// ============================================
	// INICIALIZAÇÃO COM VALORES ZERO
	// ============================================
	fmt.Println("--- INICIALIZAÇÃO COM VALORES ZERO ---")
	fmt.Println("Criando struct vazia (todos os campos recebem valores zero):")
	var usuario2 usuario
	fmt.Printf("usuario2 inicial: %+v\n", usuario2)
	fmt.Printf("Nome (zero): '%s', Idade (zero): %d, Email (zero): '%s', Ativo (zero): %t\n\n",
		usuario2.nome, usuario2.idade, usuario2.email, usuario2.ativo)

	// ============================================
	// ATRIBUIÇÃO DE CAMPOS INDIVIDUALMENTE
	// ============================================
	fmt.Println("--- ATRIBUIÇÃO DE CAMPOS INDIVIDUALMENTE ---")
	fmt.Println("Preenchendo campos da struct após criação:")
	usuario2.nome = "Cleber"
	usuario2.idade = 30
	usuario2.email = "cleber@gmail.com"
	usuario2.ativo = true
	fmt.Printf("usuario2 após preenchimento: %+v\n", usuario2)
	fmt.Printf("Nome: %s, Idade: %d, Email: %s, Ativo: %t\n\n",
		usuario2.nome, usuario2.idade, usuario2.email, usuario2.ativo)

	// ============================================
	// INICIALIZAÇÃO COM CAMPOS NOMEADOS (PARCIAL)
	// ============================================
	fmt.Println("--- INICIALIZAÇÃO COM CAMPOS NOMEADOS (PARCIAL) ---")
	fmt.Println("Criando struct especificando apenas alguns campos:")
	usuario3 := usuario{nome: "Maria"}
	fmt.Printf("usuario3 (apenas nome definido): %+v\n", usuario3)
	fmt.Printf("Nome: %s, Idade: %d, Email: '%s', Ativo: %t\n\n",
		usuario3.nome, usuario3.idade, usuario3.email, usuario3.ativo)

	// ============================================
	// STRUCTS ANINHADAS (COMPOSIÇÃO)
	// ============================================
	fmt.Println("--- STRUCTS ANINHADAS (COMPOSIÇÃO) ---")
	fmt.Println("Podemos ter structs dentro de structs:")

	type endereco struct {
		rua    string
		numero int
		cidade string
		estado string
		cep    string
	}

	type Pessoa struct {
		usuario  usuario
		endereco endereco
	}

	pessoa := Pessoa{
		usuario: usuario{
			nome:  "João",
			idade: 20,
			email: "joao@gmail.com",
			ativo: true,
		},
		endereco: endereco{
			rua:    "Rua 1",
			numero: 123,
			cidade: "São Paulo",
			estado: "SP",
			cep:    "1234567890",
		},
	}

	fmt.Printf("Pessoa completa: %+v\n\n", pessoa)
	fmt.Println("Acessando campos aninhados:")
	fmt.Printf("Nome da pessoa: %s\n", pessoa.usuario.nome)
	fmt.Printf("Endereço completo: %s, %d - %s/%s - CEP: %s\n",
		pessoa.endereco.rua, pessoa.endereco.numero,
		pessoa.endereco.cidade, pessoa.endereco.estado, pessoa.endereco.cep)
}
