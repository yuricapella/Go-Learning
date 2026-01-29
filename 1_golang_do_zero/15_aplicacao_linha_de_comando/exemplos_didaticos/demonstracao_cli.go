package demonstracao

import (
	"fmt"

	"github.com/urfave/cli"
	app "github.com/yuricapella/Go-Learning/1_golang_do_zero/15_aplicacao_linha_de_comando/app_aula"
)

// DemonstrarCriacaoAppCLI demonstra como criar uma aplicação CLI básica
func DemonstrarCriacaoAppCLI() {
	fmt.Println("--- CRIAÇÃO DE APLICAÇÃO CLI ---")
	fmt.Println("Aplicações CLI (Command Line Interface) permitem interagir com programas através do terminal.\n")

	fmt.Println("Exemplo 1: Criando uma aplicação básica")
	app := cli.NewApp()
	app.Name = "Minha Aplicação"
	app.Usage = "Uma aplicação de exemplo"
	app.Version = "1.0.0"

	fmt.Printf("  Nome: %s\n", app.Name)
	fmt.Printf("  Uso: %s\n", app.Usage)
	fmt.Printf("  Versão: %s\n", app.Version)
	fmt.Println()

	fmt.Println("Exemplo 2: Adicionando flags (opções)")
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "localhost",
			Usage: "Endereço do host",
		},
		cli.IntFlag{
			Name:  "porta",
			Value: 8080,
			Usage: "Porta do servidor",
		},
	}
	app.Flags = flags

	fmt.Println("  Flags adicionadas:")
	for _, flag := range flags {
		fmt.Printf("    - %s: %s\n", flag.GetName(), flag.String())
	}
	fmt.Println()

	fmt.Println("⚠️  IMPORTANTE:")
	fmt.Println("  - Flags são opções globais disponíveis para todos os comandos")
	fmt.Println("  - Flags de comandos são específicas para cada comando")
	fmt.Println("  - Use contexto.String(\"flag\") para obter valores de flags")
	fmt.Println()
}

// DemonstrarComandosCLI demonstra como criar e usar comandos em uma aplicação CLI
func DemonstrarComandosCLI() {
	fmt.Println("--- COMANDOS CLI ---")
	fmt.Println("Comandos são ações específicas que a aplicação pode executar.\n")

	fmt.Println("Exemplo 1: Estrutura de um comando")
	fmt.Println("  cli.Command{")
	fmt.Println("      Name:   \"nome-do-comando\",")
	fmt.Println("      Usage:  \"Descrição do comando\",")
	fmt.Println("      Flags:  []cli.Flag{...},")
	fmt.Println("      Action: funcaoQueExecuta,")
	fmt.Println("  }")
	fmt.Println()

	fmt.Println("Exemplo 2: Criando comandos")
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:  "iniciar",
			Usage: "Inicia o servidor",
			Action: func(c *cli.Context) error {
				fmt.Println("    Servidor iniciado!")
				return nil
			},
		},
		{
			Name:  "parar",
			Usage: "Para o servidor",
			Action: func(c *cli.Context) error {
				fmt.Println("    Servidor parado!")
				return nil
			},
		},
	}

	fmt.Printf("  Comandos criados: %d\n", len(app.Commands))
	for _, cmd := range app.Commands {
		fmt.Printf("    - %s: %s\n", cmd.Name, cmd.Usage)
	}
	fmt.Println()

	fmt.Println("Exemplo 3: Comandos com flags")
	comandoComFlag := cli.Command{
		Name:  "conectar",
		Usage: "Conecta a um host",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "host",
				Value: "localhost",
			},
		},
		Action: func(c *cli.Context) error {
			host := c.String("host")
			fmt.Printf("    Conectando a: %s\n", host)
			return nil
		},
	}

	fmt.Printf("  Comando: %s\n", comandoComFlag.Name)
	fmt.Println("  Flags do comando:")
	for _, flag := range comandoComFlag.Flags {
		fmt.Printf("    - %s\n", flag.GetName())
	}
	fmt.Println()

	fmt.Println("⚠️  IMPORTANTE:")
	fmt.Println("  - Comandos são executados com: app comando [flags]")
	fmt.Println("  - Exemplo: go run main.go ip --host google.com")
	fmt.Println("  - Flags de comandos são específicas para cada comando")
	fmt.Println("  - Use contexto.String(\"flag\") dentro da Action para obter valores")
	fmt.Println()
}

// DemonstrarExecucaoApp demonstra como executar uma aplicação CLI
func DemonstrarExecucaoApp() {
	fmt.Println("--- EXECUÇÃO DE APLICAÇÃO CLI ---")
	fmt.Println("Para executar uma aplicação CLI, você precisa passar os argumentos do sistema operacional.\n")

	fmt.Println("Exemplo: Executando a aplicação")
	fmt.Println("  app := Gerar()")
	fmt.Println("  if erro := app.Run(os.Args); erro != nil {")
	fmt.Println("      log.Fatal(erro)")
	fmt.Println("  }")
	fmt.Println()

	fmt.Println("Explicação:")
	fmt.Println("  - os.Args contém os argumentos passados na linha de comando")
	fmt.Println("  - app.Run() processa esses argumentos e executa o comando apropriado")
	fmt.Println("  - Se houver erro, ele é retornado e deve ser tratado")
	fmt.Println()

	fmt.Println("Exemplos de uso no terminal:")
	fmt.Println("  # Ver ajuda")
	fmt.Println("  go run main.go --help")
	fmt.Println()
	fmt.Println("  # Executar comando ip")
	fmt.Println("  go run main.go ip --host google.com")
	fmt.Println()
	fmt.Println("  # Executar comando servidores")
	fmt.Println("  go run main.go servidores --host google.com")
	fmt.Println()

	fmt.Println("⚠️  IMPORTANTE:")
	fmt.Println("  - Sempre trate erros retornados por app.Run()")
	fmt.Println("  - os.Args[0] é o nome do programa")
	fmt.Println("  - os.Args[1:] são os argumentos passados pelo usuário")
	fmt.Println()
}

// DemonstrarAppCompleta demonstra a aplicação completa funcionando
func DemonstrarAppCompleta() {
	fmt.Println("--- APLICAÇÃO CLI COMPLETA ---")
	fmt.Println("Demonstração da aplicação completa de busca de IPs e servidores.\n")

	fmt.Println("A aplicação criada possui:")
	fmt.Println("  - Nome: Aplicação de Linha de Comando")
	fmt.Println("  - Uso: Busca IPs e Nomes de Servidores na Internet")
	fmt.Println("  - Comandos disponíveis:")
	fmt.Println("    1. ip - Busca IPs de endereços na internet")
	fmt.Println("    2. servidores - Busca nomes de servidores na internet")
	fmt.Println()

	fmt.Println("Como usar:")
	fmt.Println("  1. Compile a aplicação:")
	fmt.Println("     go build -o app main.go")
	fmt.Println()
	fmt.Println("  2. Execute os comandos:")
	fmt.Println("     ./app ip --host google.com")
	fmt.Println("     ./app servidores --host google.com")
	fmt.Println()

	fmt.Println("Durante o desenvolvimento:")
	fmt.Println("  go run main.go ip --host google.com")
	fmt.Println("  go run main.go servidores --host google.com")
	fmt.Println()

	// Criar e mostrar a estrutura da aplicação
	aplicacao := app.Gerar()
	fmt.Println("Estrutura da aplicação criada:")
	fmt.Printf("  Nome: %s\n", aplicacao.Name)
	fmt.Printf("  Uso: %s\n", aplicacao.Usage)
	fmt.Printf("  Comandos: %d\n", len(aplicacao.Commands))
	for _, cmd := range aplicacao.Commands {
		fmt.Printf("    - %s: %s\n", cmd.Name, cmd.Usage)
	}
	fmt.Println()

	fmt.Println("⚠️  NOTA:")
	fmt.Println("  Esta é uma demonstração. Para executar realmente, use:")
	fmt.Println("  app := Gerar()")
	fmt.Println("  app.Run(os.Args)")
	fmt.Println()
}
