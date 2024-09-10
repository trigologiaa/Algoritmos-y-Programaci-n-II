package guia7

func Repeticiones(cadena string, caracter rune) int {
	if len(cadena) == 0 {
		return 0
	}
	var count int
	if rune(cadena[0]) == caracter {
		count = 1
	} else {
		count = 0
	}
	return count + Repeticiones(cadena[1:], caracter)
}
