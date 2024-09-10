package guia7

func EsPotencia(n, b int) bool {
	if n == 1 {
		return true
	}
	if b <= 1 {
		return n == 1
	}
	if n == 0 || n%b != 0 {
		return false
	}
	return EsPotencia(n/b, b)
}
