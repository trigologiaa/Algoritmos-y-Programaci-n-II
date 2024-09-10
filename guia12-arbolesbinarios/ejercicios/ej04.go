package ejercicios

import binarytree "github.com/untref-ayp2/data-structures/binary-tree"

func SumaPares(bt *binarytree.BinaryTree[int]) int {
	if bt.IsEmpty() {
		return 0
	}
	return sumaPares(bt.GetRoot())
}

func sumaPares(n *binarytree.BinaryNode[int]) int {
	if n == nil {
		return 0
	}
	suma := 0
	if n.GetData()%2 == 0 {
		suma = n.GetData()
	}
	sumaIzquierda := sumaPares(n.GetLeft())
	sumaDerecha := sumaPares(n.GetRight())
	return suma + sumaIzquierda + sumaDerecha
}