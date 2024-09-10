package ejercicios

import binarytree "github.com/untref-ayp2/data-structures/binary-tree"

func MayorAltura(t *binarytree.BinaryTree[int]) int {
	if t.IsEmpty() {
		return -1
	}
	return mayorAlturaRecursiva(t.GetRoot()) - 1
}

func mayorAlturaRecursiva(n *binarytree.BinaryNode[int]) int {
	if n == nil {
		return 0
	}
	alturaIzquierda := mayorAlturaRecursiva(n.GetLeft())
	alturaDerecha := mayorAlturaRecursiva(n.GetRight())
	if alturaIzquierda > alturaDerecha {
		return alturaIzquierda + 1
	}
	return alturaDerecha + 1
}