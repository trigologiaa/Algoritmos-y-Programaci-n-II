package guia7

func BusquedaBinaria(arreglo []int, elemento int) bool {
    return busquedaBinariaRecursiva(arreglo, elemento, 0, len(arreglo)-1)
}

func busquedaBinariaRecursiva(arreglo []int, elemento, inicio, fin int) bool {
    if inicio > fin {
        return false
    }
    medio := inicio + (fin-inicio)/2
    if arreglo[medio] == elemento {
        return true
    } else if elemento < arreglo[medio] {
        return busquedaBinariaRecursiva(arreglo, elemento, inicio, medio-1)
    } else {
        return busquedaBinariaRecursiva(arreglo, elemento, medio+1, fin)
    }
}