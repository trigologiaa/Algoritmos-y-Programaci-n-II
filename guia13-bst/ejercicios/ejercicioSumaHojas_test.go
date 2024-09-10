package ejercicios

import (
	"testing"
	"github.com/stretchr/testify/assert"
	tree "github.com/untref-ayp2/data-structures/tree"
)

func TestSumaHojasMenoresQue10(t *testing.T) {
	// Caso de árbol vacío
	btEmpty := tree.NewBinarySearchTree[int]()
	assert.Equal(t, 0, SumaHojasMenoresQue10(btEmpty))

	// Caso con hojas menores que 10
	bt1 := tree.NewBinarySearchTree[int]()
	bt1.Insert(5)
	bt1.Insert(8)
	bt1.Insert(12)
	bt1.Insert(3)
	bt1.Insert(9)
	assert.Equal(t, 8, SumaHojasMenoresQue10(bt1))

	// Caso sin hojas menores que 10
	bt2 := tree.NewBinarySearchTree[int]()
	bt2.Insert(15)
	bt2.Insert(18)
	bt2.Insert(12)
	assert.Equal(t, 0, SumaHojasMenoresQue10(bt2))

	// Caso con árbol nulo
	assert.Equal(t, 0, SumaHojasMenoresQue10(nil))
}