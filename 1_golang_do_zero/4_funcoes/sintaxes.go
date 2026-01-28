package main

import "fmt"

func sintaxeFuncaoSimples(a int, b int) int {
	return a + b
}

func sintaxeFuncaoMultiplosRetornos(n1, n2 int) (int, int, int, int) {
	soma := n1 + n2
	subtracao := n1 - n2
	multiplicacao := n1 * n2
	divisao := n1 / n2
	return soma, subtracao, multiplicacao, divisao
}

func sintaxeFuncaoAnonima() {
	var f = func() {
	}

	f()

	var f2 = func(txt string) string {
		return txt
	}

	resultado2 := f2("Função f2")
	_ = resultado2
}

func sintaxeIgnorarRetornos() {
	soma2, _, _, _ := sintaxeFuncaoMultiplosRetornos(10, 20)
	_ = soma2
}

func sintaxeRetornoComErro(nome string, idade int) (string, int, error) {
	return nome, idade, nil
}

func sintaxeUsarRetornoComErro() {
	nome, idade, erro := sintaxeRetornoComErro("João", 20)
	if erro != nil {
		_ = erro
	}
	_ = nome
	_ = idade
}

func sintaxeRetornoNomeado(a, b int) (resultado int, erro error) {
	if b == 0 {
		erro = fmt.Errorf("divisão por zero não permitida")
		return
	}
	resultado = a / b
	return
}

func sintaxeUsarRetornoNomeado() {
	resultado, erro := sintaxeRetornoNomeado(10, 5)
	if erro != nil {
		fmt.Println("Erro:", erro)
	}
	fmt.Println("Resultado:", resultado)
}
