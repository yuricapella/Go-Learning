package main

import "fmt"

func sintaxeForBasico() {
	i := 0

	for i < 10 {
		i++
	}
}

func sintaxeForTradicional() {
	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}

	i := 0
	for j := &i; *j < 5; *j++ {
		fmt.Println("valor de i dentro do for: ", *j)
	}
	fmt.Println("valor total de i fora do for: ", i)
}

func sintaxeForRange() {
	letras := []string{"a", "b", "c", "d", "e"}

	for indice, valor := range letras {
		fmt.Println(indice, valor)
	}
}

func sintaxeForRangeComValorOculto() {
	letras := []string{"a", "b", "c", "d", "e"}

	for _, valor := range letras {
		fmt.Println(valor)
	}

	for indice := range letras {
		fmt.Println(indice)
	}
}

func sintaxeForRangeString() {
	texto := "Hello"

	for indice, rune := range texto {
		fmt.Println(indice, rune)
		caractere := string(rune)
		fmt.Println(caractere)
	}

	texto2 := "Olá, 世界"
	for indice, rune := range texto2 {
		fmt.Println(indice, rune)
		caractere := string(rune)
		fmt.Println(caractere)
	}
}

func sintaxeForRangeMap() {
	mapa := map[string]string{
		"nome":  "João",
		"idade": "20",
		"email": "joao@example.com",
	}

	for chave, valor := range mapa {
		fmt.Println(chave, valor)
	}

	for chave := range mapa {
		fmt.Println(chave)
	}
}

func sintaxeBreakContinue() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
	}
}

func sintaxeForInfinito() {
	contador := 0
	for {
		contador++

		if contador >= 3 {
			break
		}
	}
}
