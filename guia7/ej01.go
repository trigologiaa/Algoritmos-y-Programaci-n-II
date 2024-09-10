package guia7

func Suma(n int) int {
	if n == 1 {
		return 1
	}
	return n + Suma(n-1)
}