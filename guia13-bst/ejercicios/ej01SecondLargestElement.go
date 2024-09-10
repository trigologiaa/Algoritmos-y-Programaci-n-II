package ejercicios

import (
	"errors"
	tree "github.com/untref-ayp2/data-structures/tree"
)

func SecondLargestElement(bt *tree.BinarySearchTree[int]) (int, error) {
	var x int
	if bt.Size() < 2 {
		return x, errors.New("no hay valores")
	}
	return findSecLargeElem(bt.GetRoot()), nil
}

func findSecLargeElem(n *tree.BinaryNode[int]) int {
	var x int
	actual := n
	for actual.GetRight() != nil {
		x = actual.GetData()
		actual = actual.GetRight()
	}
	if actual.GetLeft() != nil {
		actual = actual.GetLeft()
		for actual.GetRight() != nil {
			actual = actual.GetRight()
		}
		x = actual.GetData()
	}
	return x
}