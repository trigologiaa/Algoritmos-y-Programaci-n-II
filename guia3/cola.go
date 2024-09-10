package guia3

import (
	dll	"github.com/untref-ayp2/data-structures/lists/double_linked_list"
	t 	"github.com/untref-ayp2/data-structures/types"
		"errors"
)

/*
Tipo de estructura que utiliza un tipo T Ordered {
	Se declara la variable 'lista' de tipo *DoubleLinkedList[T]
}
*/
type Cola[T t.Ordered] struct {
	lista *dll.DoubleLinkedList[T]
}

/*
Función que utiliza un tipo T Ordered y retorna un *Cola[T] {
	Se retorna la dirección de Cola[T] {
		Se inicializa el campo 'lista' con el retorno de la función 'NewList' de 'DoubleLinkedList'
	}
}
*/
func NewCola[T t.Ordered]() *Cola[T] {
	return &Cola[T] {
		lista: dll.NewList[T](),
	}
}

/*
Función que utiliza una variable 'c' de tipo *Cola[T] y recibe como parámetro una variable 'dat' de tipo T {
	Se llama a la función 'Append' del campo 'lista' de 'c' pasándole 'dato' como parámetro
}
*/
func (c *Cola[T]) Enqueue(dato T) {
	c.lista.Append(dato)
}

/*
Función que utiliza una variable 'c' de tipo *Cola[T] y retorna un T y un error {
	Si la función 'IsEmpty' de 'c' es true {
		Se declara la variable 'zero' de tipo T
		Se retornan 'zero' y un error
	}
	Se declara la variable 'frontData' de tipo T inicializada con el retorno de la función 'Data' de la función 'Head' del campo 'lista' de 'c'
	Se llama a la función 'RemoveFirst' del campo 'lista' de 'c'
	Se retornan 'frontData' y nil
}
*/
func (c *Cola[T]) Dequeue() (T, error) {
	if c.IsEmpty() {
		var zero T
		return zero, errors.New("cola vacía")
	}
	frontData := c.lista.Head().Data()
	c.lista.RemoveFirst()
	return frontData, nil
}

/*
Función que utiliza una variable 'c' de tipo *Cola[T] y retorna un T y un error {
	Si la función 'IsEmpty' de 'c' es true {
		Se declara la variable 'zero' de tipo T sin inicializar
		Se retornan 'zero' y un error
	}
	Se retornan el retorno de la función 'Data' de la función 'Head' del campo 'lista' de 'c' y nil
}
*/
func (c *Cola[T]) Front() (T, error) {
	if c.IsEmpty() {
		var zero T
		return zero, errors.New("cola vacía")
	}
	return c.lista.Head().Data(), nil
}

/*
Función que utiliza una variable 'c' de tipo *Cola[T] y retorna un bool {
	Se retorna el retorno de la función 'IsEmpty' del campo 'lista' de 'c'
}
*/
func (c *Cola[T]) IsEmpty() bool {
	return c.lista.IsEmpty()
}