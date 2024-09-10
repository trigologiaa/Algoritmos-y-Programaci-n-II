package utiles

import (
	"math/rand"
	"time"
)

// GenerarArreglo genera un arreglo de tamaño tam con números aleatorios entre 0 y numeros.
/*
Función que recibe las variables 'numeros' y 'tam' de tipo int y devuelve un []int {
	Se declara la variable 'r' de tipo puntero a rand.Rand que se inicializa con el resultado de la función 'New' de 'math/rand'
	Se declara la variable 'arreglo' de tipo []int inicializada con la función 'make' y con el tamaño de la variable 'tam'
	Para cuando la variable 'i' de tipo int inicializada en 0 sea menor a 'tam' {
		Se cambia el valor de la posición 'i' de 'arreglo' por el resultado de la función 'Intn' que recibe 'numeros' como parámetro
	}
	Se retorna 'arreglo'
}
*/
func GenerarArreglo(numeros, tam int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	arreglo := make([]int, tam)
	for i := 0; i < tam; i++ {
		arreglo[i] = r.Intn(numeros)
	}
	return arreglo
}
