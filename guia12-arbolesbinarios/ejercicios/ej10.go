package ejercicios

import bt "github.com/untref-ayp2/data-structures/binary-tree"

func LadronDeCasas(arbol *bt.BinaryTree[int]) int {
	return ladronDeCasas(arbol.GetRoot())
}
//
func ladronDeCasas(nodo *bt.BinaryNode[int]) int {
	if nodo == nil {
		return 0
	}
	var roboIzquierda, roboDerecha int
	if izquierda := nodo.GetLeft(); izquierda != nil {
		roboIzquierda += ladronDeCasas(izquierda.GetLeft()) + ladronDeCasas(izquierda.GetRight())
	}
	if derecha := nodo.GetRight(); derecha != nil {
		roboDerecha += ladronDeCasas(derecha.GetLeft()) + ladronDeCasas(derecha.GetRight())
	}
	roboActual := nodo.GetData()
	return maximo(roboActual + roboIzquierda + roboDerecha, ladronDeCasas(nodo.GetLeft()) + ladronDeCasas(nodo.GetRight()))
}

func maximo(a, b int) int {
	if a > b {
		return a
	}
	return b
}