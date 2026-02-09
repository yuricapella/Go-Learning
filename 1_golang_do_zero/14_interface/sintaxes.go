package main

import "fmt"

type AnimalSintaxe interface {
	EmitirSom() string
}

type CachorroSintaxe struct {
	Nome string
}

func (cachorro CachorroSintaxe) EmitirSom() string {
	return "Au au!"
}

type GatoSintaxe struct {
	Nome string
}

func (gato GatoSintaxe) EmitirSom() string {
	return "Miau!"
}

func sintaxeInterfaceBasica() {
	cachorroExemplo := CachorroSintaxe{Nome: "Rex"}
	gatoExemplo := GatoSintaxe{Nome: "Mimi"}
	fmt.Println(cachorroExemplo.EmitirSom())
	fmt.Println(gatoExemplo.EmitirSom())

	var animalExemplo AnimalSintaxe
	animalExemplo = cachorroExemplo
	fmt.Println(animalExemplo.EmitirSom())
	animalExemplo = gatoExemplo
	fmt.Println(animalExemplo.EmitirSom())
}

func sintaxeInterfaceVazia() {
	var valorQualquer interface{}
	valorQualquer = "texto"
	fmt.Println(valorQualquer)
	valorQualquer = 42
	fmt.Println(valorQualquer)
	valorQualquer = true
	fmt.Println(valorQualquer)

	listaValores := []interface{}{
		"texto",
		42,
		true,
		3.14,
	}
	for _, valor := range listaValores {
		fmt.Println(valor)
	}

	exibirTipoSintaxe("texto")
	exibirTipoSintaxe(42)
}

func exibirTipoSintaxe(valorQualquer interface{}) {
	fmt.Printf("Valor: %v, Tipo: %T\n", valorQualquer, valorQualquer)
}

func sintaxeTypeAssertion() {
	var valorQualquer interface{} = "texto"
	textoConvertido := valorQualquer.(string)
	fmt.Println(textoConvertido)

	valorQualquer = 42
	numeroConvertido, conversaoSucedida := valorQualquer.(int)
	if conversaoSucedida {
		fmt.Println(numeroConvertido)
	}

	valorQualquer = "texto"
	numeroConvertido, conversaoSucedida = valorQualquer.(int)
	if conversaoSucedida {
		fmt.Println(numeroConvertido)
	} else {
		fmt.Println("Conversão falhou")
	}
}

func sintaxeTypeSwitch() {
	var valorQualquer interface{} = 42
	switch valorVerificado := valorQualquer.(type) {
	case string:
		fmt.Println(valorVerificado)
	case int:
		fmt.Println(valorVerificado)
	case bool:
		fmt.Println(valorVerificado)
	}
}

type FormaSintaxe interface {
	Area() float64
	Perimetro() float64
}

type RetanguloSintaxe struct {
	Largura float64
	Altura  float64
}

func (retangulo RetanguloSintaxe) Area() float64 {
	return retangulo.Largura * retangulo.Altura
}

func (retangulo RetanguloSintaxe) Perimetro() float64 {
	return 2 * (retangulo.Largura + retangulo.Altura)
}

type CirculoSintaxe struct {
	Raio float64
}

func (circulo CirculoSintaxe) Area() float64 {
	return 3.14159 * circulo.Raio * circulo.Raio
}

func (circulo CirculoSintaxe) Perimetro() float64 {
	return 2 * 3.14159 * circulo.Raio
}

func sintaxeInterfaceMultiplosMetodos() {
	retanguloExemplo := RetanguloSintaxe{Largura: 5.0, Altura: 3.0}
	fmt.Println(retanguloExemplo.Area())
	fmt.Println(retanguloExemplo.Perimetro())

	circuloExemplo := CirculoSintaxe{Raio: 4.0}
	fmt.Println(circuloExemplo.Area())
	fmt.Println(circuloExemplo.Perimetro())

	var formaExemplo FormaSintaxe
	formaExemplo = retanguloExemplo
	fmt.Println(formaExemplo.Area())
	fmt.Println(formaExemplo.Perimetro())

	formaExemplo = circuloExemplo
	fmt.Println(formaExemplo.Area())
	fmt.Println(formaExemplo.Perimetro())
}

func sintaxePolimorfismo() {
	cachorroExemplo := CachorroSintaxe{Nome: "Rex"}
	gatoExemplo := GatoSintaxe{Nome: "Mimi"}
	listaAnimais := []AnimalSintaxe{cachorroExemplo, gatoExemplo}
	for _, animalExemplo := range listaAnimais {
		fmt.Println(animalExemplo.EmitirSom())
	}

	retanguloExemplo := RetanguloSintaxe{Largura: 5.0, Altura: 3.0}
	circuloExemplo := CirculoSintaxe{Raio: 4.0}
	listaFormas := []FormaSintaxe{retanguloExemplo, circuloExemplo}
	for _, formaExemplo := range listaFormas {
		fmt.Println(formaExemplo.Area())
		fmt.Println(formaExemplo.Perimetro())
	}

	fazerAnimaisEmitiremSomSintaxe(cachorroExemplo)
	fazerAnimaisEmitiremSomSintaxe(gatoExemplo)
}

func fazerAnimaisEmitiremSomSintaxe(animal AnimalSintaxe) {
	fmt.Println(animal.EmitirSom())
}
