package structs

import "fmt"

func DemonstrarMetodos() {
	fmt.Println("--- MÉTODOS ---")
	fmt.Println("Métodos são funções associadas a um tipo (geralmente structs).")
	fmt.Println("Eles têm um receiver que especifica o tipo ao qual o método pertence.\n")

	fmt.Println("Exemplo 1: Método com receiver por valor")
	fmt.Println("Receiver por valor cria uma cópia da struct.")
	fmt.Println("Mudanças feitas no método não afetam a struct original.\n")

	usuarioExemplo := Usuario{
		Nome:  "João",
		Idade: 25,
		Email: "joao@gmail.com",
		Ativo: true,
	}

	fmt.Println("Chamando método com receiver por valor:")
	usuarioExemplo.ExibirInformacoes()
	fmt.Println()

	fmt.Println("Exemplo 2: Método com receiver por ponteiro")
	fmt.Println("Receiver por ponteiro trabalha com referência à struct original.")
	fmt.Println("Mudanças feitas no método afetam a struct original.\n")

	fmt.Println("Antes de modificar:")
	fmt.Printf("Nome: %s, Email: %s\n", usuarioExemplo.Nome, usuarioExemplo.Email)

	usuarioExemplo.AtualizarNome("João Silva")
	usuarioExemplo.AtualizarEmail("joao.silva@gmail.com")

	fmt.Println("Depois de modificar:")
	fmt.Printf("Nome: %s, Email: %s\n", usuarioExemplo.Nome, usuarioExemplo.Email)
	fmt.Println()

	fmt.Println("Exemplo 3: Método que retorna valor calculado")
	nomeCompleto := usuarioExemplo.ObterNomeCompleto()
	fmt.Printf("Nome completo: %s\n", nomeCompleto)
	fmt.Println()

	fmt.Println("Exemplo 4: Método para verificar propriedades")
	maioridade := usuarioExemplo.VerificarMaioridade()
	fmt.Printf("É maior de idade? %t\n", maioridade)
	fmt.Println()

	fmt.Println("⚠️  IMPORTANTE:")
	fmt.Println("  - Receiver por valor: use quando o método apenas lê dados")
	fmt.Println("  - Receiver por ponteiro: use quando o método modifica dados")
	fmt.Println("  - Receiver por ponteiro é mais eficiente para structs grandes")
	fmt.Println()
}

// Método com receiver por valor - apenas leitura
func (usuario Usuario) ExibirInformacoes() {
	fmt.Printf("Nome: %s\n", usuario.Nome)
	fmt.Printf("Idade: %d\n", usuario.Idade)
	fmt.Printf("Email: %s\n", usuario.Email)
	fmt.Printf("Ativo: %t\n", usuario.Ativo)
}

// Método com receiver por valor - retorna valor calculado
func (usuario Usuario) ObterNomeCompleto() string {
	return fmt.Sprintf("%s (%d anos)", usuario.Nome, usuario.Idade)
}

// Método com receiver por valor - verifica propriedade
func (usuario Usuario) VerificarMaioridade() bool {
	return usuario.Idade >= 18
}

// Método com receiver por ponteiro - modifica a struct
func (usuario *Usuario) AtualizarNome(novoNome string) {
	usuario.Nome = novoNome
}

// Método com receiver por ponteiro - modifica a struct
func (usuario *Usuario) AtualizarEmail(novoEmail string) {
	usuario.Email = novoEmail
}

// Método com receiver por ponteiro - modifica status
func (usuario *Usuario) AtivarUsuario() {
	usuario.Ativo = true
}

// Método com receiver por ponteiro - modifica status
func (usuario *Usuario) DesativarUsuario() {
	usuario.Ativo = false
}
