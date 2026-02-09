package main

func sintaxePassagemPorValor() {
	variavel1 := 10
	variavel2 := variavel1
	variavel1++
	_ = variavel2
}

func sintaxeOperadorEndereco() {
	variavel1 := 10
	variavel3 := &variavel1
	_ = variavel3
}

func sintaxeOperadorDesreferenciacao() {
	variavel1 := 10
	variavel3 := &variavel1
	valor := *variavel3
	_ = valor
}

func sintaxeModificacaoPonteiro() {
	variavel1 := 10
	variavel3 := &variavel1
	variavel1++
	*variavel3 = 100
}

func sintaxeFuncaoComPonteiro(ponteiro *int) {
	*ponteiro++
}

func sintaxeUsarFuncaoComPonteiro() {
	valor := 5
	sintaxeFuncaoComPonteiro(&valor)
}
