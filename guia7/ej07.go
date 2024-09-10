package guia7

func Multiplicar(a, b int) int {
	if b == 0 {
		return 0
	} else if b > 0 {
		return a + Multiplicar(a, b-1)
	} else {
		return -Multiplicar(a, -b)
	}
}