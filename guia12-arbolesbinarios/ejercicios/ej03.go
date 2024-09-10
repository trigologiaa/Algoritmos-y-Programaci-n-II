package ejercicios

import binarytree "github.com/untref-ayp2/data-structures/binary-tree"

func SumaInternos(bt *binarytree.BinaryTree[int]) int {
	if bt.IsEmpty() {
		return 0
	}
	raiz := bt.GetRoot().GetData()
	var raizNodo int = int(raiz)
	return sumaInternos(bt.GetRoot()) - raizNodo
}

func sumaInternos(n *binarytree.BinaryNode[int]) int {
	if n == nil || (n.GetLeft() == nil && n.GetRight() == nil) {
		return 0
	}
	sumaIzquierda := sumaInternos(n.GetLeft())
	sumaDerecha := sumaInternos(n.GetRight())
	return n.GetData() + sumaIzquierda + sumaDerecha
}