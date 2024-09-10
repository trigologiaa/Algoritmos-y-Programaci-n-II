package guia6

import (
	"strings"

	"github.com/agrison/go-commons-lang/stringUtils"
	"github.com/untref-ayp2/data-structures/dictionary"
	"github.com/untref-ayp2/data-structures/list"
)

func Traducir(texto string, dict dictionary.Dictionary[string, string]) string {
	palabras := strings.Split(texto, " ")
	var traduccion []string
	for _, palabra := range palabras {
		traducida := dict.Get(palabra)
		if traducida == "" {
			traduccion = append(traduccion, "error")
		} else {
			traduccion = append(traduccion, traducida)
		}
	}
	return strings.Join(traduccion, " ")
}

func Frecuencia(texto string) *dictionary.Dictionary[string, int] {
	freq := dictionary.NewDictionary[string, int]()
	for _, r := range texto {
		char := string(r)
		if !stringUtils.IsWhitespace(char) {
			currentCount := freq.Get(char)
			freq.Put(char, currentCount+1)
		}
	}
	return freq
}

func Interseccion(s1 []string, s2 []string) *list.LinkedList[string] {
	lista := list.NewLinkedList[string]()
	d := dictionary.NewDictionary[string, bool]()
	for _, s := range s1 {
		d.Put(s, true)
	}
	for _, s := range s2 {
		if d.Contains(s) {
			lista.Append(s)
		}
	}
	return lista
}

func InformacionSolicitada(entrada dictionary.Dictionary[string, []string]) *dictionary.Dictionary[string, []string] {
	retorno := dictionary.NewDictionary[string, []string]()
	for _, dia := range entrada.Keys() {
		alumnosDelDia := entrada.Get(dia)
		for _, alumno := range alumnosDelDia {
			var diasPorAlumno []string
			if retorno.Contains(alumno) {
				diasPorAlumno = retorno.Get(alumno)
			} else {
				diasPorAlumno = make([]string, 0)
			}
			diasPorAlumno = append(diasPorAlumno, dia)
			retorno.Put(alumno, diasPorAlumno)
		}
	}
	return retorno
}
