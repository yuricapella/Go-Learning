package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// app struct principal para podermos fazer o aplicativo
// Gerar vai retornar a aplicação de linha de comando para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidores na Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereços na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca nomes de servidores na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}

// utilizando o comando no terminal go run main.go ip --host site, ele retorna os ips desse site
// go run main.go ip --host google.com
func buscarIps(contexto *cli.Context) {
	host := contexto.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

// utilizando o comando no terminal go run main.go servidores --host site, ele retorna os servidores desse site
// go run main.go servidores --host google.com
func buscarServidores(contexto *cli.Context) {
	host := contexto.String("host")

	servidores, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
