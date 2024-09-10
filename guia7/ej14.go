package guia7

func BusquedaTernariaRecursiva(arr []int, x int) int {
	return busquedaTernaria(arr, x, 0, len(arr)-1)
}

func busquedaTernaria(arr []int, x, left, right int) int {
	if right >= left {
		tercio1 := left + (right-left)/3
		tercio2 := right - (right-left)/3
		if arr[tercio1] == x {
			return tercio1
		}
		if arr[tercio2] == x {
			return tercio2
		}
		if x < arr[tercio1] {
			return busquedaTernaria(arr, x, left, tercio1-1)
		} else if x > arr[tercio2] {
			return busquedaTernaria(arr, x, tercio2+1, right)
		} else {
			return busquedaTernaria(arr, x, tercio1+1, tercio2-1)
		}
	}
	return -1
}
