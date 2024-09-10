package ejercicios
//16/05/2024
func Mochila01(objetos []Objeto, capacidad int) int {
	mejorValor := 0
	backtrack2(objetos, capacidad, 0, 0, &mejorValor)
	return mejorValor
}

func backtrack2(objetos []Objeto, capacidad, indice, valorActual int, mejorValor *int) {
	if indice == len(objetos) || capacidad == 0 {
		if valorActual > *mejorValor {
			*mejorValor = valorActual
		}
		return
	}
	backtrack2(objetos, capacidad, indice + 1, valorActual, mejorValor)
	if objetos[indice].Peso <= capacidad {
		incluido := valorActual + objetos[indice].Valor
		backtrack2(objetos, capacidad - objetos[indice].Peso, indice + 1, incluido, mejorValor)
	}
}