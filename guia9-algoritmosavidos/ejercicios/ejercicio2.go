package ejercicios

type Tarea struct {
	Nombre string
	Tiempo int
}

func Ejercicio2(tareas []Tarea) {
	for i := 0; i < len(tareas); i++ {
		minimo := i
		for j := i + 1; j < len(tareas); j++ {
			if tareas[j].Tiempo < tareas[minimo].Tiempo {
				minimo = j
			}
		}
		tareas[i], tareas[minimo] = tareas[minimo], tareas[i]
	}
}