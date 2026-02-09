package util

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func CapitalizaPrimeiraLetra(texto string, idioma language.Tag) string {
	return cases.Title(idioma).String(texto)
}
