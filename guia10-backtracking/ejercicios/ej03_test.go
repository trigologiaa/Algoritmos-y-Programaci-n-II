package ejercicios
//16/05/2024
import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMochila3(t *testing.T) {
	objetos := []Objeto{
		{Peso: 1, Valor: 1},
		{Peso: 2, Valor: 6},
		{Peso: 3, Valor: 18},
		{Peso: 4, Valor: 22},
		{Peso: 5, Valor: 28},
	}
	capacidad := 11
	valor, solucion := Mochila2(objetos, capacidad)
	fmt.Println(solucion)
	assert.Equal(t, 56, valor)
}

func TestMochila4(t *testing.T) {
	objetos := []Objeto{
		{Peso: 1, Valor: 10},
		{Peso: 2, Valor: 6},
		{Peso: 3, Valor: 18},
		{Peso: 4, Valor: 22},
		{Peso: 5, Valor: 28},
	}
	capacidad := 11
	valor, solucion := Mochila2(objetos, capacidad)
	fmt.Println(solucion)
	assert.Equal(t, 62, valor)
}

func TestMochila3Punto2(t *testing.T) {
	objetos := []Objeto{
		{Peso: 1, Valor: 1},
		{Peso: 2, Valor: 6},
		{Peso: 3, Valor: 18},
		{Peso: 4, Valor: 22},
		{Peso: 5, Valor: 28},
	}
	capacidad := 11
	valor, solucion := Mochila2Punto2(objetos, capacidad)
	fmt.Println(solucion)
	assert.Equal(t, 56, valor)
}

func TestMochila4Punto2(t *testing.T) {
	objetos := []Objeto{
		{Peso: 1, Valor: 10},
		{Peso: 2, Valor: 6},
		{Peso: 3, Valor: 18},
		{Peso: 4, Valor: 22},
		{Peso: 5, Valor: 28},
	}
	capacidad := 11
	valor, solucion := Mochila2Punto2(objetos, capacidad)
	fmt.Println(solucion)
	assert.Equal(t, 62, valor)
}