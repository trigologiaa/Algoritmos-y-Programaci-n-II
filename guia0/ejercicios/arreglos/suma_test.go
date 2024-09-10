package arreglos

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSuma23(test *testing.T) {
	lista := []int {
		4,
		7,
		2,
		9,
		1,
	}
	resultado := Suma(lista)
	assert.Equal(test, 23, resultado)
}

func TestSuma45(test *testing.T) {
	lista := []int {
		0,
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
	}
	resultado := Suma(lista)
	assert.Equal(test, 45, resultado)
}

func TestSumaMenos10(test *testing.T) {
	lista := []int {
		5,
		-15,
	}
	resultado := Suma(lista)
	assert.Equal(test, -10, resultado)
}

func TestSuma0(test *testing.T) {
	lista := []int {
		0,
	}
	resultado := Suma(lista)
	assert.Equal(test, 0, resultado)
}