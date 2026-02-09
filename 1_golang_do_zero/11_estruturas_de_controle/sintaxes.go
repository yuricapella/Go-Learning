package main

import "fmt"

func sintaxeIfElse(numero int) {
	if numero >= 15 {
		fmt.Println("numero é maior ou igual a 15")
	} else {
		fmt.Println("numero é menor que 15")
	}
}

func sintaxeIfComDeclaracao(numero int) {
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("numero é maior que 0")
	} else if outroNumero < -10 {
		fmt.Println("numero é menor que -10")
	} else {
		fmt.Println("numero é entre 0 e -10")
	}
}

func sintaxeIfElseIf(nota int) {
	if nota >= 90 {
		fmt.Println("nota é maior ou igual a 90")
	} else if nota >= 80 {
		fmt.Println("nota é maior ou igual a 80")
	} else if nota >= 70 {
		fmt.Println("nota é maior ou igual a 70")
	} else {
		fmt.Println("nota é menor que 70")
	}
}

func sintaxeSwitchBasico(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

func sintaxeSwitchSemExpressao(numero int) string {
	diaDaSemana := ""
	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
	case numero == 2:
		diaDaSemana = "Segunda-feira"
	case numero == 3:
		diaDaSemana = "Terça-feira"
	case numero == 4:
		diaDaSemana = "Quarta-feira"
	case numero == 5:
		diaDaSemana = "Quinta-feira"
	case numero == 6:
		diaDaSemana = "Sexta-feira"
	case numero == 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Número inválido"
	}
	return diaDaSemana
}

func sintaxeSwitchFallthrough(numero int) {
	switch numero {
	case 1:
		fmt.Println("entrou no case 1")
		fallthrough
	case 2:
		fmt.Println("entrou no case 2")
	default:
		fmt.Println("entrou no default")
	}
}

func sintaxeSwitchMultiplosValores(numero int) {
	switch numero {
	case 2, 4, 6, 8:
		fmt.Println("entrou no case 2, 4, 6, 8")
	case 1, 3, 5, 7:
		fmt.Println("entrou no case 1, 3, 5, 7")
	default:
		fmt.Println("entrou no default")
	}
}
