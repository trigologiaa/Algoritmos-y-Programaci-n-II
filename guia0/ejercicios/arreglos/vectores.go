package arreglos

/*
Función que recibe las variables 'vectr1' y 'vector2' de tipo []int como parámetros y devuelve dos int que son la suma y el producto escalar de los vectores {
	Se declara la variable 'suma' de tipo int inicializada en 0
	Se declara la variable 'productoEscalar' de tipo int inicializada en 0
	Para cuando la variable 'i' de tipo int inicializada en 0 sea menor al largo de 'vector1' {
		Se suman los valores 'i' de 'vector1' y 'vector2' a 'suma'
		Se suma el producto 'i' de 'vector1' y 'vector2' a 'productoEscalar'
	}
	Se retornan 'suma' y 'productoEscalar'
}
*/
func SumaProductoEscalar(vector1, vector2 []int) (int, int) {
	suma := 0
	productoEscalar := 0
	for i := 0; i < len(vector1); i++ {
		suma += vector1[i] + vector2[i]
		productoEscalar += vector1[i] * vector2[i]
	}
	return suma, productoEscalar
}