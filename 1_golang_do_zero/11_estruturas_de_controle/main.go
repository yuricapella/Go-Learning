package main

import "fmt"

func main() {
	fmt.Println("=== ESTRUTURAS DE CONTROLE EM GO ===")
	fmt.Println()

	// ============================================
	// IF/ELSE BÁSICO
	// ============================================
	fmt.Println("--- IF/ELSE BÁSICO ---")
	numero := 5
	exemploIfElse(numero)
	fmt.Println()

	// ============================================
	// IF COM DECLARAÇÃO DE VARIÁVEL
	// ============================================
	fmt.Println("--- IF COM DECLARAÇÃO DE VARIÁVEL ---")
	exemploIfComDeclaracao(10)
	exemploIfComDeclaracao(-5)
	exemploIfComDeclaracao(-15)
	fmt.Println()

	// ============================================
	// IF/ELSE IF/ELSE
	// ============================================
	fmt.Println("--- IF/ELSE IF/ELSE ---")
	exemploIfElseIf(85)
	exemploIfElseIf(75)
	exemploIfElseIf(60)
	exemploIfElseIf(45)
	fmt.Println()

	// ============================================
	// SWITCH BÁSICO
	// ============================================
	fmt.Println("--- SWITCH BÁSICO ---")
	fmt.Println("No Go não existe 'break' - ele sai automaticamente após entrar no case.")
	fmt.Println()

	diaDaSemana := switchCase(7)
	fmt.Printf("switchCase(7) = %s\n", diaDaSemana)

	diaDaSemana = switchCase(1)
	fmt.Printf("switchCase(1) = %s\n", diaDaSemana)

	diaDaSemana = switchCase(10)
	fmt.Printf("switchCase(10) = %s\n", diaDaSemana)
	fmt.Println()

	// ============================================
	// SWITCH SEM EXPRESSÃO (SWITCH TRUE)
	// ============================================
	fmt.Println("--- SWITCH SEM EXPRESSÃO (SWITCH TRUE) ---")
	fmt.Println("Podemos usar switch sem expressão para fazer múltiplas comparações.")
	fmt.Println()

	diaDaSemana2 := switchCase2(1)
	fmt.Printf("switchCase2(1) = %s\n", diaDaSemana2)

	diaDaSemana2 = switchCase2(7)
	fmt.Printf("switchCase2(7) = %s\n", diaDaSemana2)
	fmt.Println()

	// ============================================
	// FALLTHROUGH
	// ============================================
	fmt.Println("--- FALLTHROUGH ---")
	fmt.Println("O 'fallthrough' faz o código continuar para o próximo case")
	fmt.Println("sem avaliar a condição (executa automaticamente).")
	fmt.Println()

	exemploSwitchFallthrough(1)
	exemploSwitchFallthrough(7)
	fmt.Println()

	// ============================================
	// SWITCH COM MÚLTIPLOS VALORES
	// ============================================
	fmt.Println("--- SWITCH COM MÚLTIPLOS VALORES ---")
	exemploSwitchMultiplosValores(2)
	exemploSwitchMultiplosValores(4)
	exemploSwitchMultiplosValores(6)
	exemploSwitchMultiplosValores(8)
	fmt.Println()

	// ============================================
	// RESUMO
	// ============================================
	fmt.Println("--- RESUMO ---")
	fmt.Println("IF/ELSE:")
	fmt.Println("  ✓ Sempre usa chaves {}")
	fmt.Println("  ✓ Pode declarar variáveis dentro do if")
	fmt.Println("  ✓ Suporta else if para múltiplas condições")
	fmt.Println()
	fmt.Println("SWITCH:")
	fmt.Println("  ✓ Não precisa de 'break' (sai automaticamente)")
	fmt.Println("  ✓ Pode usar sem expressão (switch true)")
	fmt.Println("  ✓ Pode ter múltiplos valores no case")
	fmt.Println("  ✓ 'fallthrough' continua para o próximo case")
	fmt.Println("  ✓ 'default' é opcional mas recomendado")
}

// ============================================
// EXEMPLOS DE IF/ELSE
// ============================================

func exemploIfElse(numero int) {
	fmt.Printf("Número: %d\n", numero)
	if numero >= 15 {
		fmt.Println("  → Número é maior ou igual a 15")
	} else {
		fmt.Println("  → Número é menor que 15")
	}
}

func exemploIfComDeclaracao(numero int) {
	fmt.Printf("Número: %d\n", numero)
	// Podemos declarar a variável dentro do if
	// Ela só existe dentro do escopo do if
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("  → Número é maior que 0")
	} else if outroNumero < -10 {
		fmt.Println("  → Número é menor que -10")
	} else {
		fmt.Println("  → Entre 0 e -10")
	}
	// outroNumero não existe aqui fora do if
}

func exemploIfElseIf(nota int) {
	fmt.Printf("Nota: %d\n", nota)
	if nota >= 90 {
		fmt.Println("  → Conceito: A (Excelente)")
	} else if nota >= 80 {
		fmt.Println("  → Conceito: B (Bom)")
	} else if nota >= 70 {
		fmt.Println("  → Conceito: C (Regular)")
	} else {
		fmt.Println("  → Conceito: D (Precisa melhorar)")
	}
}

// ============================================
// EXEMPLOS DE SWITCH
// ============================================

func switchCase(numero int) string {
	// No Go não existe break - ele sai automaticamente após entrar no case
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
	// Caso não seja nenhum número de 1 a 7, cai no default
	// O default também serve como return final que a função precisa
	default:
		return "Número inválido"
	}
}

func switchCase2(numero int) string {
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

func exemploSwitchFallthrough(numero int) {
	fmt.Printf("exemploFallthrough(%d):\n", numero)
	switch numero {
	case 1:
		fmt.Println("  → Entrou no case 1")
		fallthrough // Continua para o próximo case
	case 2:
		fmt.Println("  → Entrou no case 2")
	case 3:
		fmt.Println("  → Entrou no case 3")
	default:
		fmt.Println("  → Entrou no default")
	}
}

func exemploSwitchMultiplosValores(numero int) {
	fmt.Printf("Número: %d - ", numero)
	switch numero {
	case 2, 4, 6, 8:
		fmt.Println("É um número par entre 2 e 8")
	case 1, 3, 5, 7:
		fmt.Println("É um número ímpar entre 1 e 7")
	default:
		fmt.Println("Não está no intervalo esperado")
	}
}
