package estructurasinterfaces

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCirculoArea(t *testing.T) {
	centro := NewPunto(0, 0)
	c := NewCirculo(centro, 3.0)
	resultado := c.Area()
	assert.InDelta(t, 28.25, resultado, 0.05)
}

func TestCirculoPerimetro(t *testing.T) {
	centro := NewPunto(0, 0)
	c := NewCirculo(centro, 3.0)
	resultado := c.Perimetro()
	assert.InDelta(t, 18.84, resultado, 0.05)
}

func TestCirculoToString(t *testing.T) {
	centro := NewPunto(0, 0)
	c := NewCirculo(centro, 3.0)
	resultado := c.ToString()
	assert.Equal(t, "Círculo con centro en (0, 0) y radio 3", resultado)
}