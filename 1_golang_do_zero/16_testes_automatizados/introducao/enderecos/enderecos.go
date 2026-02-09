package enderecos

import (
	"strings"

	"github.com/yuricapella/Go-Learning/1_golang_do_zero/16_testes_automatizados/introducao/util"
	"golang.org/x/text/language"
)

// TipoDeEndereco verifica se um endereço tem um tipo valido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	//transforma a string em minuscula para nao ter problemas com case sensitive
	enderecoEmLetrasMinusculas := strings.ToLower(endereco)
	//transforma a string em um slice de acordo com os espaços na string
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetrasMinusculas, " ")[0]

	//rua dos bobos
	//["rua","dos","bobos"]

	enderecoTemUmTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		//transforma a primeira letra da string em maiuscula - strings.Title esta depreciada
		return util.CapitalizaPrimeiraLetra(primeiraPalavraDoEndereco, language.BrazilianPortuguese)
	}

	return "Tipo Inválido"
}
