package main

func sintaxeVariaveis() {
	var variavel1 string = "Variavel 1"
	variavel2 := "Variavel 2"
	_ = variavel1
	_ = variavel2

	var (
		variavel3 string = "Variavel 3"
		variavel4 string = "Variavel 4"
	)
	_ = variavel3
	_ = variavel4

	variavel5, variavel6 := "Variavel 5", "Variavel 6"

	const constante1 = "Constante 1"
	_ = constante1

	variavel5, variavel6 = variavel6, variavel5
	_ = variavel5
	_ = variavel6
}
