package punteros

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSwap(test *testing.T) {
	a := 1
	b := 2
	Swap(&a, &b)
	assert.Equal(test, 2, a)
	assert.Equal(test, 1, b)
}