package ejercicios
//16/05/2024
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDados(t *testing.T) {
	n, k, x := 3, 6, 7
	formas := Dados(n, k, x)
	assert.Equal(t, 15, formas)
}