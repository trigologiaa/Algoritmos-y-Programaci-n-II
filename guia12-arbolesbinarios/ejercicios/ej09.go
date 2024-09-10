package ejercicios

import "github.com/untref-ayp2/data-structures/binary-tree"

func RecorrerPorNiveles(tree *binarytree.BinaryTree[string]) []string {
	if tree.IsEmpty() {
		return nil
	}
	var resultado []string
	cola := []*binarytree.BinaryNode[string]{tree.GetRoot()}
	for len(cola) > 0 {
		nodoActual := cola[0]
		cola = cola[1:]
		resultado = append(resultado, nodoActual.GetData())
		if izquierda := nodoActual.GetLeft(); izquierda != nil {
			cola = append(cola, izquierda)
		}
		if derecha := nodoActual.GetRight(); derecha != nil {
			cola = append(cola, derecha)
		}
	}
	return resultado
}