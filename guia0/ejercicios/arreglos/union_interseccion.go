package arreglos

/*
Función que recibe las variables 'lista1' y 'lista2' de tipo []int como parámetro y devuelve dos []int que serían la unión y la intersección {
	Se declara la variable 'union' de tipo []int inicialmente vacío
	Se declara la variable 'interseccion' de tipo []int inicialmente vacío
	Se declara la variable 'existeEnUnion' de tipo 'func' que recibe la variable 'valor' de tipo int y la variable lista de tipo []int como parámetros y retorna bool {
		Para cada valor de 'lista' inicializado en la variable 'v' {
			Si 'v' es igual a 'valor' {
				Se retorna true
			}
		}
		Se retorna false
	}
	Para cada valor de 'lista1' inicializado en la variable 'v' {
		Si el valor 'v' no existe en 'union' {
			Se agrega el valor 'v' a 'union'
		}
	}
	Para cada valor de 'lista2' inicializado en la variable 'v' {
		Si el valor 'v' existe en 'lista1' {
			Si el valor 'v' no existe en 'interseccion' {
				Se agrega el valor 'v' a 'interseccion'
			}
		}
		Si el valor 'v' no existe en 'union' {
			Se agrega el valor 'v' a 'union'
		}
	}
	Se retornan 'union' e 'interserccion'
}
*/
func UnionInterseccion(lista1, lista2 []int) ([]int, []int) {
	union := []int {}
	interseccion := []int {}
	existeEnUnion := func(valor int, lista []int) bool {
		for _, v := range lista {
			if v == valor {
				return true
			}
		}
		return false
	}
	for _, v := range lista1 {
		if !existeEnUnion(v, union) {
			union = append(union, v)
		}
	}
	for _, v := range lista2 {
		if existeEnUnion(v, lista1) {
			if !existeEnUnion(v, interseccion) {
				interseccion = append(interseccion, v)
			}
		}
		if !existeEnUnion(v, union) {
			union = append(union, v)
		}
	}
	return union, interseccion
}