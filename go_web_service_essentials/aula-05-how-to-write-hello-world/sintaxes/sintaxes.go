package sintaxes

import "fmt"

func HelloWorld() {
	fmt.Println("Hello, Go")
}

func MontarMensagem(nome string) string {
	return fmt.Sprintf("Hello, %s", nome)
}
