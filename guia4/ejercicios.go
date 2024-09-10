package guia4

import (
	s	"github.com/untref-ayp2/data-structures/set"
	t	"github.com/untref-ayp2/data-structures/types"
)

/*
Función que utiliza un tipo T Ordered recibiendo como parámetro las variables 'set1' y 'set2' de tipo *SetList[T] y retorna un *SetList[T] {
	Se declara la variable 'resultado' de tipo *SetList[T] inicializado con el retorno de la función 'NewSetList[T]()' de 'SetList'
	Se llama al método 'Add' de 'resultado' pasándole como parámetro la función 'Values' de 'set1'
	Se llama al método 'Add' de 'resultado' pasándole como parámetro la función 'Values' de 'set2'
	Se retorna 'resultado'
}
*/
func Union[T t.Ordered](set1, set2 *s.SetList[T]) *s.SetList[T] {
	resultado := s.NewSetList[T]()
	resultado.Add(set1.Values()...)
	resultado.Add(set2.Values()...)
	return resultado
}

/*
Función que utiliza un tipo T Ordered recibiendo como parámetro las variables 'set1' y 'set2' de tipo *SetList[T] y retorna un *SetList[T] {
	Se declara la variable 'resultado' de tipo *SetList[T] inicializado con el retorno de la función 'NewSetList[T]()' de 'SetList'
	Para cada valor 'v' inicializado con el rango del retorno de la función 'Values' de 'set1' {
		Si el retorno de la función 'Contains' de 'set2' pasándole 'v' como parámetro es true {
			Se llama a la función 'Add' de 'resultado' pasándole 'v' como parámetro
		}
	}
	Se retorna 'resultado'
}
*/
func Intersection[T t.Ordered](set1, set2 *s.SetList[T]) *s.SetList[T] {
	resultado := s.NewSetList[T]()
	for _, v := range set1.Values() {
		if set2.Contains(v) {
			resultado.Add(v)
		}
	}
	return resultado
}

/*
Función que utiliza un tipo T Ordered recibiendo como parámetro las variables 'set1' y 'set2' de tipo *SetList[T] y retorna un *SetList[T] {
	Se declara la variable 'resultado' de tipo *SetList[T] inicializado con el retorno de la función 'NewSetList[T]()' de 'SetList'
	Para cada valor 'v' inicializado con el rango del retorno de la función 'Values' de 'set1' {
		Si el retorno de la función 'Contains' de 'set2' pasándole 'v' como parámetro es false {
			Se llama a la función 'Add' de 'resultado' pasándole 'v' como parámetro
		}
	}
	Se retorna 'resultado'
}
*/
func Difference[T t.Ordered](set1, set2 *s.SetList[T]) *s.SetList[T] {
	resultado := s.NewSetList[T]()
	for _, v := range set1.Values() {
		if !set2.Contains(v) {
			resultado.Add(v)
		}
	}
	return resultado
}

/*
Función que utiliza un tipo T Ordered recibiendo como parámetro las variables 'set1' y 'set2' de tipo *SetList[T] y retorna un bool {
	Para cada valor 'v' inicializado con el rango del retorno de la función 'Values' de 'set2' {
		Si el retorno de la función 'Contains' de 'set1' pasándole 'v' como parámetro es false {
			Se retorna false
		}
	}
	Se retorna true
}
*/
func Subset[T t.Ordered](set1, set2 *s.SetList[T]) bool {
	for _, v := range set2.Values() {
		if !set1.Contains(v) {
			return false
		}
	}
	return true
}

/*
Función que utiliza un tipo T Ordered recibiendo como parámetro las variables 'set1' y 'set2' de tipo *SetList[T] y retorna un bool {
	Si el retorno del llamado a la función 'Size' de 'set1' no es igual al retorno del llamado a la función 'Size' de 'set2' {
		Se retorna false
	}
	Se retorna el retorno del llamado a la función 'Subset' con las variables 'set1' y 'set2' como parámetros
}
*/
func Equal[T t.Ordered](set1, set2 *s.SetList[T]) bool {
	if set1.Size() != set2.Size() {
		return false
	}
	return Subset(set1, set2)
}

/*
Función que utiliza un tipo T Ordered recibiendo como parámetro las variables 'set1' y 'set2' de tipo *SetList[T] y retorna un *SetList[T] {
	Se declara la variable 'resultado' de tipo *SetList[T] inicializado con el retorno de la función 'NewSetList[T]()' de 'SetList'
	Para cada valor 'v' inicializado con el rango del retorno de la función 'Values' de 'set1' {
		Si el retorno de la función 'Contains' de 'set2' pasándole 'v' como parámetro es false {
			Se llama a la función 'Add' de 'resultado' pasándole 'v' como parámetro
		}
	}
	Para cada valor 'v' inicializado con el rango del retorno de la función 'Values' de 'set2' {
		Si el retorno de la función 'Contains' de 'set1' pasándole 'v' como parámetro es false {
			Se llama a la función 'Add' de 'resultado' pasándole 'v' como parámetro
		}
	}
	Se retorna 'resultado'
}
*/
func SymmetricDifference[T t.Ordered](set1, set2 *s.SetList[T]) *s.SetList[T] {
	resultado := s.NewSetList[T]()
	for _, v := range set1.Values() {
		if !set2.Contains(v) {
			resultado.Add(v)
		}
	}
	for _, v := range set2.Values() {
		if !set1.Contains(v) {
			resultado.Add(v)
		}
	}
	return resultado
}

/*
Función que utiliza un tipo T Ordered recibiendo como parámetro una variable 'arreglo' de tipo []T y retorna un []T {
	Se declara la variable 'setUnico' de tipo *SetList[T] inicializada con el retorno del llamado a la función 'NewSetList[T]' que tiene como parámetro 'arreglo...'
	Se retorna el retorno del llamado a la función 'Values' de 'setUnico'
}
*/
func EliminarRepetidos[T t.Ordered](arreglo []T) []T {
	setUnico := s.NewSetList[T](arreglo...)
	return setUnico.Values()
}

/*
Función que utiliza un tipo T Ordered recibiendo como parámetro una variable 'sets...' de tipo *SetList[T] y retorna un *SetList[T] {
	Si el largo de 'sets' es igual a 0 o 1 {
		Se retorna el retorno del llamado a la función 'NewSetList[T]()' de 's'
	}
	Se declara la variable 'resultado' de tipo *SetList[T] inicializada con el primer elemento de 'sets'
	Para cada valor 'set' inicializado con el 'rango' a partir del segundo elemento de 'sets' {
		Se declara la variable 'resultadoTemporal' de tipo *SetList[T] inicializada con el retorno del llamado a la función 'NewSetList[T]()' de 's'
		Para cada valor 'v' del rango del retorno del llamado a la función 'Values' de 'resultado' {
			Si el retorno del llamado a la función 'Contains' de 'set' pasándole 'v' como parámetro es true {
				Se llama a la función 'Add' de 'resultadoTemporal' pasándole 'v' como parámetro
			}
		}
		Se cambia el valor de 'resultado' por el valor de 'resultadoTemporal'
	}
	Se retorna 'resultado'
}
*/
func InterseccionMultiple[T t.Ordered](sets ...*s.SetList[T]) *s.SetList[T] {
    if len(sets) == 0 || len(sets) == 1 {
        return s.NewSetList[T]()
    }
    resultado := sets[0]
    for _, set := range sets[1:] {
        resultadoTemporal := s.NewSetList[T]()
        for _, v := range resultado.Values() {
            if set.Contains(v) {
                resultadoTemporal.Add(v)
            }
        }
        resultado = resultadoTemporal
    }
    return resultado
}