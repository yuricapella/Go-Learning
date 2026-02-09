package main

func sintaxeMapSimples() {
	usuario := map[string]string{
		"nome":      "João",
		"sobrenome": "Silva",
	}

	nome := usuario["nome"]
	_ = nome
}

func sintaxeMapAninhado() {
	usuario := map[string]map[string]string{
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
	}

	primeiroNome := usuario["nome"]["primeiroNome"]
	_ = primeiroNome
}

func sintaxeMapDelete() {
	usuario := map[string]map[string]string{
		"profissao": {
			"nome": "Engenheiro",
		},
	}
	delete(usuario, "profissao")
}

func sintaxeMapAdicionar() {
	usuario := map[string]map[string]string{}
	usuario["telefone"] = map[string]string{
		"ddd":    "11",
		"numero": "1234567890",
	}
}

func sintaxeMapVerificarExistencia() {
	mapa := map[string]string{
		"nome": "João",
	}

	valor, existe := mapa["nome"]
	if existe {
		_ = valor
	}

	valor2, existe2 := mapa["idade"]
	if !existe2 {
		_ = valor2
	}
}

func sintaxeMapMake() {
	usuario3 := make(map[string]string)
	usuario3["nome"] = "Maria"
	usuario3["email"] = "maria@example.com"
}

func sintaxeMapRange() {
	usuario := map[string]string{
		"nome":  "João",
		"idade": "20",
	}

	for chave, valor := range usuario {
		_ = chave
		_ = valor
	}

	for chave := range usuario {
		_ = chave
	}
}
