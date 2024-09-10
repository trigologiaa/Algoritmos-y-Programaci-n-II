package estructurasrepetitivas

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFactorial5(test *testing.T) {
	resultado := Factorial(5)
	assert.Equal(test, 120, resultado)
}

func TestFactorial0(test *testing.T) {
	resultado := Factorial(0)
	assert.Equal(test, 1, resultado)
}

func TestFactorial1(test *testing.T) {
	resultado := Factorial(1)
	assert.Equal(test, 1, resultado)
}