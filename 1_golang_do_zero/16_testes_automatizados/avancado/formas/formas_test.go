package formas

import (
	"fmt"
	"math"
	"testing"
)

/* Podemos fazer testes com funções que utilizam subtestes
Se usarmos Fatalf ele para no primeiro erro, enquanto errorf ele roda tudo

Podemos utilizar go test -v ./... na pasta principal dos testes para rodar todos os testes
*/

func TestArea(test *testing.T) {
	test.Run("Retangulo", func(test *testing.T) {
		retangulo := Retangulo{Altura: 10, Largura: 12}
		areaEsperada := float64(120)
		areaRecebida := retangulo.Area()

		if areaRecebida != areaEsperada {
			test.Fatalf("Area recebida: %.2f, Area esperada: %.2f", areaRecebida, areaEsperada)
		}
	})

	test.Run("Circulo", func(test *testing.T) {
		circulo := Circulo{Raio: 10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circulo.Area()

		if areaRecebida != areaEsperada {
			test.Fatalf("Area recebida: %.2f, Area esperada: %.2f", areaRecebida, areaEsperada)
		}
	})

}

func TestExibirArea(test *testing.T) {
	retangulo := Retangulo{Altura: 10, Largura: 12}
	areaEsperadaRetangulo := float64(120)
	fraseEsperadaRetangulo := fmt.Sprintf("A área da forma é: %.2f", areaEsperadaRetangulo)

	circulo := Circulo{Raio: 10}
	areaEsperadaCirculo := float64(math.Pi * 100)
	fraseEsperadaCirculo := fmt.Sprintf("A área da forma é: %.2f", areaEsperadaCirculo)

	formas := []Forma{retangulo, circulo}

	for _, forma := range formas {
		fraseRecebida := ExibirArea(forma)

		if fraseRecebida != fraseEsperadaRetangulo && fraseRecebida != fraseEsperadaCirculo {
			test.Fatalf("Frase recebida: %s, Frase esperada: %s", fraseRecebida, fraseEsperadaRetangulo)
		}
	}

}
