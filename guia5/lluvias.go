package guia5

import (
	bm	"github.com/untref-ayp2/data-structures/bitmap"
	s	"github.com/untref-ayp2/data-structures/set"
)

type Mes uint8

const (
	enero Mes = iota
	febrero
	marzo
	abril
	mayo
	junio
	julio
	agosto
	septiembre
	octubre
	noviembre
	diciembre
)

/*
Función que recibe como parámetro una variable 'dia' de tipo uint8 y una variable 'mes' de tipo Mes y retorna un bool {
	Para cada caso en 'mes' {
	En caso de que sea 'abril', 'junio', 'septiembre' o 'noviembre':
		Se retorna 'dia' como mayor o igual a 1 y menor o igual a 30
	En caso de que sea 'febrero':
		Se retorna 'dia' como mayor o igual a 1 y menor o igual a 29
	En otro caso:
		Se retorna 'dia' como mayor o igual a 1 y menor o igual a 31
	}
}
*/
func validaFecha(dia uint8, mes Mes) bool {
	switch mes {
	case abril, junio, septiembre, noviembre:
		return dia >= 1 && dia <= 30
	case febrero:
		return dia >= 1 && dia <= 29
	default:
		return dia >= 1 && dia <= 31
	}
}

/*
Tipo de estructura {
	Se declara el campo 'registros' de tipo [12]*BitMap
}
*/
type Lluvias struct {
	registros	[12]*bm.BitMap
}

/*
Función que retorna un *Lluvias {
	Se declara la variable 'lluvia' de tipo *Lluvias inicializada con la dirección de 'Lluvias' vacío {}
	Para cada valor 'i' inicializado con el indice del rango del campo 'registros' de 'lluvia' {
		Se cambia el valor de la posición 'i' del campo 'registros' de 'lluvia' por el retorno del llamado a la función 'NewBitMap'
	}
	Se retorna 'lluvia'
}
*/
func NewLluvias() *Lluvias {
	lluvia := &Lluvias {}
	for i := range lluvia.registros {
		lluvia.registros[i] = bm.NewBitMap()
	}
	return lluvia
}

/*
Función que utiliza una variable 'l' de tipo *Lluvias y recibe como parámetros una variable 'dia' de tipo uint8 y una variable 'mes' de tipo Mes {
	Si el retorno del llamado a la función 'validaFecha' pasándole como parámetros 'dia' y 'mes' es false {
		Se retorna
	}
	Se declara la variable 'err' que es el retorno del llamado a la función 'On' del campo 'registros[mes]' de 'l' pasándole 'día' - 1 como parámetro {
		Si 'err' no es nil {
			Se retorna
		}
	}
}
*/
func (l *Lluvias) Registrar(dia uint8, mes Mes) {
	if !validaFecha(dia, mes) {
		return
	}
	err := l.registros[mes].On(dia - 1)
	if err != nil {
		return
	}
}

/*
Función que utiliza una variable 'l' de tipo *Lluvias y recibe como parámetros una variable 'dia' de tipo uint8 y una variable 'mes' de tipo Mes {
	Si el retorno del llamado a la función 'validaFecha' pasándole los parámetros 'dia' y 'mes' es false {
		Se retorna
	}
	Se declara la variable 'err' que es el retorno del llamado a la función 'Off' del campo 'registros[mes]' de 'l' pasándole 'día' - 1 como parámetro
	Si 'err' no es nil {
		Se retorna
	}
}
*/
func (l *Lluvias) Desregistrar(dia uint8, mes Mes) {
	if !validaFecha(dia, mes) {
		return
	}
	err := l.registros[mes].Off(dia - 1)
	if err != nil {
		return
	}
}

/*
Función que utiliza una variable 'l' de tipo *Lluvias y recibe como parámetros una variable 'dia' de tipo uint8 y una variable 'mes' de tipo Mes y retorna un bool {
	Si el retorno del llamado a la función 'validaFecha' pasándole los parámetros 'dia' y 'mes' es false {
		Se retorna false
	}
	Se declaran las variables 'on' y 'err' que son los retornos del llamado a la función 'IsOn' del campo 'registros[mes]' de 'l' pasándole 'dia' - 1 como parámetro
	Si 'err' no es nil {
		Se retorna false
	}
	Se retorna 'on'
}
*/
func (l *Lluvias) Llovio(dia uint8, mes Mes) bool {
	if !validaFecha(dia, mes) {
		return false
	}
	on, err := l.registros[mes].IsOn(dia - 1)
	if err != nil {
		return false
	}
	return on
}

/*
Función que utiliza una variable 'l' de tipo *Lluvias, recibe como parámetro la variable 'mes' de tipo Mes y retorna un *ListSet[uint8] {
	Se declara la variable 'dias' de tipo *ListSet[uint8] inicializada con el retorno del llamado a la función 'NewListSet[uint8]()'
	Para cuando la variable 'i' de tipo uint8 inicializada en 1 sea menor o igual a 31 {
		Si al declarar la variable 'on' que es el retorno del llamado a la función 'IsOn' del campo 'registros[mes]' de 'l' pasándole 'i' - 1 es true {
			Se llama a la variable 'Add' de 'dias' pasándole 'i' como parámetro
		}
	}
	Se retorna 'dias'
}
*/
func (l *Lluvias) ListarMes(mes Mes) *s.ListSet[uint8] {
	dias := s.NewListSet[uint8]()
	for i := uint8(1); i <= 31; i++ {
		if on, _ := l.registros[mes].IsOn(i - 1); on {
			dias.Add(i)
		}
	}
	return dias
}

/*
Función que utiliza una variable 'l' de tipo *Lluvias, recibe como parámetros las variables 'mes1' y 'mes2' de tipo Mes y retorna un *ListSet[uint8] {
	Se declara la variable 'diasComunes' de tipo *ListSet[uint8] inicializada con el retorno del llamado a la función 'NewListSet[uint8]()'
	Para cuando la variable 'i' de tipo uint8 inicializada en 1 sea menor o igual a 31 {
		Se declaran las variables 'on1' de tipo bool y '_' de tipo error que son los retornos del llamado a la función 'IsOn' del campo 'registros[mes1]' de 'l' pasándole 'i' - 1 como parámetro
		Se declaran las variables 'on2' de tipo bool y '_' de tipo error que son los retornos del llamado a la función 'IsOn' del campo 'registros[mes2]' de 'l' pasándole 'i' - 1 como parámetro
		Si 'on1' y 'on2' son true {
			Se llama a la función 'Add' de 'diasComunes' pasándole 'i' como parámetro
		}
	}
	Se retorna 'diasComunes'
}
*/
func (l *Lluvias) ListarDiasEnAmbosMeses(mes1, mes2 Mes) *s.ListSet[uint8] {
	diasComunes := s.NewListSet[uint8]()
	for i := uint8(1); i <= 31; i++ {
		on1, _ := l.registros[mes1].IsOn(i - 1)
		on2, _ := l.registros[mes2].IsOn(i - 1)
		if on1 && on2 {
			diasComunes.Add(i)
		}
	}
	return diasComunes
}
