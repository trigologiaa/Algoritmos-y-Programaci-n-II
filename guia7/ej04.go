package guia7

func Invertir(cadena string) string {
	if len(cadena) <= 1 {
		return cadena
	}
	return Invertir(cadena[1:]) + string(cadena[0])
}
