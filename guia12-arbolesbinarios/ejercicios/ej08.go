package ejercicios

import binarytree "github.com/untref-ayp2/data-structures/binary-tree"

func SumaDerechosImpares(bt *binarytree.BinaryTree[int]) int {
	if bt.IsEmpty() {
		return 0
	}
	return sumaDerechosImpares(bt.GetRoot())
}

func sumaDerechosImpares(n *binarytree.BinaryNode[int]) int {
	if n == nil {
		return 0
	}
	suma := 0
	if n.GetRight() != nil && n.GetRight().GetData()%2 != 0 {
		suma += n.GetRight().GetData()
	}
	suma += sumaDerechosImpares(n.GetLeft()) + sumaDerechosImpares(n.GetRight())
	return suma
}