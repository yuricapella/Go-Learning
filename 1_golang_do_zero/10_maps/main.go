package main

import "fmt"

func main() {
	fmt.Println("=== MAPS EM GO ===\n")

	// ============================================
	// MAPS - CONCEITOS BÁSICOS
	// ============================================
	fmt.Println("--- MAPS - CONCEITOS BÁSICOS ---")
	fmt.Println("Maps são coleções de pares chave-valor.")
	fmt.Println("Todos os campos devem seguir o mesmo tipo [tipoChave]tipoValor")
	fmt.Println("Se for map[string]string, todas as chaves são string e todos os valores são string.\n")

	// ============================================
	// MAP SIMPLES - STRING PARA STRING
	// ============================================
	fmt.Println("--- MAP SIMPLES - STRING PARA STRING ---")
	
	usuario := map[string]string{
		"nome":      "João",
		"sobrenome": "Silva",
	}

	fmt.Printf("usuario completo: %v\n", usuario)
	fmt.Printf("Acessando chave 'nome': usuario[\"nome\"] = %s\n", usuario["nome"])
	fmt.Printf("Acessando chave 'sobrenome': usuario[\"sobrenome\"] = %s\n\n", usuario["sobrenome"])

	// ============================================
	// MAPS ANINHADOS (MAPS DENTRO DE MAPS)
	// ============================================
	fmt.Println("--- MAPS ANINHADOS (MAPS DENTRO DE MAPS) ---")
	fmt.Println("Podemos ter maps dentro de maps.")
	fmt.Println("Se queremos adicionar mais informações, devemos seguir o tipo:")
	fmt.Println("chave=string, valor=map[string]string\n")

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiroNome": "João",
			"ultimoNome":   "Silva",
		},
		"endereco": {
			"rua":    "Rua 1",
			"numero": "123",
			"cidade": "São Paulo",
			"estado": "SP",
			"cep":    "1234567890",
		},
		"profissao": {
			"nome":  "Engenheiro",
			"cargo": "Desenvolvedor",
		},
	}

	fmt.Printf("usuario2 completo: %+v\n\n", usuario2)
	fmt.Println("Acessando map aninhado:")
	fmt.Printf("usuario2[\"nome\"] = %v\n", usuario2["nome"])
	fmt.Println()
	fmt.Println("Acessando valores dentro do map aninhado:")
	fmt.Printf("usuario2[\"nome\"][\"primeiroNome\"] = %s\n", usuario2["nome"]["primeiroNome"])
	fmt.Printf("usuario2[\"nome\"][\"ultimoNome\"] = %s\n\n", usuario2["nome"]["ultimoNome"])

	// ============================================
	// DELETANDO ELEMENTOS DO MAP
	// ============================================
	fmt.Println("--- DELETANDO ELEMENTOS DO MAP ---")
	fmt.Println("Para deletar um campo, usamos: delete(map, chave)\n")

	fmt.Printf("usuario2 antes de deletar: %+v\n", usuario2)
	delete(usuario2, "profissao")
	fmt.Printf("usuario2 após deletar 'profissao': %+v\n\n", usuario2)

	// ============================================
	// ADICIONANDO ELEMENTOS AO MAP
	// ============================================
	fmt.Println("--- ADICIONANDO ELEMENTOS AO MAP ---")
	fmt.Println("Para adicionar um campo, usamos: map[chave] = valor\n")
	fmt.Println("⚠️  Importante: o tipo do valor deve corresponder ao tipo do map!")
	fmt.Println("   Como usuario2 é map[string]map[string]string,")
	fmt.Println("   o valor deve ser map[string]string, não string.\n")

	// Tentativa incorreta (comentada para evitar erro)
	// usuario2["telefone"] = "1234567890" // ERRO: tipo incompatível

	// Forma correta
	usuario2["telefone"] = map[string]string{
		"ddd":    "11",
		"numero": "1234567890",
	}
	fmt.Printf("usuario2 após adicionar 'telefone': %+v\n\n", usuario2)

	// ============================================
	// VERIFICANDO EXISTÊNCIA DE CHAVE
	// ============================================
	fmt.Println("--- VERIFICANDO EXISTÊNCIA DE CHAVE ---")
	fmt.Println("Ao acessar uma chave, podemos verificar se ela existe:\n")

	valor, existe := usuario2["nome"]
	if existe {
		fmt.Printf("Chave 'nome' existe! Valor: %v\n", valor)
	} else {
		fmt.Println("Chave 'nome' não existe!")
	}

	valor2, existe2 := usuario2["idade"]
	if existe2 {
		fmt.Printf("Chave 'idade' existe! Valor: %v\n", valor2)
	} else {
		fmt.Printf("Chave 'idade' não existe! Valor zero: %v\n", valor2)
	}
	fmt.Println()

	// ============================================
	// MAP COM MAKE
	// ============================================
	fmt.Println("--- MAP COM MAKE ---")
	fmt.Println("Podemos criar maps vazios usando make():\n")

	usuario3 := make(map[string]string)
	fmt.Printf("usuario3 (vazio): %v\n", usuario3)
	
	usuario3["nome"] = "Maria"
	usuario3["email"] = "maria@example.com"
	fmt.Printf("usuario3 após adicionar valores: %v\n\n", usuario3)

	// ============================================
	// ITERANDO SOBRE UM MAP
	// ============================================
	fmt.Println("--- ITERANDO SOBRE UM MAP ---")
	fmt.Println("Podemos iterar sobre um map usando range:\n")

	fmt.Println("Iterando sobre usuario:")
	for chave, valor := range usuario {
		fmt.Printf("  %s: %s\n", chave, valor)
	}
	fmt.Println()

	// ============================================
	// RESUMO
	// ============================================
	fmt.Println("--- RESUMO ---")
	fmt.Println("Maps em Go:")
	fmt.Println("  ✓ Coleção de pares chave-valor")
	fmt.Println("  ✓ Todos os campos seguem o mesmo tipo [tipoChave]tipoValor")
	fmt.Println("  ✓ Podem ser aninhados (maps dentro de maps)")
	fmt.Println("  ✓ Podem ter elementos adicionados: map[chave] = valor")
	fmt.Println("  ✓ Podem ter elementos removidos: delete(map, chave)")
	fmt.Println("  ✓ Podem ser verificados: valor, existe := map[chave]")
	fmt.Println("  ✓ Podem ser iterados: for chave, valor := range map")
}
