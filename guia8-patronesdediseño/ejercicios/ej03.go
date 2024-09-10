package ejercicios

import "strconv"

//Componente: Interfaz para expresiones matemáticas.
type Componente interface {
	String() 	string
	Resolver() 	int
}

//Operador: Estructura que define los componentes de izquierdauierda, derechaecha y el operador.
type Operador struct {
	operador 	string
	izquierda 	Componente
	derecha 	Componente
}

//NewOperador: Función que crea los componentes de izquierdauierda, derechaecha y el operador.
func NewOperador(operador string, izquierda, derecha Componente) *Operador {
	return &Operador{operador: operador, izquierda: izquierda, derecha: derecha}
}

//String: Función que devuelve representación en string de los componentes y operadores.
func(p *Operador) String() string {
	return p.izquierda.String() + p.operador + p.derecha.String()
}

//Resolver: Función que resuelve los ejercicios dictados.
func(p *Operador) Resolver() int {
	switch p.operador {
	case "+":
		return p.izquierda.Resolver() + p.derecha.Resolver()
	case "-":
		return p.izquierda.Resolver() - p.derecha.Resolver()
	case "*":
		return p.izquierda.Resolver() * p.derecha.Resolver()
	case "/":
		return p.izquierda.Resolver() / p.derecha.Resolver()
	default:
		return 0
	}
}

//Operando: Estructura que representa un número en la expresión.
type Operando struct {
	valor int
}

//NewOperando: Función que crea un nuevo operando con un valor indicado.
func NewOperando(valor int) *Operando {
	return &Operando{valor: valor}
}

//String: Función que devuelve representación en string del operando.
func(p *Operando) String() string {
	return strconv.Itoa(p.valor)
}

//Resolver: Función que devuelve valor numérico del operando.
func(p *Operando) Resolver() int {
	return p.valor
}