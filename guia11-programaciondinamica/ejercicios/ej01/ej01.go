package ej01

func CostoCaminoMinimo(matriz [][]int) int {
    if len(matriz) == 0 || len(matriz[0]) == 0 {
        return 0
    }
    memoria := make(map[[2]int]int)
    return costoMinimo(matriz, 0, 0, memoria)
}

func costoMinimo(matriz [][]int, x, y int, memoria map[[2]int]int) int {
    if x == len(matriz) - 1 && y == len(matriz[0]) - 1 {
        return matriz[x][y]
    }
    if valor, existe := memoria[[2]int{x, y}]; existe {
        return valor
    }
    costoCorrecto, costoReducir := int(^uint(0) >> 1), int(^uint(0) >> 1)
    if y + 1 < len(matriz[0]) {
        costoCorrecto = costoMinimo(matriz, x, y + 1, memoria)
    }
    if x + 1 < len(matriz) {
        costoReducir = costoMinimo(matriz, x + 1, y, memoria)
    }
    costoMinimoDeRuta := matriz[x][y] + minimo(costoCorrecto, costoReducir)
    memoria[[2]int{x, y}] = costoMinimoDeRuta
    return costoMinimoDeRuta
}

func minimo(a, b int) int {
    if a < b {
        return a
    }
    return b
}

//Función realizada en clase con los profesores
func CostoCaminoMinimo2(matriz [][]int) int {
    aux := make([][]int, len(matriz))
    for i := 0; i < len(matriz); i++ {
        aux[i] = make([]int, len(matriz[i]))
    }
    aux[0][0] = matriz[0][0]
    for i := 1; i < len(matriz[0]); i++ {
        aux[0][i] = aux[0][i - 1] + matriz[0][i]
    }
    for i := 1; i < len(matriz); i++ {
        aux[i][0] = aux[i - 1][0] + matriz[i][0]
    }
    for i := 1; i < len(matriz); i++ {
        for j := 1; j < len(matriz[i]); j++ {
            aux[i][j] = matriz[i][j] + min(aux[i - 1][j], aux[i][j - 1])
        }
    }
    return aux[len(aux) - 1][len(aux[len(aux) - 1]) - 1]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}