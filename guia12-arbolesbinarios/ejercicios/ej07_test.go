package ejercicios

import (
	"testing"
	"github.com/stretchr/testify/assert"
	binarytree "github.com/untref-ayp2/data-structures/binary-tree"
)

//	   4
//	  / \
//	 2   5
//	/ \   \
// 1   3   6
//        / \
//       8   7

func TestSumaHojasIzquierdas(t *testing.T) {
	t1 := binarytree.NewBinaryTree(1)
	t3 := binarytree.NewBinaryTree(3)
	t2 := binarytree.NewBinaryTree(2)
	t5 := binarytree.NewBinaryTree(5)
	t4 := binarytree.NewBinaryTree(4)
	t6 := binarytree.NewBinaryTree(6)
	t7 := binarytree.NewBinaryTree(7)
	t8 := binarytree.NewBinaryTree(8)
	t2.InsertLeft(t1)
	t2.InsertRight(t3)
	t4.InsertLeft(t2)
	t4.InsertRight(t5)
	t6.InsertRight(t7)
	t6.InsertLeft(t8)
	t5.InsertRight(t6)
	tree := t4
	assert.Equal(t, 9, SumaHojasIzquierdas(tree))
}