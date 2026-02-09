package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	// int8, int16, int32, int64    / positivos e negativos

	var numero4 int8 = 100
	fmt.Println(numero4)

	var numero int16 = 100
	fmt.Println(numero)

	var numero2 int32 = 1000000000
	fmt.Println(numero2)

	var numero3 int64 = 1000000000000000000
	fmt.Println(numero3)

	// tambem tem o int que é implicito, pega o tipo do computador, se for 64 vira int64, além de comportar numeros negativos.
	var numero5 int = 1000000000000000000
	fmt.Println(numero5)

	// uint8, uint16, uint32, uint64    / apenas numeros positivos
	var numero6 uint = 100
	fmt.Println(numero6)

	//alias -- rune = int32, usado quando mexemos com numeros que representam caracteres normalmente da tabel ascii
	var numero7 rune = 123456
	fmt.Println(numero7)
	fmt.Println("rune é um alias de:", reflect.TypeOf(numero7))

	//alias -- byte = uint8
	var numero8 byte = 100
	fmt.Println(numero8)
	fmt.Println("byte é um alias de:", reflect.TypeOf(numero8))

	// numeros de ponto flutuante, float32 ou float64
	var numero9 float32 = 123.45
	fmt.Println(numero9)

	var numero10 float64 = 1234567890.1234567890
	fmt.Println(numero10)

	// inferência de float sempre vira float64, não depende do sistema, apenas int e uint dependem
	numero11 := 123.45
	fmt.Println(numero11)
	fmt.Println("Inferencia de numero com virgula vira:", reflect.TypeOf(numero11))

	// inferência de int sempre vira int, ou seja dependendo do sistema vira int32 ou int64
	numero12 := 123
	fmt.Println(numero12)
	fmt.Println("Inferencia de numero inteiro vira:", reflect.TypeOf(numero12))

	// string
	var texto string = "Texto"
	fmt.Println(texto)

	texto2 := "Texto 2"
	fmt.Println(texto2)
	fmt.Println("Inferencia de string vira:", reflect.TypeOf(texto2))

	// o mais proximo que temos de char no go é o numero do caractere que tem na tabela ascii e declaramos ele com aspas simples 'A' e só pode um caractere
	char := 'A'
	fmt.Println(char)
	fmt.Println("char é um tipo de dado:", reflect.TypeOf(char))

	// valores vazios, string = "", int = 0, boolean = false, error = nil / nulo
	var texto3 string
	fmt.Println(texto3)

	var numero13 int
	fmt.Println(numero13)

	var booleano bool
	fmt.Println(booleano)

	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("Erro de teste")
	fmt.Println(erro2)
}
