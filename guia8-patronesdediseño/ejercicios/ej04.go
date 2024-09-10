package ejercicios

import "errors"

type posicionesParesIterator[T int] struct {
	data []T
	posicion int
}

func (p *Array[T]) PosicionesParesIterator() Iterator[T] {
	return &posicionesParesIterator[T]{data: p.data, posicion: 0}
}

func (p *posicionesParesIterator[T]) HasNext() bool {
	return p.posicion < len(p.data)
}

func (p *posicionesParesIterator[T]) Next() (T, error) {
	if p.posicion >= len(p.data) {
		return 0, errors.New("no hay más elementos")
	}
	elemento := p.data[p.posicion]
	p.posicion += 2
	return elemento, nil
}