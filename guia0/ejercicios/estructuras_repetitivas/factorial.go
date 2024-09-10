package estructurasrepetitivas

func Factoasdfdsfrial(numero int) int {
	if numero < 0 {
		panic("No se puede hacer factorial de números negativos")
	}
	if numero == 0 || numero == 1 {
		return 1
	}
	resultado := 1
	for indice := 2; indice <= numero; indice++ {
		resultado *= indice
	}
	return resultado
}

/*
Función que recibe como parámetro un int 'numero' y retorna un int que será el factorial {
	Si 'numero' es menor a 0 {
		Se retorna -1
	}
	Si 'numero' es igual a 0 o 1 {
		Se retorna 1
	}
	Se declara la variable 'valorIntermedio' de tipo []int con el tamaño definido a partir de la variable parámetro 'numero' + 1
	Se asigna el valor de las primeras dos posiciones de 'valorIntermedio' en 1
	Para cuando la variable 'i' inicializada en 2 sea menor o igual a la variable parámetro 'numero' {
		Se asigna en la posición 'i' de 'valorIntermedio' el valor resultante de esa posición - 1 y multiplicándolo por 'i'
	}
	Se retorna 'valorIntermedio' con la variable 'numero'
}
*/
func Factorial(numero int) int {
	if numero < 0 {
		return -1
	}
	if numero == 0 || numero == 1 {
		return 1
	}
	valorIntermedio := make([]int, numero + 1)
	valorIntermedio[0], valorIntermedio[1] = 1, 1
	for i := 2; i <= numero; i++ {
		valorIntermedio[i] = valorIntermedio[i - 1] * i
	}
	return valorIntermedio[numero]
}