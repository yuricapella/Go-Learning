package main

import "errors"

func sintaxeTiposInteiros() {
	var numero4 int8 = 100
	var numero int16 = 100
	var numero2 int32 = 1000000000
	var numero3 int64 = 1000000000000000000
	var numero5 int = 1000000000000000000

	var numero6 uint = 100

	var numero7 rune = 123456
	var numero8 byte = 100

	_ = numero4
	_ = numero
	_ = numero2
	_ = numero3
	_ = numero5
	_ = numero6
	_ = numero7
	_ = numero8
}

func sintaxeTiposFlutuantes() {
	var numero9 float32 = 123.45
	var numero10 float64 = 1234567890.1234567890
	numero11 := 123.45

	_ = numero9
	_ = numero10
	_ = numero11
}

func sintaxeTiposString() {
	var texto string = "Texto"
	texto2 := "Texto 2"
	char := 'A'

	_ = texto
	_ = texto2
	_ = char
}

func sintaxeValoresZero() {
	var texto3 string
	var numero13 int
	var booleano bool
	var erro error

	_ = texto3
	_ = numero13
	_ = booleano
	_ = erro
