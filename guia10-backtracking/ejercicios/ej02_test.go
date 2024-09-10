package ejercicios
//16/05/2024
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMochila01(t *testing.T) {
	objetos := []Objeto{
		{Peso: 1, Valor: 1},
		{Peso: 2, Valor: 6},
		{Peso: 3, Valor: 18},
		{Peso: 4, Valor: 22},
		{Peso: 5, Valor: 28},
	}
	capacidad := 11
	valor := Mochila01(objetos, capacidad)
	assert.Equal(t, 56, valor)
}