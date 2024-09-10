package funciones

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestElegirOpcionValida(test *testing.T) {
	for indice := 1; indice <= 4; indice++ {
		resultado := ElegirOpcion(indice)
		assert.Equal(test, indice, resultado)
	}
}
func TestElegirOpcionInvalida(test *testing.T) {
	resultado := ElegirOpcion(5)
	assert.Equal(test, 0, resultado)
}