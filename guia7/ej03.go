package guia7

func CantidadDeUnos(n int) int {
	if n == 0 {
		return 0
	}
	if n%2 == 0 {
		return CantidadDeUnos(n / 2)
	} else {
		return 1 + CantidadDeUnos(n/2)
	}
}
