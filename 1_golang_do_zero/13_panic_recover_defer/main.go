package main

import "fmt"

func main() {
	fmt.Println("=== PANIC, RECOVER E DEFER EM GO ===\n")

	fmt.Println("--- DEFER ---")
	fmt.Println("Defer adia a execução de uma função até que a função que a contém retorne.")
	fmt.Println("As funções defer são executadas em ordem LIFO.\n")
	exemploDefer()
	fmt.Println()

	fmt.Println("--- PANIC + RECOVER (didático) ---")
	fmt.Println("Panic interrompe o fluxo normal; recover só funciona dentro de defer.\n")
	exemploPanicComRecover()
	fmt.Println()

	fmt.Println("--- LIMITAÇÕES DE DEFER E RECOVER ---")
	fmt.Println("Exemplo 1: Panic antes do defer — defer não será executado.\n")
	exemploPanicAntesDoDefer()
	fmt.Println()
	fmt.Println("Exemplo 2: Recover fora do defer — recover não captura panics.\n")
	exemploRecoverForaDoDefer()
	fmt.Println()

	fmt.Println("--- DEFER + PANIC + RECOVER COM RETORNO ---")
	fmt.Println("Exemplo de função robusta com panic, defer, recover e retorno de valor seguro.\n")
	exemploDeferPanicRecoverComRetorno()
	fmt.Println()

	fmt.Println("--- RESUMO ---")
	fmt.Println("DEFER:")
	fmt.Println("  ✓ Adia a execução de uma função até o final da função em que está declarado.")
	fmt.Println("  ✓ Todos os defers são executados em ordem inversa (LIFO) ao término da função.")
	fmt.Println("  ✓ Muito útil para fechar arquivos, limpar recursos e garantir execução final.")
	fmt.Println()
	fmt.Println("PANIC:")
	fmt.Println("  ✓ Panic interrompe imediatamente o fluxo do programa, inicia o desenrolar da pilha.")
	fmt.Println("  ✓ Use apenas para condições realmente excepcionais e irrecuperáveis.")
	fmt.Println("  ✓ Todos os defers pendentes da pilha corrente são executados antes do programa encerrar.")
	fmt.Println()
	fmt.Println("RECOVER:")
	fmt.Println("  ✓ Só funciona dentro de funções defer (escopo do mesmo stackframe).")
	fmt.Println("  ✓ Permite recuperar o controle de uma goroutine em pânico, evitando crash.")
	fmt.Println("  ✓ Se usado fora de defer, não intercepta panics.")
	fmt.Println()
	fmt.Println("LIMITAÇÕES / ARMADILHAS:")
	fmt.Println("  ✓ Um panic acionado ANTES da declaração de um defer impede que o defer execute.")
	fmt.Println("  ✓ Recover chamado diretamente (fora do defer) não captura panics.")
	fmt.Println("  ✓ Sempre planeje a ordem de declaração de defers em funções críticas.")
	fmt.Println()
	fmt.Println("BOAS PRÁTICAS:")
	fmt.Println("  ✓ Empregue defer para liberar recursos com segurança (arquivos, lock/unlock, conexões).")
	fmt.Println("  ✓ Use panic como último recurso. Combine recover com log/contexto ao capturar.")
	fmt.Println("  ✓ Teste as armadilhas para entender ordem real de execução de defer/panic/recover.")
}

func exemploDefer() {
	fmt.Println("Entrando na função exemploDefer")
	defer fmt.Println("Defer 1: Executado por último")
	defer fmt.Println("Defer 2: Executado antes de Defer 1")
	fmt.Println("Saindo da função exemploDefer (defer executados depois deste print)")
}

func exemploPanicComRecover() {
	fmt.Println("Função vai panicar, mas recover dentro do defer captura e impede crash:")
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recover capturou panic: %v\n", r)
		}
		fmt.Println("Defer sempre executa ao fim da função (mesmo após panic)")
	}()
	fmt.Println("Antes do panic")
	panic("Algo deu muito errado!")
}

func exemploPanicAntesDoDefer() {
	fmt.Println("[Demonstração segura] Panic ocorre ANTES do defer: defer NÃO executa, recover NÃO captura.\n" +
		"Para evitar crash, panic está comentado — descomente para testar.")
	/*
		panic("Panic acontece ANTES do defer, logo defer não é executado.")
		defer func() {
			fmt.Println("Isto nunca será executado, pois defer vem depois do panic.")
		}()
	*/
	fmt.Println("(Sem crash: apenas exemplo de que ordem importa)")
}

func exemploRecoverForaDoDefer() {
	fmt.Println("[Demonstração segura] Recover chamado fora de defer NÃO captura panic.\n" +
		"Para evitar crash, panic está comentado — descomente para testar.")
	/*
		fmt.Println("Antes do panic")
		r := recover()
		panic("Panic aqui não será recuperado porque recover não está em defer")
		fmt.Printf("recover fora de defer => %v\n", r)
	*/
	fmt.Println("(Sem crash: apenas exemplo didático de limitação)")
}

func exemploDeferPanicRecoverComRetorno() {
	fmt.Println("Função robusta que trata panic e retorna valor seguro:")
	r := dividirComPanicRecover(10, 0)
	fmt.Printf("Resultado ao dividir (com recover): %d\n", r)
	r2 := dividirComPanicRecover(10, 2)
	fmt.Printf("Resultado ao dividir (normal): %d\n", r2)
}

func dividirComPanicRecover(a, b int) (res int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic capturado no dividirComPanicRecover: %v\n", r)
			res = -1 // valor seguro em caso de panic
		}
	}()
	if b == 0 {
		panic("divisão por zero (panic didático)")
	}
	return a / b
}
