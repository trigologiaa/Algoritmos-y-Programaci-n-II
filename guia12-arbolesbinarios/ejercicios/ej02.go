package ejercicios

import binarytree "github.com/untref-ayp2/data-structures/binary-tree"

func SumaHojas(bt *binarytree.BinaryTree[int]) int {
	if bt.IsEmpty() {
		return 0
	}
	return sumaHojas(bt.GetRoot())
}

func sumaHojas(n *binarytree.BinaryNode[int]) int {
	if n == nil {
		return 0
	}
	if n.GetLeft() == nil && n.GetRight() == nil {
		return n.GetData()
	}
	return sumaHojas(n.GetLeft()) + sumaHojas(n.GetRight())
}