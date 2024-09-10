package ejercicios

import (
	"errors"
	tree "github.com/untref-ayp2/data-structures/tree"
)

func PredecesorInOrder(bt *tree.BinarySearchTree[int], k int) (int, error) {
	var x int
	iterador := bt.Iterator()
	if bt.Size() == 0 {
		return x, errors.New("no hay predecesores")
	}
	if k <= bt.FindMin().GetData() {
		return x, errors.New("no hay predecesores menores que el mínimo")
	}
	for iterador.HasNext() {
		siguiente := iterador.Next()
		if siguiente < k {
			x = siguiente
		}
		if siguiente >= k {
			return x, nil
		}
	}
	return x, nil
}