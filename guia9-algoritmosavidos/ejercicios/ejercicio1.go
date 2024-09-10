package ejercicios

import "sort"

type Actividad struct {
	Nombre string
	Inicio int
	Fin    int
}

// Selector de actividades recursivo
// Reescribir la función que resuelve el problema de seleccionar actividades en forma recursiva
func SelectorRecursivo(actividades []Actividad) []Actividad {
	sort.Slice(actividades, func(i, j int) bool {
		return actividades[i].Fin < actividades[j].Fin
	})
	var seleccionRecursiva func([]Actividad, int) []Actividad
	seleccionRecursiva = func(actividades []Actividad, ultimoFin int) []Actividad {
		if len(actividades) == 0 {
			return []Actividad{}
		}
		if actividades[0].Inicio >= ultimoFin {
			seleccion := []Actividad{actividades[0]}
			return append(seleccion, seleccionRecursiva(actividades[1:], actividades[0].Fin)...)
		}
		return seleccionRecursiva(actividades[1:], ultimoFin)
	}
	return seleccionRecursiva(actividades, 0)
}