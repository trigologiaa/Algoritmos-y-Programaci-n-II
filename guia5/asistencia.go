package guia5

import (
	bm	"github.com/untref-ayp2/data-structures/bitmap"
	s	"github.com/untref-ayp2/data-structures/set"
)

/*
Tipo de estructura {
	Se declara el campo 'clases' de tipo []BitMap
	Se declara el campo 'n' de tipo uint
	Se declara el campo 'm' de tipo uint8
}
*/
type Asistencias struct {
	clases	[]bm.BitMap
	n      	uint
	m      	uint8
}

/*
Función que recibe como parámetros la variable 'cantidadAlumnos' de tipo uint y la variable 'cantidadClases' de tipo uint8 y retorna un *Asistencias {
	Se declara la variable 'clases' de tipo []BitMap inicializada con el retorno del llamado a la función 'make' pasándole '[]BitMap' y 'cantidadClases' como parámetros
	Para cada posición 'i' del rango de 'clases' {
		Se cambia el valor de la posición 'i' de 'clases' por el retorno del llamado a la función 'NewBitMap'
	}
	Se retorna la dirección de 'Asistencias' {
		Se inicializa el campo 'clases' con la variable 'clases'
		Se inicializa el campo 'n' con la variable 'cantidadAlumnos'
		Se inicializa el campo 'm' con la variable 'cantidadClases'
	}
}
*/
func NewAsistencias(cantidadAlumnos uint, cantidadClases uint8) *Asistencias {
	clases := make([]bm.BitMap, cantidadClases)
	for i := range clases {
		clases[i] = *bm.NewBitMap()
	}
	return &Asistencias {
		clases: clases,
		n:      cantidadAlumnos,
		m:      cantidadClases,
	}
}

/*
Función que utiliza una variable 'a' de tipo *Asistencias y que recibe como parámetros la variable 'alumno' de tipo uint y la variable 'clase' de tipo uint8 {
	Si 'alumno' es mayor o igual al campo 'n' de 'a', o, 'clase' es mayor o igual al campo 'm' de 'a' {
		Se retorna
	}
	Se llama a la función 'On' del campo 'clases', con la posición 'clase', de 'a', pasándole como parámetro 'alumno' pasado a uint8
}
*/
func (a *Asistencias) RegistrarAsistencia(alumno uint, clase uint8) {
	if alumno >= a.n || clase >= a.m {
		return
	}
	a.clases[clase].On(uint8(alumno))
}

/*
Función que utiliza una variable 'a' de tipo *Asistencias, que recibe como parámetros la variable 'alumno' de tipo uint y la variable 'clase' de tipo uint8 y retorna un bool {
	Si 'alumno' es mayor o igual al campo 'n' de 'a', o, 'clase' es mayor o igual al campo 'm' de 'a' {
		Se retorna false
	}
	Se declaran las variables 'on' de tipo bool y '_' de tipo error inicializados como los retornos del llamado a la función 'IsOn' del campo 'clases', con la posición 'clase', de 'a', pasándole como parámetro 'alumno' pasado a uint8
	Se retorna 'on'
}
*/
func (a *Asistencias) Asistio(alumno uint, clase uint8) bool {
	if alumno >= a.n || clase >= a.m {
		return false
	}
	on, _ := a.clases[clase].IsOn(uint8(alumno))
	return on
}

/*
Función que utiliza una variable 'a' de tipo *Asistencias, que recibe como parámetro la variable 'clase' de tipo uint8 y retorna un Set[uint] {
	Se declara la variable 'resultado' de tipo *ListSet[uint] inicializada con el retorno del llamado a la función 'NewListSet[uint]()'
	Si 'clase' es mayor o igual al campo 'm' de 'a' {
		Se retorna 'resultado'
	}
	Para cuando la variable 'i' de tipo uint inicializada en 0 sea menor al campo 'n' de 'a' {
		Si al declarar las variables 'asistio' de tipo bool y '_' de tipo error inicializadas con el retorno del llamado a la función 'IsOn' del campo 'clases', con la posición 'clase', de 'a', pasándole como parámetro 'i' pasado a uint8, 'asistio' sea true {
			Se llama a la función 'Add' de 'resultado' pasándole 'i' como parámetro
		}
	}
	Se retorna 'resultado'
}
*/
func (a *Asistencias) ListarClase(clase uint8) s.Set[uint] {
	resultado := s.NewListSet[uint]()
	if clase >= a.m {
		return resultado
	}
	for i := uint(0); i < a.n; i++ {
		if asistio, _ := a.clases[clase].IsOn(uint8(i)); asistio {
			resultado.Add(i)
		}
	}
	return resultado
}

/*
Función que utiliza una variable 'a' de tipo *Asistencias, que recibe como parámetro la variable 'alumno' de tipo uint y retorna un Set[uint8] {
	Se declara la variable 'resultado' de tipo *ListSet[uint] inicializada con el retorno del llamado a la función 'NewListSet[uint8]()'
	Si 'alumno' es mayor o igual al campo 'n' de 'a' {
		Se retorna 'resultado'
	}
	Para cuando la variable 'clase' de tipo uint8 inicializada en 0 sea menor al campo 'm' de 'a' {
		Si al declarar las variables 'asistio' de tipo bool y '_' de tipo error inicializadas con el retorno del llamado a la función 'IsOn' del campo 'clases', con la posición 'clase', de 'a', pasándole como parámetro 'alumno' pasado a uint8, 'asistio' sea true {
			Se llama a la función 'Add' de 'resultado' pasándole 'clase' como parámetro
		}
	}
	Se retorna 'resultado'
}
*/
func (a *Asistencias) ListarAlumno(alumno uint) s.Set[uint8] {
	resultado := s.NewListSet[uint8]()
	if alumno >= a.n {
		return resultado
	}
	for clase := uint8(0); clase < a.m; clase++ {
		if asistio, _ := a.clases[clase].IsOn(uint8(alumno)); asistio {
			resultado.Add(clase)
		}
	}
	return resultado
}

/*
Función que utiliza una variable 'a' de tipo *Asistencias y retorna un *Set[uint8] {
	Se declara la variable 'resultado' de tipo *ListSet[uint] inicializada con el retorno del llamado a la función 'NewListSet[uint8]()'
	
}
*/
func (a *Asistencias) ListarAsistencias() s.Set[uint8] {
	resultado := s.NewListSet[uint8]()
	for clase := uint8(0); clase < a.m; clase++ {
		allAttended := true
		for alumno := uint(0); alumno < a.n; alumno++ {
			if attended, _ := a.clases[clase].IsOn(uint8(alumno)); !attended {
				allAttended = false
				break
			}
		}
		if allAttended {
			resultado.Add(clase)
		}
	}
	return resultado
}

func (a *Asistencias) ListarAsistenciaPerfecta() s.Set[uint] {
	resultado := s.NewListSet[uint]()
	for alumno := uint(0); alumno < a.n; alumno++ {
		perfectAttendance := true
		for clase := uint8(0); clase < a.m; clase++ {
			if attended, _ := a.clases[clase].IsOn(uint8(alumno)); !attended {
				perfectAttendance = false
				break
			}
		}
		if perfectAttendance {
			resultado.Add(alumno)
		}
	}
	return resultado
}