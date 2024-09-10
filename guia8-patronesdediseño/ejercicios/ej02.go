package ejercicios

//Parte: Interfaz que define métodos a implementar.
type Parte interface {
	ObtenerPrecio() float64
	ObtenerDescripcion() string
}

//ParteSimple: Estructura que representa nombre y precio.
type ParteSimple struct {
	nombre string
	precio float64
}

//NewParteSimple: Función que crea parte simple con nombre y precio indicado.
func NewParteSimple(nombre string, precio float64) *ParteSimple {
	return &ParteSimple{nombre: nombre, precio: precio}
}

//ObtenerPrecio: Función que devuelve precio de la parte simple.
func(p *ParteSimple) ObtenerPrecio() float64 {
	return p.precio
}

//ObtenerDescripcion: Función que devuelve nombre de la parte simple.
func(p *ParteSimple) ObtenerDescripcion() string {
	return p.nombre
}

//Ensamble: Estructura que tiene conjunto de partes (partes simples o ensambles).
type Ensamble struct {
	partes []Parte
}

//NewEnsamble: Función que crea ensamble.
func NewEnsamble() *Ensamble {
	return &Ensamble{}
}

//AgregarParte: Función que agrega una parte al ensamble.
func(p *Ensamble) AgregarParte(parte Parte) {
	p.partes = append(p.partes, parte)
}

//ObtenerPrecio: Función que devuelve precio de ensamble.
func(p *Ensamble) ObtenerPrecio() float64 {
	var precio float64
	for _, parte := range p.partes {
		precio += parte.ObtenerPrecio()
	}
	return precio
}

//ObtenerDescripción: Función que devuelve nombres del ensamble.
func(p *Ensamble) ObtenerDescripcion() string {
	var nombres string
	for i, parte := range p.partes {
		if i > 0 {
			nombres += ", "
		}
		nombres += parte.ObtenerDescripcion()
	}
	return nombres
}