package ej04

func Mochila(valores, pesos []int, capacidad int) int {
	elementos := len(valores)
	matriz := make([][]int, elementos + 1)
	for i := range matriz {
		matriz[i] = make([]int, capacidad + 1)
	}
	for i := 1; i <= elementos; i++ {
		for j := 1; j <= capacidad; j++ {
			if pesos[i - 1] <= j {
				matriz[i][j] = maximo(matriz[i - 1][j], matriz[i - 1][j - pesos[i - 1]] + valores[i - 1])
			} else {
				matriz[i][j] = matriz[i - 1][j]
			}
		}
	}
	return matriz[elementos][capacidad]
}

func maximo(a, b int) int {
	if a > b {
		return a
	}
	return b
}