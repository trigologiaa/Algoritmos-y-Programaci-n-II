package estructurasinterfaces

//Demoré en entregarlo porque no me daban bien los resultados gracias a la poca tolerancia de los test en circulo, entonces si o si tenía que restarle a los resultados para llegar a la tolerancia pero no me parecía jaja

import (
	"fmt"
	"math"
)

type Figura interface {
	Area() float64
	Perimetro() float64
	ToString() string
}

type Punto struct {
	X float64
	Y float64
}

// NewPunto crea un nuevo punto con coordenadas x e y.
// Pre: Los parámetros x e y no deben ser nulos.
// Post: Devuelve un puntero a un Punto con las coordenadas x e y especificadas.
func NewPunto(x, y float64) *Punto {
	return &Punto{x, y}
}

// ToString devuelve la representación en cadena del punto.
// Pre: El punto no debe ser nulo.
// Post: Devuelve la representación en cadena del punto.
func (p *Punto) ToString() string {
	return fmt.Sprintf("(%v, %v)", p.X, p.Y)
}

type Circulo struct {
	Centro *Punto
	Radius float64
}

// NewCirculo crea un nuevo círculo con centro en pto y radio r.
// Pre: El parámetro pto no debe ser nulo.
// Post: Devuelve un puntero a un Circulo con el centro y radio especificados.
func NewCirculo(pto *Punto, r float64) *Circulo {
	return &Circulo{pto, r}
}

// Area devuelve el área del círculo con máximo 2 dígitos después de la coma.
// Pre: El círculo no debe ser nulo.
// Post: Devuelve el área del círculo.
func (c *Circulo) Area() float64 {
	area := math.Pi * c.Radius * c.Radius
	return math.Round(area*100) / 100
}

// Perimetro devuelve el perímetro del círculo con máximo 2 dígitos después de la coma.
// Pre: El círculo no debe ser nulo.
// Post: Devuelve el perímetro del círculo.
func (c *Circulo) Perimetro() float64 {
	perimetro := 2 * math.Pi * c.Radius
	return math.Round(perimetro*100) / 100
}

// ToString devuelve la representación en cadena del círculo.
// Pre: El círculo no debe ser nulo.
// Post: Devuelve la representación en cadena del círculo.
func (c *Circulo) ToString() string {
	return fmt.Sprintf("Círculo con centro en %s y radio %d", c.Centro.ToString(), int(c.Radius))
}