package funciones

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestImprimirPolinomioDe3(test *testing.T) {
	coeficientes := []float64 {
		3.5,
		2.5,
		1.5,
	}
	resultadoEsperado := "3.5 + 2.5 x + 1.5 x**2"
	resultado := ImprimirPolinomio(coeficientes)
	assert.Equal(test, resultadoEsperado, resultado)
}

func TestImprimirPolinomioDe2(test *testing.T) {
	coeficientes := []float64 {
		3.5,
		2.5,
	}
	resultadoEsperado := "3.5 + 2.5 x"
	resultado := ImprimirPolinomio(coeficientes)
	assert.Equal(test, resultadoEsperado, resultado)
}

func TestImprimirPolinomioDe1(test *testing.T) {
	coeficientes := []float64 {
		3.5,
	}
	resultadoEsperado := "3.5"
	resultado := ImprimirPolinomio(coeficientes)
	assert.Equal(test, resultadoEsperado, resultado)
}

func TestImprimirPolinomioVacio(test *testing.T) {
	coeficientes := []float64 {}
	resultadoEsperado := "El array está vacío"
	resultado := ImprimirPolinomio(coeficientes)
	assert.Equal(test, resultadoEsperado, resultado)
}