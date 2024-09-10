package guia7

func Pico(arreglo []int) int {
    return encontrarPico(arreglo, 0, len(arreglo)-1)
}

func encontrarPico(arreglo []int, izq, der int) int {
    if izq == der {
        return izq
    }
    mid := izq + (der-izq)/2
    if mid > 0 && mid < len(arreglo)-1 && arreglo[mid] > arreglo[mid-1] && arreglo[mid] > arreglo[mid+1] {
        return mid
    }
    if mid > 0 && arreglo[mid-1] > arreglo[mid] {
        return encontrarPico(arreglo, izq, mid-1)
    } else {
        return encontrarPico(arreglo, mid+1, der)
    }
}