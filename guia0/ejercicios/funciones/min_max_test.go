package funciones

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestEncontrarMinimoMaximo1(test *testing.T) {
	lista := []int {4, 7, 2, 9, 1}
	minimo, maximo := EncontrarMinimoMaximo(lista)
	assert.Equal(test, 1, minimo)
	assert.Equal(test, 9, maximo)
}

func TestEncontrarMinimoMaximo2(test *testing.T) {
	lista := []int {23, 5, 17, 9, 34, 12, 8, 41, 20, 3}
	minimo, maximo := EncontrarMinimoMaximo(lista)
	assert.Equal(test, 3, minimo)
	assert.Equal(test, 41, maximo)
}

func TestEncontrarMinimoMaximo3(test *testing.T) {
	lista := []int {10, 27, 6, 14, 38, 19, 4, 31, 22, 7}
	minimo, maximo := EncontrarMinimoMaximo(lista)
	assert.Equal(test, 4, minimo)
	assert.Equal(test, 38, maximo)
}

func TestEncontrarMinimoMaximo4(test *testing.T) {
	lista := []int {15, 33, 2, 11, 29, 18, 5, 25, 13, 30}
	minimo, maximo := EncontrarMinimoMaximo(lista)
	assert.Equal(test, 2, minimo)
	assert.Equal(test, 33, maximo)
}

func TestEncontrarMinimoMaximo5(test *testing.T) {
	lista := []int {}
	minimo, maximo := EncontrarMinimoMaximo(lista)
	assert.Equal(test, 0, minimo)
	assert.Equal(test, 0, maximo)
}

func TestEncontrarMinimoMaximo6(test *testing.T) {
	lista := []int {8, 24, 3, 17, 36, 12, 6, 29, 19, 10}
	minimo, maximo := EncontrarMinimoMaximo(lista)
	assert.Equal(test, 3, minimo)
	assert.Equal(test, 36, maximo)
}

func TestEncontrarMinimoMaximo7(test *testing.T) {
	lista := []int {14, 27, 5, 19, 35, 8, 23, 11, 31, 16, 29, 7, 22, 10, 25, 3, 18, 6, 30, 12}
	minimo, maximo := EncontrarMinimoMaximo(lista)
	assert.Equal(test, 3, minimo)
	assert.Equal(test, 35, maximo)
}