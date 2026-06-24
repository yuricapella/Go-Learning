package sintaxes

import "fmt"

// --- CONCEITO: Struct e Zero Value ---
//
// O que é:
//   Struct é a única forma de criar tipos compostos em Go.
//   Go não tem `class` — é uma linguagem orientada a dados, não a objetos.
//
// O que o compilador recebe ao declarar um tipo:
//   1. A quantidade de memória a alocar
//   2. A representação dos dados
//
// Regras do instrutor (Bill Kennedy - Ardan Labs):
//   - Use `var` para zero value (intenção clara)
//   - Use literal construction `{}` apenas quando inicializar com valores reais
//   - Evite construção parcial (var + set campos depois) — causa bugs em produção
//   - Não inclua campos no literal se o valor for zero — desnecessário

type example struct {
	flag    bool
	counter int16
	pi      float32
}

func DemonstrarStruct() {
	fmt.Println("=== AULA 01: Go Syntax - Struct ===")
	fmt.Println()

	Exemplo1ZeroValue()
	Exemplo2LiteralConstruction()
	Exemplo3FormatacaoStruct()
	Exemplo4EvitarConstrucaoParcial()
}

// Exemplo1ZeroValue — var declara e inicializa com zero value
func Exemplo1ZeroValue() {
	fmt.Println("--- Exemplo 1: Zero Value com var ---")

	var e1 example
	// flag=false, counter=0, pi=0 (zero value de cada tipo)

	fmt.Printf("e1 = %+v\n\n", e1)
}

// Exemplo2LiteralConstruction — inicializa com valores reais via {}
func Exemplo2LiteralConstruction() {
	fmt.Println("--- Exemplo 2: Literal Construction ---")

	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}
	// campo não incluído = zero value automaticamente
	// vírgula obrigatória em toda linha (consistência, facilita reordenação)

	fmt.Printf("e2 = %+v\n\n", e2)
}

// Exemplo3FormatacaoStruct — três estilos de formatação para structs
func Exemplo3FormatacaoStruct() {
	fmt.Println("--- Exemplo 3: Estilos de Formatação ---")

	e := example{flag: true, counter: 5, pi: 1.5}

	fmt.Printf("%%v  → %v\n", e)   // sem nome dos campos
	fmt.Printf("%%+v → %+v\n", e)  // com nome dos campos (preferido pelo instrutor)
	fmt.Printf("%%#v → %#v\n\n", e) // com tipo completo (mais verboso)
}

// Exemplo4EvitarConstrucaoParcial — mostra o padrão correto vs problemático
func Exemplo4EvitarConstrucaoParcial() {
	fmt.Println("--- Exemplo 4: Evitar Construção Parcial ---")

	// ⚠️ RUIM: construção parcial — risco de retornar valor incompleto por acidente
	// var e example
	// e.flag = true  ← não faça isso

	// ✓ BOM: reúna os dados antes, construa de uma vez
	flagValue := true
	e := example{
		flag: flagValue,
	}
	// counter e pi ficam em zero value — explícito e intencional

	fmt.Printf("e = %+v\n\n", e)

	// ✓ ACEITÁVEL: empty literal construction apenas em return sem variável prévia
	// return example{}
	fmt.Println("⚠️  IMPORTANTE:")
	fmt.Println("  - `var` = zero value (use para deixar intenção clara)")
	fmt.Println("  - `example{}` = empty literal (não é zero value para todos os tipos)")
	fmt.Println("  - Construção parcial (var + set depois) = smell, evite")
	fmt.Println("  - Dot operator: value.field (sem novidade, igual a outras linguagens)")
}
