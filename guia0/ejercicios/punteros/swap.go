package punteros

/*
Función que recibe dos variables 'x' e 'y' de tipo puntero int {
	Se declara la variable 'px' inicializada con el valor del puntero de 'x'
	Se declara la variable 'py' inicializada con el valor del puntero de 'y'
	Se le cambia el valor del puntero de 'x' por el valor de 'py'
	Se le cambia el valor del puntero de 'y' por el valor de 'px'
}
*/
func Swap(x, y *int) {
	px := *x
	py := *y
	*x  = py
	*y = px
}