package funciones

import(
	"fmt"
)

/*
Función que recibe como parámetro un arreglo de float64, que representa los coeficientes a imprimir {
	Si el arreglo está vacío {
		Se retorna "El array está vacío"
	}
	Se declara la variable 'resultado' como string vacío
	Por cada índice y valor del arreglo 'coeficientes' {
		Si es la primer posición del arreglo {
			Se le suma a la variable 'resultado' el valor de la primera posición arreglo 'coeficiente'
		} Si no, entonces si es la segunda posición del arreglo {
			Se le suma a la variable 'resultado' el valor de la segunda posición del arreglo 'coeficientes'
		} Si no, entonces {
			Se le suma a la variable 'resultado' el valor de cada posición siguiente del arreglo 'coeficientes'
		}
	}
	Se retorna la variable resultado
}
*/
func ImprimirPolinomio(coeficientes []float64) string {
    if len(coeficientes) == 0 {
        return "El array está vacío"
    }
    resultado := ""
    for indice, valor := range coeficientes {
        if indice == 0 {
            resultado += fmt.Sprintf("%.1f", valor)
        } else if indice == 1 {
            resultado += fmt.Sprintf(" + %.1f x", valor)
        } else {
            resultado += fmt.Sprintf(" + %.1f x**%d", valor, indice)
        }
    }
    return resultado
}