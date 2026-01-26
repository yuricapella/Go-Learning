package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Pacote 1.2 - Pacotes Externos")
	email := "teste@teste.com"
	err := checkmail.ValidateFormat(email)
	fmt.Println(err)
}
