package main

func sintaxeOperadoresAritmeticos() {
	numero1 := 10
	numero2 := 20

	soma := numero1 + numero2
	subtracao := numero1 - numero2
	multiplicacao := numero1 * numero2
	divisao := numero1 / numero2
	resto := numero1 % numero2

	_ = soma
	_ = subtracao
	_ = multiplicacao
	_ = divisao
	_ = resto
}

func sintaxeOperadoresAtribuicaoComposta() {
	valor := 10
	valor += 5
	valor -= 3
	valor *= 2
	valor /= 4
	valor %= 3
}

func sintaxeOperadoresUnarios() {
	num := 10
	num++
	num--

	positivo := +5
	negativo := -5

	_ = positivo
	_ = negativo
}

func sintaxeOperadoresComparacao() {
	a := 10
	b := 20

	igual := a == b
	diferente := a != b
	maior := a > b
	menor := a < b
	maiorOuIgual := a >= b
	menorOuIgual := a <= b

	_ = igual
	_ = diferente
	_ = maior
	_ = menor
	_ = maiorOuIgual
	_ = menorOuIgual
}

func sintaxeOperadoresLogicos() {
	x := 10
	y := 20

	e := x == y && x < y
	ou := x == y || x < y

	booleano := false
	negacao := !booleano

	_ = e
	_ = ou
	_ = negacao
}

func sintaxeConversaoTipos() {
	var numero3 int16 = 10
	var numero4 int32 = 20

	var soma2 int32 = int32(numero3) + numero4
	_ = soma2
}
