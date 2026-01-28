package main

type pessoaSintaxe struct {
	nome      string
	sobrenome string
	idade     uint
	altura    uint
}

type estudanteSintaxe struct {
	pessoaSintaxe
	curso     string
	faculdade string
}

type estudanteComNomeacaoSintaxe struct {
	pessoa pessoaSintaxe
	curso  string
}

type enderecoSintaxe struct {
	rua string
}

type trabalhosintaxe struct {
	rua string
}

type funcionarioSintaxe struct {
	pessoaSintaxe
	enderecoSintaxe
	trabalhosintaxe
	cargo string
}

func sintaxeEmbeddingAnonimo() {
	pessoa1 := pessoaSintaxe{
		nome:      "João",
		sobrenome: "Silva",
		idade:     20,
		altura:    180,
	}

	estudante1 := estudanteSintaxe{
		pessoaSintaxe: pessoa1,
		curso:         "Engenharia",
	}

	_ = estudante1.nome
	_ = estudante1.pessoaSintaxe.nome
}

func sintaxeEmbeddingComNomeacao() {
	estudante2 := estudanteComNomeacaoSintaxe{
		pessoa: pessoaSintaxe{
			nome: "Pedro",
		},
		curso: "Medicina",
	}

	_ = estudante2.pessoa.nome
}

func sintaxeEmbeddingAmbiguidade() {
	funcionario1 := funcionarioSintaxe{
		pessoaSintaxe: pessoaSintaxe{
			nome: "Ana",
		},
		enderecoSintaxe: enderecoSintaxe{
			rua: "Rua das Flores",
		},
		trabalhosintaxe: trabalhosintaxe{
			rua: "Avenida do Trabalho",
		},
		cargo: "Desenvolvedora",
	}

	_ = funcionario1.nome
	_ = funcionario1.enderecoSintaxe.rua
	_ = funcionario1.trabalhosintaxe.rua
}
