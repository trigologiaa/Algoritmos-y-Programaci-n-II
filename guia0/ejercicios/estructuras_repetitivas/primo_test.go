package estructurasrepetitivas

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestEsPrimo5(test *testing.T) {
	resultado := EsPrimo(5)
	assert.Equal(test, true, resultado)
}

func TestEsPrimo7(test *testing.T) {
	resultado := EsPrimo(7)
	assert.Equal(test, true, resultado)
}

func TestEsPrimo11(test *testing.T) {
	resultado := EsPrimo(11)
	assert.Equal(test, true, resultado)
}

func TestNoEsPrimo4(test *testing.T) {
	resultado := EsPrimo(4)
	assert.Equal(test, false, resultado)
}

func TestNoEsPrimo8(test *testing.T) {
	resultado := EsPrimo(8)
	assert.Equal(test, false, resultado)
}

func TestNoEsPrimo12(test *testing.T) {
	resultado := EsPrimo(12)
	assert.Equal(test, false, resultado)
}