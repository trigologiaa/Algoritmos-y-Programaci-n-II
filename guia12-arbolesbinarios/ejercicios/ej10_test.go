package ejercicios

import (
	"testing"
	assert "github.com/stretchr/testify/assert"
	bt "github.com/untref-ayp2/data-structures/binary-tree"
)

func TestLadronDeCasasEjemploUno(t *testing.T) {
	raiz := bt.NewBinaryTree(3)
	a := bt.NewBinaryTree(2)
	b := bt.NewBinaryTree(3)
	c := bt.NewBinaryTree(3)
	d := bt.NewBinaryTree(1)
	raiz.InsertLeft(a)
	raiz.InsertRight(b)
	a.InsertRight(c)
	b.InsertRight(d)
	assert.Equal(t, 7, LadronDeCasas(raiz))
}

func TestLadronDeCasasEjemploDos(t *testing.T) {
	raiz := bt.NewBinaryTree(3)
	a := bt.NewBinaryTree(4)
	b := bt.NewBinaryTree(5)
	c := bt.NewBinaryTree(1)
	d := bt.NewBinaryTree(3)
	e := bt.NewBinaryTree(1)
	raiz.InsertLeft(a)
	raiz.InsertRight(b)
	a.InsertLeft(c)
	a.InsertRight(d)
	b.InsertRight(e)
	assert.Equal(t, 9, LadronDeCasas(raiz))
}