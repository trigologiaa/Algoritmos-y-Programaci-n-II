package ejercicios

import (
	"testing"
	"github.com/stretchr/testify/assert"
	tree "github.com/untref-ayp2/data-structures/tree"
)

//	   4
//	  / \
//	 2   5
//	/ \   \
// 1   3   6
//        / \
//       8   7

func TestNoIsBst(t *testing.T) {
	t1 := tree.NewBinaryTree(1)
	t3 := tree.NewBinaryTree(3)
	t2 := tree.NewBinaryTree(2)
	t5 := tree.NewBinaryTree(5)
	t4 := tree.NewBinaryTree(4)
	t6 := tree.NewBinaryTree(6)
	t7 := tree.NewBinaryTree(7)
	t8 := tree.NewBinaryTree(8)
	t2.InsertLeft(t1)
	t2.InsertRight(t3)
	t4.InsertLeft(t2)
	t4.InsertRight(t5)
	t6.InsertRight(t7)
	t6.InsertLeft(t8)
	t5.InsertRight(t6)
	tree := t4
	assert.False(t, IsBst(tree))
}

//	   4
//	  / \
//	 2   15
//	/ \    \
// 1   3    17
//           \
//           18

func TestSiIsBst(t *testing.T) {
	t1 := tree.NewBinaryTree(1)
	t3 := tree.NewBinaryTree(3)
	t2 := tree.NewBinaryTree(2)
	t15 := tree.NewBinaryTree(15)
	t4 := tree.NewBinaryTree(4)
	t17 := tree.NewBinaryTree(17)
	t18 := tree.NewBinaryTree(18)
	t2.InsertLeft(t1)
	t2.InsertRight(t3)
	t4.InsertLeft(t2)
	t4.InsertRight(t15)
	t15.InsertRight(t17)
	t17.InsertRight(t18)
	tree := t4
	assert.True(t, IsBst(tree))
}

func TestNilIsBst(t *testing.T) {
	t1 := tree.NewBinaryTree(1)
	t1.Empty()
	tree := t1
	assert.True(t, IsBst(tree))
}

func TestNodoIsBst(t *testing.T) {
	t1 := tree.NewBinaryTree(1)
	tree := t1
	assert.True(t, IsBst(tree))
}

//	   4
//	  / \
//	 2   15
//	/ \    \
// 14  3    17
//           \
//           18

func TestMalMinIsBst(t *testing.T) {
	t1 := tree.NewBinaryTree(14)
	t3 := tree.NewBinaryTree(3)
	t2 := tree.NewBinaryTree(2)
	t15 := tree.NewBinaryTree(15)
	t4 := tree.NewBinaryTree(4)
	t17 := tree.NewBinaryTree(17)
	t18 := tree.NewBinaryTree(18)
	t2.InsertLeft(t1)
	t2.InsertRight(t3)
	t4.InsertLeft(t2)
	t4.InsertRight(t15)
	t15.InsertRight(t17)
	t17.InsertRight(t18)
	tree := t4
	assert.False(t, IsBst(tree))
}

//	   4
//	  / \
//	 2   15
//	/ \    \
// 1   3    17
//           \
//           -18

func TestMalMaxSiIsBst(t *testing.T) {
	t1 := tree.NewBinaryTree(1)
	t3 := tree.NewBinaryTree(3)
	t2 := tree.NewBinaryTree(2)
	t15 := tree.NewBinaryTree(15)
	t4 := tree.NewBinaryTree(4)
	t17 := tree.NewBinaryTree(17)
	t18 := tree.NewBinaryTree(-18)
	t2.InsertLeft(t1)
	t2.InsertRight(t3)
	t4.InsertLeft(t2)
	t4.InsertRight(t15)
	t15.InsertRight(t17)
	t17.InsertRight(t18)
	tree := t4
	assert.False(t, IsBst(tree))
}