package ejercicios

import "errors"

type ElementosParesIterator[T int] struct {
	data  []T
	posicion int
}

func (p *Array[T]) ElementosParesIterator() Iterator[T] {
	return &ElementosParesIterator[T]{data: p.data, posicion: 1}
}

func (p *ElementosParesIterator[T]) HasNext() bool {
	return p.posicion < len(p.data)
}

func (p *ElementosParesIterator[T]) Next() (T, error) {
	if p.posicion >= len(p.data) {
		return 0, errors.New("no hay más elementos")
	}
	elemento := p.data[p.posicion]
	p.posicion += 2
	return elemento, nil
}