package ejercicios
//16/05/2024
func Dados(n, k, x int) int {
    return backtrack1(n, k, x, 0)
}

func backtrack1(n, k, x, sumaActual int) int {
    if n == 0 {
        if sumaActual == x {
            return 1
        }
        return 0
    }
    if n == 0 && sumaActual != x {
        return 0
    }
    formas := 0
    for i := 1; i <= k; i++ {
        formas += backtrack1(n - 1, k, x, sumaActual + i)
    }
    return formas
}