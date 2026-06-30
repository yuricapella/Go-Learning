package didatica

import "fmt"

type example struct {
	flag    bool
	counter int16
	pi      float32
}

func DemonstrarStruct() {
	fmt.Println("=== AULA 01: Struct, Zero Value e Literal Construction ===")
	fmt.Println()

	fmt.Println("--- O que e struct? ---")
	fmt.Println("Struct e a forma de criar tipos compostos em Go.")
	fmt.Println("Go nao usa `class` para modelar dados. Voce declara um novo tipo com `type` e define seus campos com `struct`.")
	fmt.Println()

	fmt.Println("--- Por que isso existe? ---")
	fmt.Println("Tipos built-in como int, bool e string resolvem valores simples.")
	fmt.Println("Mas aplicacoes reais precisam agrupar dados relacionados em um tipo proprio.")
	fmt.Println("Ao declarar uma struct, voce diz ao compilador qual memoria precisa ser alocada e qual e a representacao dos dados.")
	fmt.Println()

	Exemplo1DeclararTipo()
	Exemplo2ZeroValue()
	Exemplo3LiteralConstruction()
	Exemplo4FormatacaoStruct()
	Exemplo5EvitarConstrucaoParcial()
	PontosImportantes()
}

func Exemplo1DeclararTipo() {
	fmt.Println("--- Exemplo 1: declarar um tipo com struct ---")
	fmt.Println("A sintaxe em Go segue `nome tipo`, como `flag bool` e `counter int16`.")
	fmt.Println("Isso pode parecer invertido no comeco, mas a leitura fica: campo `flag` do tipo `bool`.")
	fmt.Println()
}

func Exemplo2ZeroValue() {
	fmt.Println("--- Exemplo 2: zero value com var ---")

	var e1 example
	fmt.Printf("e1 = %+v\n", e1)
	fmt.Println("`var e1 example` cria um valor completo de `example` com todos os campos em zero value.")
	fmt.Println("Nesse caso: flag=false, counter=0 e pi=0.")
	fmt.Println()
}

func Exemplo3LiteralConstruction() {
	fmt.Println("--- Exemplo 3: literal construction ---")

	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	fmt.Printf("e2 = %+v\n", e2)
	fmt.Println("Use literal construction quando voce quer inicializar campos com valores reais.")
	fmt.Println("Se um campo deve continuar em zero value, voce pode simplesmente nao inclui-lo.")
	fmt.Println()
}

func Exemplo4FormatacaoStruct() {
	fmt.Println("--- Exemplo 4: formatacao de struct ---")

	e := example{flag: true, counter: 5, pi: 1.5}

	fmt.Printf("%%v  = %v\n", e)
	fmt.Printf("%%+v = %+v\n", e)
	fmt.Printf("%%#v = %#v\n", e)
	fmt.Println("O formato com plus costuma ser uma boa opcao didatica porque mostra os nomes dos campos sem muito ruido.")
	fmt.Println()
}

func Exemplo5EvitarConstrucaoParcial() {
	fmt.Println("--- Exemplo 5: evitar construcao parcial ---")
	fmt.Println("Evite criar uma struct vazia e preencher campos espalhados depois quando voce ja sabe os dados.")
	fmt.Println("Isso reduz o risco de retornar ou usar um valor incompleto por acidente.")

	flagValue := true
	e := example{
		flag: flagValue,
	}

	fmt.Printf("construcao direta = %+v\n", e)
	fmt.Println()
}

func PontosImportantes() {
	fmt.Println("--- Pontos importantes ---")
	fmt.Println("- `type Nome struct { ... }` cria um tipo composto.")
	fmt.Println("- `var e example` deixa claro que voce quer o zero value.")
	fmt.Println("- `example{...}` e literal construction: use quando for inicializar valores.")
	fmt.Println("- Empty literal `example{}` pode parecer zero value, mas conceitualmente e uma construcao vazia.")
	fmt.Println("- Prefira reunir os dados e construir a struct de uma vez.")
}
