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
//          \
//           7

func TestMayorAltura(t *testing.T) {
	t1 := binarytree.NewBinaryTree(1)
	t3 := binarytree.NewBinaryTree(3)
	t2 := binarytree.NewBinaryTree(2)
	t5 := binarytree.NewBinaryTree(5)
	t4 := binarytree.NewBinaryTree(4)
	t6 := binarytree.NewBinaryTree(6)
	t7 := binarytree.NewBinaryTree(7)
	t2.InsertLeft(t1)
	t2.InsertRight(t3)
	t4.InsertLeft(t2)
	t4.InsertRight(t5)
	t6.InsertRight(t7)
	t5.InsertRight(t6)
	tree := t4
	assert.Equal(t, 3, MayorAltura(tree))
}

func TestMayorAltura2(t *testing.T) {
	t20 := binarytree.NewBinaryTree(20)
	t19 := binarytree.NewBinaryTree(19)
	t1 := binarytree.NewBinaryTree(1)
	t29 := binarytree.NewBinaryTree(29)
	t27 := binarytree.NewBinaryTree(27)
	t54 := binarytree.NewBinaryTree(54)
	t13 := binarytree.NewBinaryTree(13)
	t15 := binarytree.NewBinaryTree(15)
	t20.InsertLeft(t19)
	t20.InsertRight(t29)
	t19.InsertLeft(t1)
	t29.InsertLeft(t27)
	t29.InsertRight(t54)
	t54.InsertLeft(t13)
	t13.InsertRight(t15)
	tree := t20
	assert.Equal(t, 4, MayorAltura(tree))
}