package main

import "fmt"

type usuarioSintaxe struct {
	nome  string
	idade int
	email string
	ativo bool
}

func sintaxeStructInicializacaoPosicional() {
	usuarioExemplo := usuarioSintaxe{"João", 20, "joao@gmail.com", true}
	fmt.Println(usuarioExemplo)
}

func sintaxeStructValoresZero() {
	var usuarioExemplo usuarioSintaxe
	usuarioExemplo.nome = "Cleber"
	usuarioExemplo.idade = 30
	usuarioExemplo.email = "cleber@gmail.com"
	usuarioExemplo.ativo = true
	fmt.Println(usuarioExemplo)
}

func sintaxeStructCamposNomeados() {
	usuarioExemplo := usuarioSintaxe{nome: "Maria"}
	fmt.Println(usuarioExemplo)
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

	pessoaExemplo := PessoaSintaxe{
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

	fmt.Println(pessoaExemplo.usuario.nome)
	fmt.Println(pessoaExemplo.endereco.rua)
}

type UsuarioMetodosSintaxe struct {
	Nome  string
	Idade int
	Email string
	Ativo bool
}

func (usuario UsuarioMetodosSintaxe) ExibirNome() {
	fmt.Println(usuario.Nome)
}

func sintaxeMetodoReceiverValor() {
	usuarioExemplo := UsuarioMetodosSintaxe{Nome: "João", Idade: 25}
	usuarioExemplo.ExibirNome()
}

func (usuario *UsuarioMetodosSintaxe) AtualizarNome(novoNome string) {
	usuario.Nome = novoNome
}

func sintaxeMetodoReceiverPonteiro() {
	usuarioExemplo := UsuarioMetodosSintaxe{Nome: "João", Email: "joao@gmail.com"}
	usuarioExemplo.AtualizarNome("João Silva")
	fmt.Println(usuarioExemplo.Nome)
}

func (usuario UsuarioMetodosSintaxe) ObterNomeCompleto() string {
	return fmt.Sprintf("%s (%d anos)", usuario.Nome, usuario.Idade)
}

func sintaxeMetodoRetornarValor() {
	usuarioExemplo := UsuarioMetodosSintaxe{Nome: "João", Idade: 25}
	nomeCompleto := usuarioExemplo.ObterNomeCompleto()
	fmt.Println(nomeCompleto)
}

func (usuario *UsuarioMetodosSintaxe) AtivarUsuario() {
	usuario.Ativo = true
}

func sintaxeMetodoModificarStruct() {
	usuarioExemplo := UsuarioMetodosSintaxe{Nome: "João", Ativo: false}
	usuarioExemplo.AtivarUsuario()
	fmt.Println(usuarioExemplo.Ativo)
}
