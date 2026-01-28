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
		fmt.Println("Função f")
	}

	f()

	var f2 = func(txt string) string {
		return txt
	}

	resultado2 := f2("Função f2")
	fmt.Println(resultado2)

	retorno := func(texto string) string {
		return fmt.Sprintf("recebido -> %s", texto)
	}("Passando parametro")
	fmt.Println(retorno)
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
		fmt.Println(erro)
	}
	fmt.Println(resultado)
}

func sintaxeFuncaoVariadica(numeros ...int) int {
	soma := 0
	for _, num := range numeros {
		soma += num
	}
	return soma
}

func sintaxeUsarFuncaoVariadica() {
	resultado1 := sintaxeFuncaoVariadica(1, 2, 3)
	fmt.Println(resultado1)

	resultado2 := sintaxeFuncaoVariadica(10, 20, 30, 40, 50)
	fmt.Println(resultado2)

	numeros := []int{5, 10, 15}
	resultado3 := sintaxeFuncaoVariadica(numeros...)
	fmt.Println(resultado3)
}

func sintaxeFuncaoRecursivaFatorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * sintaxeFuncaoRecursivaFatorial(n-1)
}

func sintaxeFuncaoRecursivaFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return sintaxeFuncaoRecursivaFibonacci(n-1) + sintaxeFuncaoRecursivaFibonacci(n-2)
}
