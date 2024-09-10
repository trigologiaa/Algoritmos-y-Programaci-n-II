package guia7

func DecimalABinario(n int) string {
	if n == 0 {
		return "0"
	} else if n == 1 {
		return "1"
	} else {
		return DecimalABinario(n/2) + DecimalABinario(n%2)
	}
}
