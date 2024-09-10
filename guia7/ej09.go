package guia7

func SumaArray(v []int) int {
	if len(v) == 0 {
		return 0
	}
	return v[0] + SumaArray(v[1:])
}