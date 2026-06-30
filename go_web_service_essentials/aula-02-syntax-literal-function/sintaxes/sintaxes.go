package sintaxes

func ChamarFuncaoAnonima(n int) int {
	resultado := 0
	func() {
		resultado = n * 2
	}()
	return resultado
}

func GuardarFuncaoEmVariavel(prefixo string) func(string) string {
	return func(nome string) string {
		return prefixo + nome
	}
}

func ExecutarComDefer(eventos *[]string) {
	defer func() {
		*eventos = append(*eventos, "defer")
	}()

	*eventos = append(*eventos, "body")
}
