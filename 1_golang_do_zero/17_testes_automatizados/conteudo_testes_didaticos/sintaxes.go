package main

import (
	"testing"
)

// sintaxeEstruturaBasicaTeste demonstra a estrutura básica de um teste
func sintaxeEstruturaBasicaTeste(t *testing.T) {
	valorEsperado := 10
	valorRecebido := 10

	if valorRecebido != valorEsperado {
		t.Errorf("Valor recebido: %d, Valor esperado: %d", valorRecebido, valorEsperado)
	}
}

// sintaxeTesteComError demonstra uso de t.Error
func sintaxeTesteComError(t *testing.T) {
	valorEsperado := 10
	valorRecebido := 5

	if valorRecebido != valorEsperado {
		t.Error("Valores não são iguais")
	}
}

// sintaxeTesteComFatal demonstra uso de t.Fatal
func sintaxeTesteComFatal(t *testing.T) {
	valorEsperado := 10
	valorRecebido := 5

	if valorRecebido != valorEsperado {
		t.Fatal("Teste deve parar aqui")
	}
	// Este código não será executado se Fatal for chamado
}

// sintaxeTesteComTabela demonstra teste com múltiplos cenários
func sintaxeTesteComTabela(t *testing.T) {
	type cenarioTeste struct {
		entrada   int
		esperado  int
		descricao string
	}

	cenarios := []cenarioTeste{
		{entrada: 1, esperado: 2, descricao: "caso 1"},
		{entrada: 2, esperado: 4, descricao: "caso 2"},
		{entrada: 3, esperado: 6, descricao: "caso 3"},
	}

	for _, cenario := range cenarios {
		resultado := cenario.entrada * 2
		if resultado != cenario.esperado {
			t.Errorf("Cenário %s: resultado %d, esperado %d", cenario.descricao, resultado, cenario.esperado)
		}
	}
}

// sintaxeSubteste demonstra uso de test.Run
func sintaxeSubteste(t *testing.T) {
	t.Run("Subteste 1", func(t *testing.T) {
		valor := 10
		if valor != 10 {
			t.Error("Falhou")
		}
	})

	t.Run("Subteste 2", func(t *testing.T) {
		valor := 20
		if valor != 20 {
			t.Error("Falhou")
		}
	})
}

// sintaxeSubtesteParalelo demonstra subteste com execução paralela
func sintaxeSubtesteParalelo(t *testing.T) {
	t.Run("Paralelo 1", func(t *testing.T) {
		t.Parallel()
		// Código do teste
	})

	t.Run("Paralelo 2", func(t *testing.T) {
		t.Parallel()
		// Código do teste
	})
}

// sintaxeTesteComHelper demonstra uso de helper functions
func sintaxeTesteComHelper(t *testing.T) {
	verificarIgualdade(t, 10, 10)
}

func verificarIgualdade(t *testing.T, recebido int, esperado int) {
	t.Helper()
	if recebido != esperado {
		t.Errorf("Recebido: %d, Esperado: %d", recebido, esperado)
	}
}

// sintaxeTesteComSkip demonstra como pular um teste
func sintaxeTesteComSkip(t *testing.T) {
	t.Skip("Pulando este teste")
	// Código não será executado
}

// sintaxeTesteComCleanup demonstra uso de cleanup
func sintaxeTesteComCleanup(t *testing.T) {
	// Setup
	recurso := "recurso criado"

	// Cleanup será executado ao final do teste
	t.Cleanup(func() {
		// Limpar recurso
		_ = recurso
	})

	// Código do teste
	_ = recurso
}
