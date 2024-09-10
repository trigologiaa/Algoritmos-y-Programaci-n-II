package arreglos

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSumaProductoEscalar123456(test *testing.T) {
	vector1 := []int {
		1,
		2,
		3,
	}
	vector2 := []int {
		4,
		5,
		6,
	}
	suma, producto := SumaProductoEscalar(vector1, vector2)
	assert.Equal(test, 21, suma)
	assert.Equal(test, 32, producto)
}

func TestSumaProductoEscalar987654(test *testing.T) {
	vector1 := []int {
		9,
		8,
		7,
	}
	vector2 := []int {
		6,
		5,
		4,
	}
	suma, producto := SumaProductoEscalar(vector1, vector2)
	assert.Equal(test, 39, suma)
	assert.Equal(test, 122, producto)
}