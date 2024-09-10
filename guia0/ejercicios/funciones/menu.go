package funciones

/*
Función que recibe como parámetro un int que funciona como 'opcion' y devuelve un int {
	Si la 'opcion' está entre el 1 y el 4 {
		Se retorna el int de 'opcion'
	}
	Se retorna 0
}
*/
func ElegirOpcion(opcion int) int {
	if opcion >= 1 && opcion <= 4 {
		return opcion
	}
	return 0
}