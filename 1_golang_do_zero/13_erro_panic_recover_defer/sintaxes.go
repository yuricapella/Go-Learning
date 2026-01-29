package main

import (
	"errors"
	"fmt"
)

func sintaxeDefer() {
	defer fmt.Println("Defer 1")
	defer fmt.Println("Defer 2")
	defer fmt.Println("Defer 3")
	fmt.Println("Código normal")
}

func sintaxeDeferLimpeza() {
	fmt.Println("Abrindo recurso")
	defer fmt.Println("Fechando recurso")
	fmt.Println("Usando recurso")
}

func sintaxeCriarErro() {
	erro1 := errors.New("erro de exemplo")
	fmt.Println(erro1)

	erro2 := fmt.Errorf("erro formatado: %d", 10)
	fmt.Println(erro2)
}

func sintaxeRetornarErro(valor int) (int, error) {
	if valor < 0 {
		return 0, fmt.Errorf("valor negativo não permitido: %d", valor)
	}
	return valor * 2, nil
}

func sintaxeTratarErro() {
	resultado, erro := sintaxeRetornarErro(5)
	if erro != nil {
		fmt.Println(erro)
	} else {
		fmt.Println(resultado)
	}

	resultado2, erro2 := sintaxeRetornarErro(-5)
	if erro2 != nil {
		fmt.Println(erro2)
	} else {
		fmt.Println(resultado2)
	}
}

func sintaxePanic() {
	panic("erro fatal")
}

func sintaxePanicComValor() {
	panic(fmt.Errorf("erro: %s", "divisão por zero"))
}

func sintaxeRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	panic("erro recuperável")
}

func sintaxeRecoverComRetorno() (resultado int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			resultado = -1
		}
	}()

	panic("erro")
}

func sintaxeDeferPanicRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recuperado: %v\n", r)
		}
	}()

	fmt.Println("Código normal")
	panic("erro")
}

func sintaxeFuncaoComRecover(a, b int) (resultado int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			resultado = 0
		}
	}()

	if b == 0 {
		panic("divisão por zero")
	}

	resultado = a / b
	return resultado
}

func sintaxeUsarFuncaoComRecover() {
	resultado1 := sintaxeFuncaoComRecover(10, 0)
	fmt.Println(resultado1)

	resultado2 := sintaxeFuncaoComRecover(10, 2)
	fmt.Println(resultado2)
}
