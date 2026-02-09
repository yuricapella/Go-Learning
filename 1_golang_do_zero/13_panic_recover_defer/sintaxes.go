package main

import "fmt"

func sintaxeDefer() {
	defer fmt.Println("Defer 1")
	defer fmt.Println("Defer 2")
	fmt.Println("Código normal")
}

func sintaxePanicComRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
		fmt.Println("Defer sempre executa")
	}()
	fmt.Println("Antes do panic")
	panic("Algo deu muito errado!")
}

func sintaxePanicAntesDoDefer() {
	panic("Panic acontece ANTES do defer")
	defer func() {
		fmt.Println("Isto nunca será executado")
	}()
}

func sintaxeRecoverForaDoDefer() {
	fmt.Println("Antes do panic")
	r := recover()
	panic("Panic aqui não será recuperado")
	fmt.Println(r)
}

func sintaxeDeferPanicRecoverComRetorno() {
	r := sintaxeDividirComPanicRecover(10, 0)
	fmt.Println(r)
	r2 := sintaxeDividirComPanicRecover(10, 2)
	fmt.Println(r2)
}

func sintaxeDividirComPanicRecover(a, b int) (res int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			res = -1
		}
	}()
	if b == 0 {
		panic("divisão por zero")
	}
	return a / b
}
