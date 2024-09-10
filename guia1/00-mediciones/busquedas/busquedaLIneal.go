package busquedas

// BusLineal busca un elemento en un arreglo de enteros usando el algoritmo de búsqueda lineal
/*
Función que recibe como parámetros una variable 'datos' de tipo []int y una variable 'buscado' de tipo int, la cual retorna un int {
	Para cuando esté en la posición 'i' con el valor 'v' de 'datos' {
		Si 'buscado' es igual al valor 'v' de 'datos' {
			Se retorna la posición 'i' de 'datos'
		}
	}
	Se retorna -1
}
*/
func BusLineal(datos []int, buscado int) int {
	for i, v := range datos {
		if buscado == v {
			return i
		}
	}
	return -1
}