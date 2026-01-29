package structs

import "fmt"

type Usuario struct {
	Nome  string
	Idade int
	Email string
	Ativo bool
}

func DemonstrarStructsBasicas() {
	fmt.Println("--- STRUCTS BÁSICAS ---")
	fmt.Println("Structs são uma coleção de campos nomeados que agrupam dados relacionados.\n")

	fmt.Println("Exemplo 1: Inicialização posicional")
	fmt.Println("Criando struct com valores na ordem dos campos:")
	usuarioExemplo := Usuario{"João", 20, "joao@gmail.com", true}
	fmt.Printf("usuarioExemplo: %+v\n", usuarioExemplo)
	fmt.Printf("Nome: %s, Idade: %d, Email: %s, Ativo: %t\n\n",
		usuarioExemplo.Nome, usuarioExemplo.Idade, usuarioExemplo.Email, usuarioExemplo.Ativo)

	fmt.Println("Exemplo 2: Inicialização com valores zero")
	fmt.Println("Criando struct vazia (todos os campos recebem valores zero):")
	var usuarioVazio Usuario
	fmt.Printf("usuarioVazio inicial: %+v\n", usuarioVazio)
	fmt.Printf("Nome (zero): '%s', Idade (zero): %d, Email (zero): '%s', Ativo (zero): %t\n\n",
		usuarioVazio.Nome, usuarioVazio.Idade, usuarioVazio.Email, usuarioVazio.Ativo)

	fmt.Println("Exemplo 3: Atribuição de campos individualmente")
	fmt.Println("Preenchendo campos da struct após criação:")
	usuarioVazio.Nome = "Cleber"
	usuarioVazio.Idade = 30
	usuarioVazio.Email = "cleber@gmail.com"
	usuarioVazio.Ativo = true
	fmt.Printf("usuarioVazio após preenchimento: %+v\n", usuarioVazio)
	fmt.Printf("Nome: %s, Idade: %d, Email: %s, Ativo: %t\n\n",
		usuarioVazio.Nome, usuarioVazio.Idade, usuarioVazio.Email, usuarioVazio.Ativo)

	fmt.Println("Exemplo 4: Inicialização com campos nomeados (parcial)")
	fmt.Println("Criando struct especificando apenas alguns campos:")
	usuarioParcial := Usuario{Nome: "Maria"}
	fmt.Printf("usuarioParcial (apenas nome definido): %+v\n", usuarioParcial)
	fmt.Printf("Nome: %s, Idade: %d, Email: '%s', Ativo: %t\n\n",
		usuarioParcial.Nome, usuarioParcial.Idade, usuarioParcial.Email, usuarioParcial.Ativo)
}
