package subsumamax

// Dado un arreglo devuelve la posicion inicial, la posición final y el valor de la subsecuencia cuya suma es máxima dentro del arreglo original.
/*
Función que recibe como parámetro una variable 'arreglo' de tipo []int y devuelve tres int, que son la suma máxima, la posición inicial y la posición final {
	Se declara la variable 'sumaMaxima' de tipo int inicializada en 0
	Se declaran las variables 'posInicial' y 'pos Final' de tipo int inicializadas en -1
	Para cuando la variable 'i' inicializada en 0 sea menor al largo de 'arreglo' {
		Se declara la variable 'sumaLocal' de tipo int inicializada en 0
		Para cuando la variable 'j' inicializada en el mismo valor que 'i' sea menor al largo de 'arreglo' {
			Se suma el valor de la posición 'j' de 'arreglo' a 'sumaLocal'
			Si 'sumaLocal' es mayor a 'sumaMaxima' {
				Se le cambia el valor de 'sumaMaxima' por el de 'sumaLocal'
				Se le cambia el valor de 'posInicial' por el valor de 'i'
				Se le cambia el valor de 'posFinal' por el valor de 'j'
			}
		}
	}
	Se retornan 'sumaMaxima', 'posInicial' y 'posFinal'
}
*/
func SubsecuenciaSumaMaxima(arreglo []int) (int, int, int) {
	sumaMaxima := 0
	posInicial, posFinal := -1, -1
	for i := 0; i < len(arreglo); i++ {
		sumaLocal := 0
		for j := i; j < len(arreglo); j++ {
			sumaLocal += arreglo[j]
			if sumaLocal > sumaMaxima {
				sumaMaxima = sumaLocal
				posInicial = i
				posFinal = j
			}
		}
	}
	return sumaMaxima, posInicial, posFinal
}