package funciones

/*
	Función que recibe una lista de tipo []int y devuelve un int mínimo y un int máximo {
		Si la lista está vacía {
			Se retorna 0 y 0
		}
		Se declara la variable 'minimo' que es igual al primer elemento de 'lista'
		Se declara la variable 'maximo' que es igual al primer elemento de 'lista'
		Por cada valor de 'lista' {
			Si 'minimo' es mayor al valor 'v' indicado {
				Se actualiza el valor de 'minimo' al valor actual de 'v'
			}
			Si 'maximo' es menor al valor 'v' indicado {
				Se actualiza el valor de 'maximo' al valor actual de 'v'
			}
		}
		Se retornan 'minimo' y 'maximo'
	}
*/
func EncontrarMinimoMaximo(lista []int) (int, int) {
	if len(lista) == 0 {
		return 0, 0
	}
	minimo := lista[0]
	maximo := lista[0]
	for _, v := range lista {
		if minimo > v {
			minimo = v
		}
		if maximo < v {
			maximo = v
		}
	}
	return minimo, maximo
}