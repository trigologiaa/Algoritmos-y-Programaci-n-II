package estructurasrepetitivas

/*
Función que recibe dos variables 'a' y 'b' de tipo int y devuelve un int que será el producto de los parámetros {
	Se declara la variable 'producto' de tipo int inicializada en 0
	Si 'b' es un número negativo {
		Se cambia el valor de 'a' a negativo
		Se cambia el valor de 'b' a negativo
	}
	Para cuando la variable 'i' inicializada en 0 sea menor a 'b' {
		Se suma 'a' a 'producto'
	}
	Se retorna 'producto'
}
*/
func Producto(a, b int) int {
	producto := 0
	if b < 0 {
		a = -a
		b = -b
	}
	for i := 0; i < b; i++ {
		producto += a
	}
	return producto
}