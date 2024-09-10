package guia3

import(
    t   "github.com/untref-ayp2/data-structures/types"
)

/*
Tipo de estructura que utiliza un tipo T Ordered {
    Se declara el campo 'dato' de tipo T
    Se declara el campo 'sig' de tipo *Nodo[T]
    Se declara el campo 'prev' de tipo *Nodo[T]
}
*/
type Nodo[T t.Ordered] struct {
    dato    T
    sig     *Nodo[T]
    prev    *Nodo[T]
}

/*
Función que utiliza un tipo T Ordered, que recibe una variable 'dato' de tipo T y que retorna un *Nodo[T] {
    Se retorna la dirección de Nodo[T] {
        Se inicializa el campo 'dato' con el parámetro 'dato'
        Se inicializa el campo 'sig' en nil
        Se inicializa el campo 'prev' en nil
    }
}
*/
func NuevoNodo[T t.Ordered](dato T) *Nodo[T] {
    return &Nodo[T] {
        dato:   dato,
        sig:    nil,
        prev:   nil,
    }
}

/*
Función que utiliza una variable 'n' de tipo *Nodo[T] y que retorna un T {
    Se retorna el campo 'dato' de 'n'
}
*/
func (n *Nodo[T]) Data() T {
    return n.dato
}

/*
Función que utiliza una variable 'n' de tipo *Nodo[T] y que retorna un *Nodo[T] {
    Se retorna el campo 'sig' de 'n'
}
*/
func (n *Nodo[T]) Next() *Nodo[T] {
    return n.sig
}

/*
Función que utiliza una variable 'n' de tipo *Nodo[T] y que retorna un *Nodo[T] {
    Se retorna el campo 'prev' de 'n'
}
*/
func (n *Nodo[T]) Prev() *Nodo[T] {
    return n.prev
}

/*
Función que utiliza una variable 'n' de tipo *Nodo[T] y que recibe como parámetro una variable 'dato' de tipo T {
    Se cambia el valor del campo 'dato' de 'n' por el valor del parámetro dato
}
*/
func (n *Nodo[T]) SetData(dato T) {
    n.dato = dato
}

/*
Función que utiliza una variable 'n' de tipo *Nodo[T] y que recibe como parámetro una variable 'sig' de tipo *Nodo[T] {
    Se cambia el valor del campo 'sig' de 'n' por el valor del parámetro 'sig'
}
*/
func (n *Nodo[T]) SetNext(sig *Nodo[T]) {
    n.sig = sig
}

/*
Función que utiliza una variable 'n' de tipo *Nodo[T] y que recibe como parámetro una variable 'prev' de tipo *Nodo[T] {
    Se cambia el valor del campo 'prev' de 'n' por el valor del parámetro 'prev'
}
*/
func (n *Nodo[T]) SetPrev(prev *Nodo[T]) {
    n.prev = prev
}

/*
Función que utiliza una variable 'n' de tipo *Nodo[T] y que recibe como parámetro una variable 'dato' de tipo T {
    Se declara la variable 'newNode' de tipo *Nodo[T] inicializada con el retorno de la función 'NuevoNodo' pasándole 'dato' como parámetro
    Se cambia el valor del campo 'sig' de 'newNode' por el valor del campo 'sig' de 'n'
    Se cambia el valor del campo 'prev' de 'newNode' por el valor de 'n'
    Si el campo 'sig' de 'n' no es nil {
        Se cambia el valor del campo 'prev' del campo 'sig' de 'n' por el valor de 'newNode'
    }
    Se cambia el valor del campo 'sig' de 'n' por el valor de 'newNode'
}
*/
func (n *Nodo[T]) InsertAfter(dato T) {
    newNode := NuevoNodo(dato)
    newNode.sig = n.sig
    newNode.prev = n
    if n.sig != nil {
        n.sig.prev = newNode
    }
    n.sig = newNode
}

/*
Función que utiliza una variable 'n' de tipo *Nodo[T] y que recibe como parámetro una variable 'dato' de tipo T {
    Se declara la variable 'newNode' de tipo *Nodo[T] inicializada con el retorno de la función 'NuevoNodo' pasándole 'dato' como parámetro
     Se cambia el valor del campo 'prev' de 'newNode' por el valor del campo 'prev' de 'n'
     Se cambia el valor del campo 'sig' de 'newNode' por el valor de 'n'
     Si el campo 'prev' de 'n' no es nil {
        Se cambia el valor del campo 'sig' del campo 'prev' de 'n' por el valor de 'newNode'
    }
    Se cambia el valor del campo 'prev' de 'n' por el valor de 'newNode'
}
*/
func (n *Nodo[T]) InsertBefore(dato T) {
    newNode := NuevoNodo(dato)
    newNode.prev = n.prev
    newNode.sig = n
    if n.prev != nil {
        n.prev.sig = newNode
    }
    n.prev = newNode
}

/*
Función que utiliza una variable 'n' de tipo *Nodo[T] {
    Si el valor del campo 'prev' de 'n' no es nil {
        Se cambia el valor del campo 'sig' del campo 'prev' de 'n' por el valor del campo 'sig' de 'n'
    }
    Si el valor del campo 'sig' de 'n' no es nil {
        Se cambia el valor del campo 'prev' del campo 'sig' de 'n' por el valor del campo 'prev' de 'n'
    }
}
*/
func (n *Nodo[T]) Remove() {
    if n.prev != nil {
        n.prev.sig = n.sig
    }
    if n.sig != nil {
        n.sig.prev = n.prev
    }
}

/*
Función que utiliza una variable 'n' de tipo *Nodo[T] {
    Si el valor del campo 'sig' de 'n' no es nil {
        Se llama a la función 'Remove' del campo 'sig' de 'n'
    }
}
*/
func (n *Nodo[T]) RemoveNext() {
    if n.sig != nil {
        n.sig.Remove()
    }
}

/*
Función que utiliza una variable 'n' de tipo *Nodo[T] {
    Si el valor del campo 'priv' de 'n' no es nil {
        Se llama a la función 'Remove' del campo 'prev' de 'n'
    }
}
*/
func (n *Nodo[T]) RemovePrev() {
    if n.prev != nil {
        n.prev.Remove()
    }
}