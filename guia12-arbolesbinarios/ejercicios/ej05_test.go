package ejercicios

import (
	"testing"
	"github.com/stretchr/testify/assert"
	binarytree "github.com/untref-ayp2/data-structures/binary-tree"
)

//	   4
//	  / \
//	 2   5
//	/ \
// 1   3

func TestSumaMayoresQue(t *testing.T) {
	t1 := binarytree.NewBinaryTree(1)
	t3 := binarytree.NewBinaryTree(3)
	t2 := binarytree.NewBinaryTree(2)
	t5 := binarytree.NewBinaryTree(5)
	t4 := binarytree.NewBinaryTree(4)
	t2.InsertLeft(t1)
	t2.InsertRight(t3)
	t4.InsertLeft(t2)
	t4.InsertRight(t5)
	tree := t4
	assert.Equal(t, 14, SumaMayoresQue(tree, 1))
	assert.Equal(t, 9, SumaMayoresQue(tree, 3))
	assert.Equal(t, 0, SumaMayoresQue(tree, 5))
}