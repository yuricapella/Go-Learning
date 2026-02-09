package formas

import (
	"fmt"
	"math"
)

type Forma interface {
	Area() float64
}

func ExibirArea(forma Forma) string {
	return fmt.Sprintf("A área da forma é: %.2f", forma.Area())
}

type Retangulo struct {
	Altura  float64
	Largura float64
}

func (retangulo Retangulo) Area() float64 {
	return retangulo.Altura * retangulo.Largura
}

type Circulo struct {
	Raio float64
}

func (circulo Circulo) Area() float64 {
	return math.Pi * math.Pow(circulo.Raio, 2)
}
