package estructurasrepetitivas

/*
Función que recibe como parámetro una variable 'numero' de tipo int y retorna bool dependiendo de si el número es primo o no {
	Si 'numero' es menor o igual a 1 o el resto de su división por 2 o 3 da 0 {
		Se retorna false
	}
	Si 'numero' es menor o igual a 3 {
		Se retorna true
	}
	Para cuando la variable 'i' inicializada en 5, multiplicada por sí misma sea menor o igual a 'numero' y aumentando en 6 {
		Si el resto de la división de 'numero' por 'i' o de 'i' + 2 es 0 {
			Se retorna false
		}
	}
	Se retorna true
}
*/
func EsPrimo(numero int) bool {
	if numero <= 1 || numero % 2 == 0 || numero % 3 == 0 {
		return false
	}
	if numero <= 3 {
		return true
	}
	for i := 5; i * i <= numero; i += 6 {
		if numero % i == 0 || numero % (i + 2) == 0 {
			return false
		}
	}
	return true
}