package guia3

import (
        "errors"
    sll "github.com/untref-ayp2/data-structures/lists/single_linked_list"
    t   "github.com/untref-ayp2/data-structures/types"
)

/*
Tipo de estructura que utiliza un tipo T Ordered {
    Se declara el campo 'lista' de tipo *SingleLinkedList[T]
}
*/
type Pila[T t.Ordered] struct {
    lista *sll.SingleLinkedList[T]
}

/*
Función que utiliza un tipo T Ordered y retorna un *Pila[T] {
    Se retorna la dirección de Pila[T] {
        Se inicializa el campo 'lista' como el retorno de la función 'NewList[T]' de SingleLinkedList[T]
    }
}
*/
func NewPila[T t.Ordered]() *Pila[T] {
    return &Pila[T] {
        lista: sll.NewList[T](),
    }
}

/*
Función que utiliza la variable 'p' de tipo *Pila[T] y que recibe como parámetro una variable 'dato' de tipo T {
    Se llama a la función 'Prepend' del campo 'lista' de 'p' pasándole 'dato' como parámetro
}
*/
func (p *Pila[T]) Push(dato T) {
    p.lista.Prepend(dato)
}

/*
Función que utiliza la variable 'p' de tipo *Pila[T] y que retorna un T y un error {
    Si la función 'IsEmpty' de 'p' es true {
        Se declara la variable 'zero' de tipo T
        Se retornan 'zero' y un error
    }
    Se declara la función 'topNode' de tipo *Node[T] inicializada con el retorno de la función 'Head' del campo 'lista' de 'p'
    Se llama a la función 'RemoveFirst' del campo 'lista' de 'p'
    Se retornan la función 'Data' de la variable 'topNode' y nil
}
*/
func (p *Pila[T]) Pop() (T, error) {
    if p.IsEmpty() {
        var zero T
        return zero, errors.New("pila vacía")
    }
    topNode := p.lista.Head()
    p.lista.RemoveFirst()
    return topNode.Data(), nil
}

/*
Función que utiliza la variable 'p' de tipo *Pila[T] y que retorna un T y un error {
    Si la función 'IsEmpty' de 'p' es true {
        Se declara la variable 'zero' de tipo T
        Se retornan 'zero' y un error
    }
    Se retornan la función 'Data' de la función 'Head' del campo 'Lista' de 'p' y nil
}
*/
func (p *Pila[T]) Top() (T, error) {
    if p.IsEmpty() {
        var zero T
        return zero, errors.New("pila vacía")
    }
    return p.lista.Head().Data(), nil
}

/*
Función que utiliza la variable 'p' de tipo *Pila[T] y que retorna un bool {
    Se retorna el retorno de la función 'IsEmpty' del campo 'lista' de 'p'
}
*/
func (p *Pila[T]) IsEmpty() bool {
    return p.lista.IsEmpty()
}