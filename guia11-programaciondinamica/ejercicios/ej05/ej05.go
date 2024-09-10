package ej05

func SumaSubconjunto(arr []int, k int) int {
    memoria := make(map[int]int)
    return contarSubconjuntos(arr, len(arr)-1, k, memoria)
}

func contarSubconjuntos(arr []int, indice, suma int, memoria map[int]int) int {
    if suma == 0 {
        return 1
    }
    if suma < 0 || indice < 0 {
        return 0
    }
    llave := (indice << 16) | suma
    if valor, encontrado := memoria[llave]; encontrado {
        return valor
    }
    incluido := contarSubconjuntos(arr, indice - 1, suma - arr[indice], memoria)
    excluido := contarSubconjuntos(arr, indice - 1, suma, memoria)
    memoria[llave] = incluido + excluido
    return memoria[llave]
}