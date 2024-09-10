package subsumamax

// Dado un arreglo devuelve la posicion inicial, la posición final y el valor de la subsecuencia cuya suma es máxima dentro del arreglo original.
/*
Función que recibe como parámetro una variable 'arreglo' de tipo []int y devuelve tres int, que son la suma máxima, la posición inicial y la posición final {
    Si el largo del arreglo es igual a 0 {
        Se retornan 0, -1 y -1
    }
    Se declara la variable 'sumaMaxima' de tipo int inicializada con la primera posición de 'arreglo'
    Se declara la variable 'sumaActual' de tipo int inicializada con la primera posición de 'arreglo'
    Se declaran las variables 'posInicial' y 'posFinal' de tipo int inicializadas en 0
    Se declara la variable 'inicioTemporal' de tipo int inicializada en 0
    Para cuando la variable 'i' de tipo init inicializada en 1 sea menor al largo de 'arreglo' {
        Si 'sumaActual' es menor a 0 {
            Se cambia el valor de 'sumaActual' por el de la posición 'i' de 'arreglo'
            Se cambia el valor de 'inicioTemporal' por el valor 'i'
        } Sino {
            Se suma el valor de la posición 'i' de 'arreglo' a 'sumaActual' 
        }
        Si 'sumaActual' es mayor a 'sumaMaxima' {
            Se cambia el valor de 'sumaMaxima' por el valor de 'sumaActual'
            Se cambia el valor de 'posInicial' por el valor de 'inicioTemporal'
            Se cambia el valor de 'posFinal' por el valor de 'i'
        }
    }
    Se retornan 'sumaMaxima', 'posInicial' y 'posFinal'
}
*/
func SubsecuenciaSumaMaximaOdeN(arreglo []int) (int, int, int) {
    if len(arreglo) == 0 {
        return 0, -1, -1
    }
    sumaMaxima := arreglo[0]
    sumaActual := arreglo[0]
    posInicial, posFinal := 0, 0
    inicioTemporal := 0
    for i := 1; i < len(arreglo); i++ {
        if sumaActual < 0 {
            sumaActual = arreglo[i]
            inicioTemporal = i
        } else {
            sumaActual += arreglo[i]
        }
        if sumaActual > sumaMaxima {
            sumaMaxima = sumaActual
            posInicial = inicioTemporal
            posFinal = i
        }
    }
    return sumaMaxima, posInicial, posFinal
}