package ejercicios

import (
	"testing"
	assert "github.com/stretchr/testify/assert"
	bt "github.com/untref-ayp2/data-structures/binary-tree"
)

func TestRecorrerPorNiveles(t *testing.T) {
	a := bt.NewBinaryTree("A")
	b := bt.NewBinaryTree("B")
	c := bt.NewBinaryTree("C")
	d := bt.NewBinaryTree("D")
	e := bt.NewBinaryTree("E")
	f := bt.NewBinaryTree("F")
	g := bt.NewBinaryTree("G")
	h := bt.NewBinaryTree("H")
	i := bt.NewBinaryTree("I")
	j := bt.NewBinaryTree("J")
	k := bt.NewBinaryTree("K")
	l := bt.NewBinaryTree("L")
	a.InsertLeft(b)
	b.InsertLeft(d)
	d.InsertLeft(g)
	d.InsertRight(h)
	h.InsertLeft(k)
	k.InsertRight(l)
	a.InsertRight(c)
	c.InsertLeft(e)
	c.InsertRight(f)
	e.InsertRight(i)
	f.InsertRight(j)
	raiz := a
	expected := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L"}
	assert.Equal(t, expected, RecorrerPorNiveles(raiz))
}