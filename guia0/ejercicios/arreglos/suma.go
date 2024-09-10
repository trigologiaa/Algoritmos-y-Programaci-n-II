package arreglos

/*
Función que recibe una variable 'arreglo' de tipo []int y devuelve un int que sería la suma de todos los valores {
	Se declara la variable 'suma' inicializada en 0
	Para cada valor de 'arreglo' inicializado en la variable 'v' {
		Se suma 'v' a 'suma'
	}
	Se retorna 'suma'
}
*/
func Suma(arreglo []int) int {
	suma := 0
	for _, v := range arreglo {
		suma += v
	}
	return suma
}