package estructurasrepetitivas

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestProducto3x4(test *testing.T) {
	resultado := Producto(3, 4)
	assert.Equal(test, 12, resultado)
}

func TestProducto5x5(test *testing.T) {
	resultado := Producto(5, 5)
	assert.Equal(test, 25, resultado)
}

func TestProducto6x8(test *testing.T) {
	resultado := Producto(6, 8)
	assert.Equal(test, 48, resultado)
}

func TestProducto264x536(test *testing.T) {
	resultado := Producto(264, 536)
	assert.Equal(test, 141504, resultado)
}

func TestProducto5Negativox6(test *testing.T) {
	resultado := Producto(-5, 6)
	assert.Equal(test, -30, resultado)
}

func TestProducto6x8Negativo(test *testing.T) {
	resultado := Producto(6, -8)
	assert.Equal(test, -48, resultado)
}