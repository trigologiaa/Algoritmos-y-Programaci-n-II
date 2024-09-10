package busquedas

// BusquedaBinaria busca un elemento en un arreglo de enteros usando el algoritmo de búsqueda binaria
/*
Función que recibe como parámetros una variable 'datos' de tipo []int y una variable 'buscado' de tipo int, la cual retorna un int {
	Se declara la variable 'inicio' de tipo int inicializada en 0
	Se declara la variable 'fin' de tipo int la cual se inicializa como la última posición de 'datos'
	Para cuando 'inicio' sea menor o igual a 'fin' {
		Se declara la variable 'medio' de tipo int inicializada como la división entre 2 de la suma de 'inicio' y 'fin', siendo esto la mitad de 'datos'
		Para cada caso {
			En caso de que 'buscado' sea menor al medio de 'datos':
				Se cambia el valor de 'fin' por el de 'medio' - 1
			En caso de que 'buscado' sea mayor al medio de 'datos':
				Se cambia el valor de 'inicio' por el de 'medio' + 1
			En otro caso:
				Se retorna 'medio'
		}
	}
	Se retorna -1
}
*/
func BusquedaBinaria(datos []int, buscado int) int {
	inicio := 0
	fin := len(datos) - 1
	for inicio <= fin {
		medio := (inicio + fin) / 2
		switch {
		case buscado < datos[medio]:
			fin = medio - 1
		case buscado > datos[medio]:
			inicio = medio + 1
		default:
			return medio
		}
	}
	return -1
}