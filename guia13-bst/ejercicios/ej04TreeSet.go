package ejercicios

import (
	tree "github.com/untref-ayp2/data-structures/tree"
	"github.com/untref-ayp2/data-structures/types"
	"strconv"
)

type TreeSet[T types.Ordered] struct {
	set *tree.BinarySearchTree[T]
}

func NewTreeSet[T types.Ordered](elementos ...T) *TreeSet[T] {
	ts := &TreeSet[T]{
		set: tree.NewBinarySearchTree[T](),
	}
	ts.Add(elementos...)
	return ts
}

func (ts *TreeSet[T]) Add(elementos ...T) {
	for _, elemento := range elementos {
		if !ts.Contains(elemento) {
			ts.set.Insert(elemento)
		}
	}
}

func (ts *TreeSet[T]) Size() int {
	return ts.set.Size()
}

func (ts *TreeSet[T]) Contains(elemento T) bool {
	return ts.set.Search(elemento) != nil
}

func (ts *TreeSet[T]) Remove(elemento T) {
	ts.set.Remove(elemento)
}

func (ts *TreeSet[T]) Values() []T {
	if ts.Size() == 0 {
		return []T{}
	}
	return ts.set.GetRoot().GetInOrder()
}

func (ts *TreeSet[T]) String() string {
	valores := ts.Values()
	resultado := "Set: {"
	for i, v := range valores {
		if i > 0 {
			resultado += " "
		}
		resultado += toString(v)
	}
	resultado += "}"
	return resultado
}

func toString[T types.Ordered](v T) string {
	switch any(v).(type) {
	case int, int8, int16, int32, int64:
		return strconv.Itoa(any(v).(int))
	case uint, uint8, uint16, uint32, uint64:
		return strconv.Itoa(int(any(v).(uint)))
	case float32, float64:
		return strconv.FormatFloat(any(v).(float64), 'f', -1, 64)
	case string:
		return any(v).(string)
	default:
		return ""
	}
}