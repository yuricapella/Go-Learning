package interfaces

import "fmt"

func DemonstrarPolimorfismo() {
	fmt.Println("--- POLIMORFISMO COM INTERFACES ---")
	fmt.Println("Polimorfismo permite tratar diferentes tipos de forma uniforme através de interfaces.")
	fmt.Println("Uma função que aceita uma interface pode trabalhar com qualquer tipo que a implemente.\n")

	fmt.Println("Exemplo 1: Função que aceita interface Animal")
	fmt.Println("A função ProcessarAnimais aceita qualquer tipo que implemente Animal:\n")

	cachorroExemplo := Cachorro{Nome: "Rex"}
	gatoExemplo := Gato{Nome: "Mimi"}

	listaAnimais := []Animal{cachorroExemplo, gatoExemplo}
	ProcessarAnimais(listaAnimais)
	fmt.Println()

	fmt.Println("Exemplo 2: Função que aceita interface Forma")
	fmt.Println("A função ExibirInformacoesForma aceita qualquer tipo que implemente Forma:\n")

	retanguloExemplo := Retangulo{Largura: 5.0, Altura: 3.0}
	circuloExemplo := Circulo{Raio: 4.0}

	listaFormas := []Forma{retanguloExemplo, circuloExemplo}
	ExibirInformacoesFormas(listaFormas)
	fmt.Println()

	fmt.Println("Exemplo 3: Polimorfismo em ação")
	fmt.Println("O mesmo código funciona para diferentes tipos:\n")

	FazerAnimaisEmitiremSom(cachorroExemplo)
	FazerAnimaisEmitiremSom(gatoExemplo)
	fmt.Println()

	fmt.Println("⚠️  IMPORTANTE:")
	fmt.Println("  - Polimorfismo permite código mais flexível e reutilizável")
	fmt.Println("  - Funções que aceitam interfaces são mais genéricas")
	fmt.Println("  - Novos tipos podem ser adicionados sem modificar o código existente")
	fmt.Println("  - O comportamento específico de cada tipo é mantido")
	fmt.Println()
}

func ProcessarAnimais(animais []Animal) {
	fmt.Println("Processando animais:")
	for indice, animal := range animais {
		fmt.Printf("  [%d] %s\n", indice+1, animal.EmitirSom())
	}
}

func ExibirInformacoesFormas(formas []Forma) {
	fmt.Println("Informações das formas:")
	for indice, forma := range formas {
		fmt.Printf("  [%d] Área: %.2f, Perímetro: %.2f\n", indice+1, forma.Area(), forma.Perimetro())
	}
}

func FazerAnimaisEmitiremSom(animal Animal) {
	fmt.Printf("Animal emite: %s\n", animal.EmitirSom())
}
