package ejercicios

import binarytree "github.com/untref-ayp2/data-structures/binary-tree"

func SumaHojasIzquierdas(t *binarytree.BinaryTree[int]) int {
	if t.IsEmpty() {
		return 0
	}
	return sumaHojasIzquierdas(t.GetRoot(), false)
}

func sumaHojasIzquierdas(n *binarytree.BinaryNode[int], esHijoIzquierdo bool) int {
	if n == nil {
		return 0
	}
	if n.GetLeft() == nil && n.GetRight() == nil && esHijoIzquierdo {
		return n.GetData()
	}
	return sumaHojasIzquierdas(n.GetLeft(), true) + sumaHojasIzquierdas(n.GetRight(), false)
}