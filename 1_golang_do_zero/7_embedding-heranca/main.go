package main

import "fmt"

// ============================================
// STRUCT BASE
// ============================================
type pessoa struct {
	nome      string
	sobrenome string
	idade     uint
	altura    uint
}

// ============================================
// EMBEDDING (HERANÇA EM GO)
// ============================================
// Estudante "herda" de pessoa usando embedding
// Como estamos usando apenas o tipo (sem nomear),
// os atributos de pessoa ficam diretamente acessíveis em estudante
type estudante struct {
	pessoa    // Embedding anônimo - campos acessíveis diretamente
	curso     string
	faculdade string
}

// ============================================
// EMBEDDING COM NOMEAÇÃO EXPLÍCITA
// ============================================
// Se nomeássemos o campo pessoa, teríamos que acessar via estudante.pessoa.nome
type estudanteComNomeacao struct {
	pessoa pessoa // Campo nomeado - acesso via estudante.pessoa.nome
	curso  string
}

// ============================================
// EXEMPLO DE AMBIGUIDADE
// ============================================
type endereco struct {
	rua string
}

type trabalho struct {
	rua string // Campo com mesmo nome que endereco.rua
}

type funcionario struct {
	pessoa
	endereco
	trabalho
	cargo string
}

func main() {
	// ============================================
	// CRIAÇÃO DE PESSOA
	// ============================================
	fmt.Println("--- CRIAÇÃO DE PESSOA ---")
	pessoa1 := pessoa{
		nome:      "João",
		sobrenome: "Silva",
		idade:     20,
		altura:    180,
	}
	fmt.Printf("pessoa1: %+v\n\n", pessoa1)

	// ============================================
	// EMBEDDING ANÔNIMO - ACESSO DIRETO
	// ============================================
	fmt.Println("--- EMBEDDING ANÔNIMO (ACESSO DIRETO) ---")
	fmt.Println("Como 'estudante' usa apenas o tipo 'pessoa' (sem nomear),")
	fmt.Println("os campos de pessoa ficam diretamente acessíveis:\n")

	estudante1 := estudante{
		pessoa: pessoa{
			nome:      "Maria",
			sobrenome: "Santos",
			idade:     22,
			altura:    165,
		},
		curso:     "Engenharia",
		faculdade: "Universidade Federal do Rio de Janeiro",
	}

	fmt.Printf("estudante1 completo: %+v\n\n", estudante1)
	fmt.Println("Acesso direto aos campos de pessoa:")
	fmt.Printf("estudante1.nome = %s\n", estudante1.nome)
	fmt.Printf("estudante1.sobrenome = %s\n", estudante1.sobrenome)
	fmt.Printf("estudante1.idade = %d\n", estudante1.idade)
	fmt.Printf("estudante1.altura = %d\n", estudante1.altura)
	fmt.Printf("estudante1.curso = %s\n", estudante1.curso)
	fmt.Printf("estudante1.faculdade = %s\n\n", estudante1.faculdade)

	// Também podemos acessar via pessoa (mas não é necessário)
	fmt.Println("Também podemos acessar via estudante1.pessoa.nome:")
	fmt.Printf("estudante1.pessoa.nome = %s\n\n", estudante1.pessoa.nome)

	// ============================================
	// EMBEDDING COM NOMEAÇÃO EXPLÍCITA
	// ============================================
	fmt.Println("--- EMBEDDING COM NOMEAÇÃO EXPLÍCITA ---")
	fmt.Println("Se nomearmos o campo pessoa, precisamos acessar via estudante.pessoa.nome:\n")

	estudante2 := estudanteComNomeacao{
		pessoa: pessoa{
			nome:      "Pedro",
			sobrenome: "Oliveira",
			idade:     25,
			altura:    175,
		},
		curso: "Medicina",
	}

	fmt.Printf("estudante2 completo: %+v\n\n", estudante2)
	fmt.Println("Acesso com nomeação explícita:")
	fmt.Printf("estudante2.pessoa.nome = %s\n", estudante2.pessoa.nome)
	fmt.Printf("estudante2.pessoa.idade = %d\n", estudante2.pessoa.idade)
	fmt.Printf("estudante2.curso = %s\n\n", estudante2.curso)

	// Tentar acessar diretamente causaria erro de compilação:
	// fmt.Println(estudante2.nome) // ERRO: estudante2.nome undefined

	// ============================================
	// AMBIGUIDADE COM CAMPOS DE MESMO NOME
	// ============================================
	fmt.Println("--- AMBIGUIDADE COM CAMPOS DE MESMO NOME ---")
	fmt.Println("Quando duas structs embedded têm campos com o mesmo nome,")
	fmt.Println("Go exige que especifiquemos qual struct usar:\n")

	funcionario1 := funcionario{
		pessoa: pessoa{
			nome:      "Ana",
			sobrenome: "Costa",
			idade:     30,
			altura:    160,
		},
		endereco: endereco{
			rua: "Rua das Flores, 123",
		},
		trabalho: trabalho{
			rua: "Avenida do Trabalho, 456",
		},
		cargo: "Desenvolvedora",
	}

	fmt.Printf("funcionario1 completo: %+v\n\n", funcionario1)
	fmt.Println("Acesso aos campos únicos (sem ambiguidade):")
	fmt.Printf("funcionario1.nome = %s\n", funcionario1.nome)
	fmt.Printf("funcionario1.cargo = %s\n\n", funcionario1.cargo)

	fmt.Println("Acesso aos campos com mesmo nome (precisa especificar):")
	fmt.Printf("funcionario1.endereco.rua = %s\n", funcionario1.endereco.rua)
	fmt.Printf("funcionario1.trabalho.rua = %s\n\n", funcionario1.trabalho.rua)

	// Tentar acessar diretamente causaria erro de compilação:
	// fmt.Println(funcionario1.rua) // ERRO: ambiguous selector funcionario1.rua
	fmt.Println("⚠️  Se tentássemos acessar 'funcionario1.rua' diretamente,")
	fmt.Println("   Go retornaria erro: 'ambiguous selector funcionario1.rua'")
	fmt.Println("   Por isso precisamos especificar: funcionario1.endereco.rua")
}
