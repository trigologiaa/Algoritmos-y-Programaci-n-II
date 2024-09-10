package ej03

func Escalera(n int, formas []int) int {
    memoria := make(map[int]int)
    return contarFormas(n, formas, memoria)
}

func contarFormas(n int, formas []int, memoria map[int]int) int {
    if n == 0 {
        return 1
    }
    if valor, encontrado := memoria[n]; encontrado {
        return valor
    }
    formasTotales := 0
    for _, forma := range formas {
        if n - forma >= 0 {
            formasTotales += contarFormas(n - forma, formas, memoria)
        }
    }
    memoria[n] = formasTotales
    return formasTotales
}