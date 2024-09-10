package ejercicios

import (
	"github.com/untref-ayp2/data-structures/tree"
)

func SumaHojasMenoresQue10(bt *tree.BinarySearchTree[int]) int {
	if bt.IsEmpty() {
		return 0
	}
	return sumaHojasMenoresQue10Rec(bt.GetRoot())
}

func sumaHojasMenoresQue10Rec(n *tree.BinaryNode[int]) int {
	if n == nil {
		return 0
	}
	if n.GetLeft() == nil && n.GetRight() == nil {
		if n.GetData() < 10 {
			return n.GetData()
		} else {
			return 0
		}
	}
	return sumaHojasMenoresQue10Rec(n.GetLeft()) + sumaHojasMenoresQue10Rec(n.GetRight())
}