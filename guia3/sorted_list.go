package guia3

import (
	dll	"github.com/untref-ayp2/data-structures/lists/double_linked_list"
	t	"github.com/untref-ayp2/data-structures/types"
)

/*
Tipo de estructura que utiliza un tipo T Ordered {
	Se declara el campo 'lista' de tipo *DoubleLinkedList[T]
}
*/
type SortedList[T t.Ordered] struct {
	lista *dll.DoubleLinkedList[T]
}

/*
Función que utiliza un tipo T Ordered y retorna un *SortedList[T] {
	Se retorna la dirección de SortedList[T] {
		Se llama a la función 'NewList[T]' de 'DoubleLinkedList[T]'
	}
}
*/
func NewList[T t.Ordered]() *SortedList[T] {
	return &SortedList[T] {
		dll.NewList[T](),
	}
}

/*
Función que utiliza una variable 's' de tipo *SortedList[T] y recibe como parámetro una variable 'dato' de tipo T {
	Para cada caso {
	En caso de que el llamado a la función 'IsEmpty' del campo 'lista' de 's' retorne true:
		Se llama a la función 'Append' del campo 'lista' de 's' pasándole 'dato' como parámetro
	En caso de que el retorno de la función 'Data' de la función 'Head' del campo 'lista' de 's' sea mayor a 'dato':
		Se llama a la función 'Prepend' del campo 'lista' de 's' pasándole 'dato' como parámetro
	En caso de que el retorno de la función 'Data' de la función 'Tail' del campo 'lista' de 's' sea menor a 'dato':
		Se llama a la función 'Append' del campo 'lista' de 's' pasándole 'dato' como parámetro
	En otro caso:
		Se declara la variable 'nodo' de tipo *Node[T] inicializada con el retorno de la función 'Head' del campo 'lista' de 's'
		Para cuando el retorno de la función 'Data' de la función 'Next' de 'nodo' sea menor a 'dato' {
			Se cambia el valor de 'nodo' por el valor del retorno de la función 'Next' de 'nodo'
		}
		Se declara la variable 'siguiente' de tipo *Node[T] inicializada con el retorno de la función 'Next' de 'nodo'
		Se declara la variable 'nuevoNodo' de tipo *Node[T] inicializada con el retorno de la función 'NewNode' de 'DoubleLinkedList' pasándole 'dato' como parámetro
		Se llama a la función 'SetNext' de 'nodo' pasándole 'nuevoNodo' como parámetro
		Se llama a la función 'SetNext' de 'nuevoNodo' pasándole 'siguiente' como parámetro
		Se llama a la función 'SetPrev' de 'siguiente' pasándole 'nuevoNodo' como parámetro
		Se llama a la función 'SetPrev' de 'nuevoNodo' pasándole 'nodo' como parámetro
	}
}
*/
func (s *SortedList[T]) Insert(dato T) {
	switch {
	case s.lista.IsEmpty():
		s.lista.Append(dato)
	case s.lista.Head().Data() > dato:
		s.lista.Prepend(dato)
	case s.lista.Tail().Data() < dato:
		s.lista.Append(dato)
	default:
		nodo := s.lista.Head()
		for nodo.Next().Data() < dato {
			nodo = nodo.Next()
		}
		siguiente := nodo.Next()
		nuevoNodo := dll.NewNode(dato)
		nodo.SetNext(nuevoNodo)
		nuevoNodo.SetNext(siguiente)
		siguiente.SetPrev(nuevoNodo)
		nuevoNodo.SetPrev(nodo)
	}
}

/*
Función que utiliza una variable 's' de tipo *SortedList[T] {
	Se llama a la función 'RemoveFirst' del campo 'lista' de 's'
}
*/
func (s *SortedList[T]) RemoveFirst() {
	s.lista.RemoveFirst()
}

/*
Función que utiliza una variable 's' de tipo *SortedList[T] {
	Se llama a la función 'RemoveLast' del campo 'lista' de 's'
}
*/
func (s *SortedList[T]) RemoveLast() {
	s.lista.RemoveLast()
}

/*
Función que utiliza una variable 's' de tipo *SortedList[T] y recibe como parámetro una variable 'data' de tipo T {
	Se llama a la función 'Remove' del campo 'lista' de 's' pasándole 'data' como parámetro
}
*/
func (s *SortedList[T]) Remove(data T) {
	s.lista.Remove(data)
}

/*
Función que utiliza una variable 's' de tipo *SortedList[T] y que retorna un bool {
	Se retorna el retorno de la función 'IsEmpty' del campo 'lista' de 's'
}
*/
func (s *SortedList[T]) IsEmpty() bool {
	return s.lista.IsEmpty()
}

/*
Función que utiliza una variable 's' de tipo *SortedList[T] y que retorna un int {
	Se retorna el retorno de la función 'Size' del campo 'lista' de 's'
}
*/
func (s *SortedList[T]) Size() int {
	return s.lista.Size()
}

/*
Función que utiliza una variable 's' de tipo *SortedList[T] y que retorna un T {
	Se retorna el retorno de la función 'Data' de la función 'Head' del campo 'lista' de 's'
}
*/
func (s *SortedList[T]) Head() T {
	return s.lista.Head().Data()
}

/*
Función que utiliza una variable 's' de tipo *SortedList[T] y que retorna un T {
	Se retorna el retorno de la función 'Data' de la función 'Tail' del campo 'lista' de 's'
}
*/
func (s *SortedList[T]) Tail() T {
	return s.lista.Tail().Data()
}

/*
Función que utiliza una variable 's' de tipo *SortedList[T], que recibe una variable 'data' de tipo T y que retorna un *Node[T] {
	Se retorna el retorno de la función 'Find' del campo 'lista' de 's' pasándole 'data' como parámetro
}
*/
func (s *SortedList[T]) Find(data T) *dll.Node[T] {
	return s.lista.Find(data)
}

/*
Función que utiliza una variable 's' de tipo *SortedList[T] {
	Se llama a la función 'Clear' del campo 'lista' de 's'
}
*/
func (s *SortedList[T]) Clear() {
	s.lista.Clear()
}