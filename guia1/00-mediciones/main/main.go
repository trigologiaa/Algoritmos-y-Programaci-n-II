package main

import (
	"busquedas"
	"fmt"
	"sort"
	"time"
	"utiles"
)


func main() {
	arreglo := utiles.GenerarArreglo(1000, 10000000)
	buscado := 849
	//fmt.Println(arreglo) //IMPRIME EL ARREGLO COMPLETO
	inicio := time.Now()
	// Busqueda Lineal
	fmt.Println("Indice donde se encontró: ", busquedas.BusLineal(arreglo, buscado)) //IMPRIME EL ÍNDICE EN EL QUE SE ENCONTRÓ EL VALOR
	fmt.Println("Busqueda Lineal: ", time.Since(inicio))                             //IMPRIME EL TIEMPO QUE DEMORÓ EN ENCONTRAR O NO EL VALOR
	inicio = time.Now()
	// Ordenar el arreglo para la busqueda binaria
	sort.Ints(arreglo)
	fmt.Println("Ordenamiento: ", time.Since(inicio)) //IMPRIME EL TIEMPO QUE DEMORA EN ORDENAR
	//fmt.Println(arreglo)                              //IMPRIME EL ARREGLO COMPLETO
	inicio = time.Now()
	// Busqueda Binaria
	fmt.Println("Indice donde se encontró: ", busquedas.BusquedaBinaria(arreglo, buscado)) //IMPRIME EL ÍNDICE EN EL QUE SE ENCONTRÓ EL VALOR
	fmt.Println("Busqueda Binaria: ", time.Since(inicio))                                  //IMPRIME EL TIEMPO QUE DEMORÓ EN ENCONTRAR O NO EL VALOR
	//Busqueda Burbujeo
	inicio = time.Now()
	fmt.Println("Indice donde se encontró: ", busquedas.BusquedaBurbujeo(arreglo, buscado))
	fmt.Println("Busqueda Burbujeo: ", time.Since(inicio))
}
