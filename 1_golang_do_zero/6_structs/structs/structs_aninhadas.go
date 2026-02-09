package structs

import "fmt"

type Endereco struct {
	Rua    string
	Numero int
	Cidade string
	Estado string
	Cep    string
}

type Pessoa struct {
	Usuario  Usuario
	Endereco Endereco
}

func DemonstrarStructsAninhadas() {
	fmt.Println("--- STRUCTS ANINHADAS (COMPOSIÇÃO) ---")
	fmt.Println("Podemos ter structs dentro de structs:\n")

	pessoaExemplo := Pessoa{
		Usuario: Usuario{
			Nome:  "João",
			Idade: 20,
			Email: "joao@gmail.com",
			Ativo: true,
		},
		Endereco: Endereco{
			Rua:    "Rua 1",
			Numero: 123,
			Cidade: "São Paulo",
			Estado: "SP",
			Cep:    "1234567890",
		},
	}

	fmt.Printf("Pessoa completa: %+v\n\n", pessoaExemplo)
	fmt.Println("Acessando campos aninhados:")
	fmt.Printf("Nome da pessoa: %s\n", pessoaExemplo.Usuario.Nome)
	fmt.Printf("Endereço completo: %s, %d - %s/%s - CEP: %s\n",
		pessoaExemplo.Endereco.Rua, pessoaExemplo.Endereco.Numero,
		pessoaExemplo.Endereco.Cidade, pessoaExemplo.Endereco.Estado, pessoaExemplo.Endereco.Cep)
	fmt.Println()
}
