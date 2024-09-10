package main

import (
	"fmt"
	"subsumamax"
)

/*
Función principal {
	Se declara la variable 'arreglo' de tipo []int con determinados valores
	Se declaran las variables 'sumaMaxima', 'posInicial' y 'posFinal' que son el resultado de los retornos del llamado a la función 'SubsecuenciaSumaMaxima' pasándole 'arreglo' como parámetro
	Se imprime la subsecuencia con suma máxima
	Se imprime la suma máxima
	Se imprime la posición inicial
	Se imprime la posición final
}
*/
func main() {
	arreglo := []int {-1, 6, -2, 5, -1, 4, 3, -4, 3}
	sumaMaxima, posInicial, posFinal := subsumamax.SubsecuenciaSumaMaxima(arreglo)
	fmt.Printf("La subsecuencia con suma máxima es: %v\n", arreglo[posInicial:posFinal + 1])
	fmt.Printf("La suma máxima es: %v\n", sumaMaxima)
	fmt.Printf("La posición inicial es: %v\n", posInicial)
	fmt.Printf("La posición final es: %v\n", posFinal)
}