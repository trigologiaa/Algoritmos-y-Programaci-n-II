package guia7

func DivisionEntera(dividendo, divisor int) (cociente, resto int) {
	if divisor == 0 {
		return 0, 0
	}
	if dividendo < divisor {
		return 0, dividendo
	}
	cociente, resto = DivisionEntera(dividendo-divisor, divisor)
	return cociente + 1, resto
}
