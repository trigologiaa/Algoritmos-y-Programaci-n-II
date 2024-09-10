package ejercicios

import binarytree "github.com/untref-ayp2/data-structures/binary-tree"

func SumaMayoresQue(bt *binarytree.BinaryTree[int], k int) int {
	if bt.IsEmpty() {
		return 0
	}
	return sumaMayoresQue(bt.GetRoot(), k)
}

func sumaMayoresQue(n *binarytree.BinaryNode[int], k int) int {
	if n == nil {
		return 0
	}
	sumaIzquierda := sumaMayoresQue(n.GetLeft(), k)
	sumaDerecha := sumaMayoresQue(n.GetRight(), k)
	if n.GetData() > k {
		return n.GetData() + sumaIzquierda + sumaDerecha
	}
	return sumaIzquierda + sumaDerecha
}