package interfaces

import "fmt"

type Animal interface {
	EmitirSom() string
}

type Cachorro struct {
	Nome string
}

func (cachorro Cachorro) EmitirSom() string {
	return "Au au!"
}

type Gato struct {
	Nome string
}

func (gato Gato) EmitirSom() string {
	return "Miau!"
}

func DemonstrarInterfacesBasicas() {
	fmt.Println("--- INTERFACES BÁSICAS ---")
	fmt.Println("Interfaces definem um conjunto de métodos que um tipo deve implementar.")
	fmt.Println("Em Go, a implementação é implícita: se um tipo tem os métodos da interface, ele a implementa.\n")

	fmt.Println("Exemplo 1: Definição de interface")
	fmt.Println("Interface Animal com método EmitirSom():")
	fmt.Println("  type Animal interface {")
	fmt.Println("      EmitirSom() string")
	fmt.Println("  }\n")

	fmt.Println("Exemplo 2: Implementação implícita")
	fmt.Println("Structs Cachorro e Gato implementam Animal automaticamente")
	fmt.Println("porque possuem o método EmitirSom().\n")

	cachorroExemplo := Cachorro{Nome: "Rex"}
	gatoExemplo := Gato{Nome: "Mimi"}

	fmt.Printf("Cachorro: %s emite: %s\n", cachorroExemplo.Nome, cachorroExemplo.EmitirSom())
	fmt.Printf("Gato: %s emite: %s\n", gatoExemplo.Nome, gatoExemplo.EmitirSom())
	fmt.Println()

	fmt.Println("Exemplo 3: Usando a interface como tipo")
	fmt.Println("Podemos criar variáveis do tipo Animal que aceitam qualquer tipo que implemente Animal:\n")

	var animalExemplo Animal
	animalExemplo = cachorroExemplo
	fmt.Printf("Animal (Cachorro): %s\n", animalExemplo.EmitirSom())

	animalExemplo = gatoExemplo
	fmt.Printf("Animal (Gato): %s\n", animalExemplo.EmitirSom())
	fmt.Println()

	fmt.Println("⚠️  IMPORTANTE:")
	fmt.Println("  - Interfaces definem comportamentos, não dados")
	fmt.Println("  - Implementação é implícita (duck typing)")
	fmt.Println("  - Se um tipo tem os métodos da interface, ele a implementa automaticamente")
	fmt.Println()
}
