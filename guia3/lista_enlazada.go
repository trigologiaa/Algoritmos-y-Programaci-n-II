package guia3

import(
	t	"github.com/untref-ayp2/data-structures/types"
)

// ListaEnlazada es un lista enlazada simple con nodo ficticio
/*
Tipo de estructura que utiliza un tipo T Ordered {
	Se declara el campo 'cabeza' de tipo *Nodo[T]
	Se declara el campo 'cola' de tipo *Nodo[T]
	Se declara el campo 'tamanio' de tipo int
}
*/
type ListaEnlazada[T t.Ordered] struct {
	cabeza  *Nodo[T]
	cola    *Nodo[T]
	tamanio	int
}

// NuevaLista crea una nueva lista enlazada
/*
Función que utiliza un tipo T Ordered y que retorna un *ListaEnlazada[T] {
	Se declara la variable 'dato' de tipo T sin inicializar
	Se declara la variable 'n1' de tipo *Nodo[T] inicializada con el retorno de la función 'NuevoNodo[T]' pasándole 'dato' como parámetro
	Se declara la variable 'n2' de tipo *Nodo[T] inicializada con el retorno de la función 'NuevoNodo[T]' pasándole 'dato' como parámetro
	Se llama a la función 'SetNext' de 'n1' pasándole 'n2' como parámetro
	Se llama a la función 'SetPrev' de 'n2' pasándole 'n1' como parámetro
	Se retorna la dirección de 'ListaEnlazada[T] {
		Se inicializa el campo 'cabeza' con el valor de 'n1'
		Se inicializa el campo 'cola' con el valor de 'n2'
	}
}
*/
func NuevaLista[T t.Ordered]() *ListaEnlazada[T] {
	var dato T
	n1 := NuevoNodo[T](dato)
	n2 := NuevoNodo[T](dato)
	n1.SetNext(n2)
	n2.SetPrev(n1)
	return &ListaEnlazada[T] {
		cabeza:	n1,
		cola: 	n2,
	}
}

// Append agrega un nuevo nodo al final de la lista
/*
Función que utiliza una variable 'l' de tipo *ListaEnlazada[T] y que recibe como parámetro una variable 'dato' de tipo T {
	Se llama a la función 'InsertBefore' del campo 'cola' de 'l' pasándole 'dato' como parámetro
	Se aumenta el valor del campo 'tamanio' de 'l' en 1
}
*/
func (l *ListaEnlazada[T]) Append(dato T) {
	l.cola.InsertBefore(dato)
	l.tamanio++
}

// Prepend agrega un nuevo nodo al principio de la lista
/*
Función que utiliza una variable 'l' de tipo *ListaEnlazada[T] y que recibe como parámetro una variable 'dato' de tipo T {
	Se llama a la función 'InsertAfter' del campo 'cabeza' de 'l' pasándole 'dato' como parámetro
	Se aumenta el valor del campo 'tamanio' de 'l' en 1
}
*/
func (l *ListaEnlazada[T]) Prepend(dato T) {
	l.cabeza.InsertAfter(dato)
	l.tamanio++
}

// RemoveFirst elimina el primer nodo de la lista
/*
Función que utiliza una variable 'l' de tipo *ListaEnlazada[T] {
	Si el llamado a la función 'IsEmpty' de 'l' retorna true {
		Se retorna
	}
	Se llama a la función 'RemoveNext' del campo 'cabeza' de 'l'
	Se disminuye el valor del campo 'tamanio' de 'l' en 1
}
*/
func (l *ListaEnlazada[T]) RemoveFirst() {
	if l.IsEmpty() {
		return
	}
	l.cabeza.RemoveNext()
	l.tamanio--
}

// RemoveLast elimina el último nodo de la lista
/*
Función que utiliza una variable 'l' de tipo *ListaEnlazada[T] {
	Si el llamado a la función 'IsEmpty' de 'l' retorna true {
		Se retorna
	}
	Se llama a la función 'Remove' de la función 'Prev' del campo 'cola' de 'l'
	Se disminuye el valor del campo 'tamanio' en 1
}
*/
func (l *ListaEnlazada[T]) RemoveLast() {
	if l.IsEmpty() {
		return
	}
	l.cola.Prev().Remove()
	l.tamanio--
}

// Remove elimina el nodo con el dato pasado por parámetro
/*
Función que utiliza una variable 'l' de tipo *ListaEnlazada[T] y recibe como parámetro una variable 'dato' de tipo T {
	Se declara la variable 'nodo' de tipo *Nodo[T] inicializada con el retorno de la función 'Next' del campo 'cabeza' de 'l'
	Para cuando 'nodo' no sea igual al campo 'cola' de 'l' {
		Si el retorno de la función 'Data' de 'nodo' es igual a 'dato' {
			Se llama a la función 'Remove' de 'nodo'
			Se disminuye el valor de 'tamanio' en 1
			Se retorna
		}
		Se cambia el valor de 'nodo' por el retorno de la función 'Next' de 'nodo'
	}
}
*/
func (l *ListaEnlazada[T]) Remove(dato T) {
	nodo := l.cabeza.Next()
	for nodo != l.cola {
		if nodo.Data() == dato {
			nodo.Remove()
			l.tamanio--
			return
		}
		nodo = nodo.Next()
	}
}

// IsEmpty devuelve true si la lista está vacía
/*
Función que utiliza una variable 'l' de tipo *ListaEnlazada[T] y retorna un bool {
	Se retorna el resultado bool de la comparación == entre el campo 'tamanio' de 'l' y 0
}
*/
func (l *ListaEnlazada[T]) IsEmpty() bool {
	return l.tamanio == 0
}

// Size devuelve la cantidad de nodos en la lista
/*
Función que utiliza una variable 'l' de tipo *ListaEnlazada[T] y retorna un int {
	Se retorna el campo 'tamanio' de 'l'
}
*/
func (l *ListaEnlazada[T]) Size() int {
	return l.tamanio
}

// Head devuelve el dato del primer nodo de la lista
/*
Función que utiliza una variable 'l' de tipo *ListaEnlazada[T] y retorna un *Nodo[T] {
	Si el llamado a la función 'IsEmpty' de 'l' retorna true {
		Se retorna nil
	}
	Se retorna el retorno del llamado a la función 'Next' del campo 'cabeza' de 'l'
}
*/
func (l *ListaEnlazada[T]) Head() *Nodo[T] {
	if l.IsEmpty() {
		return nil
	}
	return l.cabeza.Next()
}

// Tail devuelve el dato del último nodo de la lista
/*
Función que utiliza una variable 'l' de tipo *ListaEnlazada[T] y retorna un *Nodo[T] {
	Si el llamado a la función 'IsEmpty' de 'l' retorna true {
		Se retorna nil
	}
	Se retorna el retorno del llamado a la función 'Prev' del campo 'cola' de 'l'
}
*/
func (l *ListaEnlazada[T]) Tail() *Nodo[T] {
	if l.IsEmpty() {
		return nil
	}
	return l.cola.Prev()
}

// Find busca un nodo con el dato pasado por parámetro y lo devuelve
/*
Función que utiliza una variable 'l' de tipo *ListaEnlazada[T] que recibe como parámetro una variable 'dato' de tipo T y retorna un *Nodo[T] {
	Se declara la variable 'nodo' de tipo *Nodo[T] inicializada con el retorno de la función 'Next' del campo 'cola' de 'l'
	Para cuando 'nodo' no sea igual al campo 'cola' de 'l' {
		Si el retorno de la función 'Data' de 'nodo' es igual a 'dato' {
			Se retorna 'nodo'
		}
		Se cambia el valor de 'nodo' por el valor que retorna el llamado a la función 'Next' de 'nodo'
	}
	Se retorna nil
}
*/
func (l *ListaEnlazada[T]) Find(dato T) *Nodo[T] {
	nodo := l.cabeza.Next()
	for nodo != l.cola {
		if nodo.Data() == dato {
			return nodo
		}
		nodo = nodo.Next()
	}
	return nil
}

// Clear elimina todos los nodos de la lista
/*
Función que utiliza una variable 'l' de tipo *ListaEnlazada[T] {
	Se llama a la función 'SetNext' del campo 'cabeza' de 'l' pasándole el campo 'cola' de 'l' como parámetro
	Se llama a la función 'SetPrev' del campo 'cola' de 'l' pasándole el campo 'cabeza' de 'l' como parámetro
	Se cambia el valor del campo 'tamanio' de 'l' a 0
}
*/
func (l *ListaEnlazada[T]) Clear() {
	l.cabeza.SetNext(l.cola)
	l.cola.SetPrev(l.cabeza)
	l.tamanio = 0
}