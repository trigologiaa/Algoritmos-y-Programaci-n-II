package ejercicios

type Objeto struct {
	Nombre string
	Peso   int
	Valor  int
}

func Ejercicio3(objetos []Objeto, capacidad int) map[Objeto]float32 {
	type ValorPorPeso struct {
		Objeto		Objeto
		ValorUnidad	float32
	}
	valorPorPeso := make([]ValorPorPeso, len(objetos))
	for i, objeto := range objetos {
		valorPorPeso[i] = ValorPorPeso {
			Objeto: objeto,
			ValorUnidad: float32(objeto.Valor) / float32(objeto.Peso),
		}
	}
	for i := 0; i < len(valorPorPeso) - 1; i++ {
		for j := i + 1; j < len(valorPorPeso); j++ {
			if valorPorPeso[i].ValorUnidad < valorPorPeso[j].ValorUnidad {
				valorPorPeso[i], valorPorPeso[j] = valorPorPeso[j], valorPorPeso[i]
			}
		}
	}
	seleccionados := make(map[Objeto]float32)
	capacidadRestante := capacidad
	for _, valorPeso := range valorPorPeso {
		if capacidadRestante == 0 {
			break
		}
		if valorPeso.Objeto.Peso <= capacidadRestante {
			seleccionados[valorPeso.Objeto] = 1
			capacidadRestante -= valorPeso.Objeto.Peso
		} else {
			fraccion := float32(capacidadRestante) / float32(valorPeso.Objeto.Peso)
			seleccionados[valorPeso.Objeto] = fraccion
			capacidadRestante = 0
		}
	}
	return seleccionados
}