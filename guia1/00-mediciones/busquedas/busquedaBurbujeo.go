package busquedas

// BusquedaBurbujeo busca un elemento en un arreglo de enteros usando el algoritmo de búsqueda por burbujeo
/*
Función que recibe como parámetros una variable 'datos' de tipo []int y una variable 'buscado' de tipo int, la cual retorna un int {
    Se declara la variable 'n' de tipo int inicializada como el largo de 'datos'
    Para cuando 'i' inicializado en 0 sea menor a 'n' {
        Si la posición 'i' de 'datos' es igual a 'buscado' {
            Para cuando 'j' inicializado en el mismo valor que 'i' sea menor a 'n' - 1 {
                Se cambian el valor de la posición 'j' de 'datos' y el valor 'j' + 1 de 'datos' entre sí
            }
            Se retorna 'n' - 1
        }
    }
    Se retorna -1
}
*/
func BusquedaBurbujeo(datos []int, buscado int) int {
    n := len(datos)
    for i := 0; i < n; i++ {
        if datos[i] == buscado {
            for j := i; j < n - 1; j++ {
                datos[j], datos[j + 1] = datos[j + 1], datos[j]
            }
            return n - 1
        }
    }
    return -1
}