package ejercicios

import tree "github.com/untref-ayp2/data-structures/tree"

func IsBst(bt *tree.BinaryTree[int]) bool {
	if bt.IsEmpty() {
		return true
	}
	return isBstByNode(bt.GetRoot(), nil, nil)
}

func isBstByNode(n *tree.BinaryNode[int], minimo, maximo *int) bool {
	if n == nil {
		return true
	}
	data := n.GetData()
	if (minimo != nil && data <= *minimo) || (maximo != nil && data >= *maximo) {
		return false
	}
	izquierda := n.GetLeft()
	derecha := n.GetRight()
	return isBstByNode(izquierda, minimo, &data) && isBstByNode(derecha, &data, maximo)
}