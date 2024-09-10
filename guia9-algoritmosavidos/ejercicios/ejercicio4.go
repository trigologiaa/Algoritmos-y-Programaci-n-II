package ejercicios

import "errors"

type Farolas struct {
	Posicion  int
	Radio     int
	Encendida bool
}

func EncenderFarolas(farolas []Farolas, puntos []int) ([]Farolas, error) {
	var resultado []Farolas
	n := len(farolas)
	m := len(puntos)
	for i := 0; i < m; {
		mejor := -1
		for j := 0; j < n; j++ {
			if farolas[j].Posicion - farolas[j].Radio <= puntos[i] && farolas[j].Posicion + farolas[j].Radio >= puntos[i] {
				if mejor == -1 || farolas[j].Posicion + farolas[j].Radio > farolas[mejor].Posicion + farolas[mejor].Radio {
					mejor = j
				}
			}
		}
		if mejor == -1 {
			return nil, errors.New("no se puede iluminar todos los puntos :p")
		}
		farolas[mejor].Encendida = true
		resultado = append(resultado, farolas[mejor])
		for i < m && puntos[i] <= farolas[mejor].Posicion + farolas[mejor].Radio {
			i++
		}
	}
	return resultado, nil
}