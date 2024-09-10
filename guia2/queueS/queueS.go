package queueS

import (
    "errors"
    "github.com/untref-ayp2/data-structures/stack"
)

// QueueS representa una cola implementada usando dos pilas.
/*
Tipo de estructura que recibe un parámetro de tipo 'T' cualquiera {
    Se declara el campo 'dentro' que es un puntero a una pila 'Stack' que recibe un parámetro de tipo 'T' cualquiera
    Se declara el campo 'fuera' que es un puntero a una pila 'Stack' que recibe un parámetro de tipo 'T' cualquiera
}
*/
type QueueS[T any] struct {
    dentro *stack.Stack[T]  // Pila para encolar elementos
    fuera  *stack.Stack[T]  // Pila para desencolar elementos
}

// NewQueueS crea y devuelve una nueva cola vacía.
/*
Función que utiliza cualquier tipo 'T' y que devuelve un puntero a 'QueueS[T]' {
    Se retorna la dirección de 'QueueS[T] {
        En el campo 'dentro' se llama a la función 'New' de 'stack'
        En el campo 'fuera' se llama a la función 'New' de 'stack'
    }
}
*/
func NewQueueS[T any]() *QueueS[T] {
    return &QueueS[T] {
        dentro: stack.New[T](),
        fuera:  stack.New[T](),
    }
}

// Enqueue agrega un elemento al final de la cola.
/*
Función que utiliza una variable 'q' de tipo *QueueS[T] y que recibe como parámetro una variable 'v' de tipo T {
    Se llama a la función 'Push' desde el campo 'dentro' de la variable 'q', pasándole 'v' como parámetro
}
*/
func (q *QueueS[T]) Enqueue(v T) {
    q.dentro.Push(v)
}

// Dequeue elimina y devuelve el elemento al frente de la cola.
/*
Función que utiliza una variable 'q' de tipo *QueueS[T] y que retorna un T y un error {
    Se declara la variable 'zero' de tipo T sin inicializar
    Si 'q' está vacío {
        Se retornan 'zero' y un error
    }
    Si el campo 'fuera' de 'q' está vacío {
        Para cuando el campo 'dentro' de 'q' no está vacío {
            Se declara la variable 'item' y '_' que son los retornos de la función 'Pop' del campo 'dentro' de 'q'
            Se llama a la función 'Push' del campo 'fuera' de 'q' pasándole 'item' como parámetro
        }
    }
    Se retorna la función 'Pop' del campo 'fuera' de 'q'
}
*/
func (q *QueueS[T]) Dequeue() (T, error) {
    var zero T
    if q.IsEmpty() {
        return zero, errors.New("cola vacía")
    }
    if q.fuera.IsEmpty() {
        for !q.dentro.IsEmpty() {
            item, _ := q.dentro.Pop()
            q.fuera.Push(item)
        }
    }
    return q.fuera.Pop()
}

// Front devuelve el elemento al frente de la cola sin eliminarlo.
/*
Función que utiliza una variable 'q' de tipo *QueueS[T] y que retorna un T y un error {
    Se declara la variable 'zero' de tipo T
    Si 'q' está vacío {
        Se retornan 'zero' y un error
    }
    Si el campo 'fuera' de 'q' está vacío {
        Para cuando el campo 'dentro' de 'q' no está vacío {
            Se declaran las variables 'item' y '_' que son los retornos de la función 'Pop' del campo 'dentro' de 'q'
            Se llama a la función 'Push' del campo 'fuera' de 'q' pasándole 'item' como parámetro
        }
    }
    Se declaran las variables 'item' y 'err' que son los retornos la función 'Pop' del campo 'fuera' de 'q'
    Si 'err' no es nulo {
        Se retornan 'zero' y 'err'
    }
    Se llama a la función 'Push' del campo 'fuera' de 'q' pasándole 'item' como parámetro
    Se retornan 'item' y nil
}
*/
func (q *QueueS[T]) Front() (T, error) {
    var zero T
    if q.IsEmpty() {
        return zero, errors.New("cola vacía")
    }
    if q.fuera.IsEmpty() {
        // Trasladar todos los elementos de 'dentro' a 'fuera' para acceder al elemento del frente
        for !q.dentro.IsEmpty() {
            item, _ := q.dentro.Pop()
            q.fuera.Push(item)
        }
    }
    item, err := q.fuera.Pop()
    if err != nil {
        return zero, err
    }
    q.fuera.Push(item)  // Reinserta el elemento al frente de 'fuera'
    return item, nil
}

// IsEmpty devuelve true si la cola está vacía.
/*
Función que utiliza una variable 'q' de tipo *QueueS[T] y que retorna un bool {
    Se retornan la función 'IsEmpty' de los campos 'dentro' y 'fuera' de 'q'
}
*/
func (q *QueueS[T]) IsEmpty() bool {
    return q.dentro.IsEmpty() && q.fuera.IsEmpty()
}