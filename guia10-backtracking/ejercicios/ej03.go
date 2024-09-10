package ejercicios
//16/05/2024
func Mochila2(objetos []Objeto, capacidad int) (int, []Objeto) {
	mejorValor := 0
	mejor := make([]Objeto, len(objetos))
	actual := make([]Objeto, len(objetos))
	backtrack3(objetos, 0, capacidad, 0, &mejorValor, actual, &mejor)
	return mejorValor, mejor
}

func backtrack3(objetos []Objeto, indice, capacidad, actualValor int, mejorValor *int, actual []Objeto, mejor *[]Objeto) {
	if indice == len(objetos) || capacidad == 0 {
		if actualValor > *mejorValor {
			*mejorValor = actualValor
			*mejor = make([]Objeto, len(actual))
			copy(*mejor, actual)
		}
		return
	}
	backtrack3(objetos, indice + 1, capacidad, actualValor, mejorValor, actual, mejor)
	if objetos[indice].Peso <= capacidad {
		nuevoObjeto := Objeto{Peso: objetos[indice].Peso, Valor: objetos[indice].Valor}
		actual[indice] = nuevoObjeto
		backtrack3(objetos, indice + 1, capacidad - objetos[indice].Peso, actualValor + objetos[indice].Valor, mejorValor, actual, mejor)
		actual[indice] = Objeto{}
	}
}

func Mochila2Punto2(objetos []Objeto, capacidad int) (int, []Objeto) {
	mejorValor := 0
	mejor := make([]Objeto, len(objetos))
	backtrack3Punto2(objetos, 0, capacidad, 0, &mejorValor, []Objeto{}, mejor)
	return mejorValor, mejor
}

func backtrack3Punto2(objetos []Objeto, indice, capacidad, actualValor int, mejorValor *int, seleccionados []Objeto, mejor []Objeto) {
	if indice == len(objetos) || capacidad == 0 {
		if actualValor > *mejorValor {
			*mejorValor = actualValor
			copy(mejor, seleccionados)
		}
		return
	}
	backtrack3Punto2(objetos, indice + 1, capacidad, actualValor, mejorValor, seleccionados, mejor)
	if objetos[indice].Peso <= capacidad {
		objetoSeleccionado := objetos[indice]
		seleccionados = append(seleccionados, objetoSeleccionado)
		backtrack3Punto2(objetos, indice + 1, capacidad - objetoSeleccionado.Peso, actualValor + objetoSeleccionado.Valor, mejorValor, seleccionados, mejor)
	}
}