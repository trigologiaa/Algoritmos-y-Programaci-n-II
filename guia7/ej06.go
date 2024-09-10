package guia7

func MCD(a, b int) int {
	if b == 0 {
		return a
	}
	return MCD(b, a%b)
}
