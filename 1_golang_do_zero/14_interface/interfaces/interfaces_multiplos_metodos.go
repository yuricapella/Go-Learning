package interfaces

import "fmt"

type Forma interface {
	Area() float64
	Perimetro() float64
}

type Retangulo struct {
	Largura float64
	Altura  float64
}

func (retangulo Retangulo) Area() float64 {
	return retangulo.Largura * retangulo.Altura
}

func (retangulo Retangulo) Perimetro() float64 {
	return 2 * (retangulo.Largura + retangulo.Altura)
}

type Circulo struct {
	Raio float64
}

func (circulo Circulo) Area() float64 {
	return 3.14159 * circulo.Raio * circulo.Raio
}

func (circulo Circulo) Perimetro() float64 {
	return 2 * 3.14159 * circulo.Raio
}

func DemonstrarInterfacesMultiplosMetodos() {
	fmt.Println("--- INTERFACES COM MÚLTIPLOS MÉTODOS ---")
	fmt.Println("Interfaces podem definir múltiplos métodos.")
	fmt.Println("Um tipo implementa a interface apenas se implementar TODOS os métodos.\n")

	fmt.Println("Exemplo 1: Interface Forma com dois métodos")
	fmt.Println("  type Forma interface {")
	fmt.Println("      Area() float64")
	fmt.Println("      Perimetro() float64")
	fmt.Println("  }\n")

	fmt.Println("Exemplo 2: Retangulo implementa Forma")
	retanguloExemplo := Retangulo{Largura: 5.0, Altura: 3.0}
	fmt.Printf("Retângulo (largura: %.1f, altura: %.1f):\n", retanguloExemplo.Largura, retanguloExemplo.Altura)
	fmt.Printf("  Área: %.2f\n", retanguloExemplo.Area())
	fmt.Printf("  Perímetro: %.2f\n", retanguloExemplo.Perimetro())
	fmt.Println()

	fmt.Println("Exemplo 3: Circulo implementa Forma")
	circuloExemplo := Circulo{Raio: 4.0}
	fmt.Printf("Círculo (raio: %.1f):\n", circuloExemplo.Raio)
	fmt.Printf("  Área: %.2f\n", circuloExemplo.Area())
	fmt.Printf("  Perímetro: %.2f\n", circuloExemplo.Perimetro())
	fmt.Println()

	fmt.Println("Exemplo 4: Usando a interface Forma")
	fmt.Println("Ambos Retangulo e Circulo podem ser tratados como Forma:\n")

	var formaExemplo Forma
	formaExemplo = retanguloExemplo
	fmt.Printf("Forma (Retângulo) - Área: %.2f, Perímetro: %.2f\n", formaExemplo.Area(), formaExemplo.Perimetro())

	formaExemplo = circuloExemplo
	fmt.Printf("Forma (Círculo) - Área: %.2f, Perímetro: %.2f\n", formaExemplo.Area(), formaExemplo.Perimetro())
	fmt.Println()

	fmt.Println("⚠️  IMPORTANTE:")
	fmt.Println("  - Um tipo deve implementar TODOS os métodos da interface")
	fmt.Println("  - Se faltar um método, o tipo não implementa a interface")
	fmt.Println("  - Interfaces com múltiplos métodos permitem polimorfismo")
	fmt.Println()
}
