package main

type usuarioSintaxe struct {
	nome  string
	idade int
	email string
	ativo bool
}

func sintaxeStructInicializacaoPosicional() {
	usuario1 := usuarioSintaxe{"João", 20, "joao@gmail.com", true}
	_ = usuario1
}

func sintaxeStructValoresZero() {
	var usuario2 usuarioSintaxe
	usuario2.nome = "Cleber"
	usuario2.idade = 30
	usuario2.email = "cleber@gmail.com"
	usuario2.ativo = true
}

func sintaxeStructCamposNomeados() {
	usuario3 := usuarioSintaxe{nome: "Maria"}
	_ = usuario3
}

func sintaxeStructsAninhadas() {
	type enderecoSintaxe struct {
		rua    string
		numero int
		cidade string
		estado string
		cep    string
	}

	type PessoaSintaxe struct {
		usuario  usuarioSintaxe
		endereco enderecoSintaxe
	}

	pessoa := PessoaSintaxe{
		usuario: usuarioSintaxe{
			nome:  "João",
			idade: 20,
			email: "joao@gmail.com",
			ativo: true,
		},
		endereco: enderecoSintaxe{
			rua:    "Rua 1",
			numero: 123,
			cidade: "São Paulo",
			estado: "SP",
			cep:    "1234567890",
		},
	}

	_ = pessoa.usuario.nome
	_ = pessoa.endereco.rua
}
