package ejercicios

import (
	"github.com/untref-ayp2/data-structures/queue"
	"github.com/untref-ayp2/data-structures/stack"
)

// Escribir una función que que reciba una cadena de caracteres y devuelva la cadena invertida. Analizar el orden.
/*
Función que recibe como parámetro una variable 'cadena' de tipo string y retorna un string {
	Se declara la variable 'aux' de tipo *stack.Stack[string] que se inicializa con la función 'New' de 'Stack'
	Se declara la variable 'invertida' de tipo string sin inicializar
	Para cada variable '_' e 'i' del rango de 'cadena' {
		Se llama a la función 'Push' de 'aux' pasándole el string 'i' como parámetro
	}
	Para cuando 'aux' no esté vacío {
		Se declaran las variables 'letras' y '_' que son los retornos de la función 'Pop' de 'aux'
		A la variable 'invertida' se le suma la variable 'letras'
	}
	Se retorna 'invertida'
}
*/
func InvertirCadena(cadena string) string {
	aux := stack.New[string]()
	var invertida string
	for _, i := range cadena {
		aux.Push(string(i))
	}
	for !aux.IsEmpty() {
		letras, _ := aux.Pop()
		invertida += letras
	}
	return invertida
}

// Escribir una función que verifique si una palabra es palíndromo (es decir que una cadena es igual a su inversa. Por ejemplo: las cadenas "1456541" y "145541" son palíndromos). Analizar el orden.
/*
Función que recibe como parámetro una variable 'cadena' de tipo string y retorna un bool {
	Se declara la variable 'aux' de tipo *stack.Stack[string] que se inicializa con la función 'New' de 'Stack'
	Se declara la variable 'invertida' de tipo string sin inicializar
	Para cada variable '_' e 'i' del rango de 'cadena' {
		Se llama a la función 'Push' de 'aux' pasándole el string 'i' como parámetro
	}
	Para cuando 'aux' no esté vacío {
		Se declaran las variables 'letras' y '_' que son los retornos de la función 'Pop' de 'aux'
		A la variable 'invertida' se le suma la variable 'letras'
	}
	Se retorna el bool resultante de la comparación == entre 'invertida' y 'cadena'
}
*/
func Palindromo(cadena string) bool {
	aux := stack.New[string]()
	var invertida string
	for _, i := range cadena {
		aux.Push(string(i))
	}
	for !aux.IsEmpty() {
		letras, _ := aux.Pop()
		invertida += letras
	}
	return invertida == cadena
}

// Escribir una función que evalúe si una cadena de paréntesis, corchetes y llaves está bien balanceada y devuelve `true` si está bien balanceada y `false` si está mal balanceada. Por ejemplo: `[()]{}{[()()]()}` debe devolver `true`, mientras que `[(])` debe devolver `false`. Analizar el orden.
/*
Función que recibe como parámetro una variable 'cadena' de tipo string y retorna un bool {
	Se declara la variable 'aux' de tipo *stack.Stack[string] que se inicializa con la función 'New' de 'Stack'
	Para cada variable '_' y 'runa' del rango de 'cadena' {
		Se declara la variable 'caracter' de tipo string que es el retorno de la función 'string' que recibe 'runa' como parámetro
		Para cada caso en 'caracter' {
			En el caso de que coincida con "(", "[" o "{":
				Se llama a la función 'Push' de 'aux' pasándole 'caracter' como parámetro
			En el caso de que coincida con ")":
				Si con los parámetros 'top' y 'error' que son los retornos de la función 'Pop' de 'aux', si 'error' no es nil o 'top' no es igual a "(" {
					Se retorna false
				}
			En el caso de que coincida con "]":
				Si con los parámetros 'top' y 'error' que son los retornos de la función 'Pop' de 'aux', si 'error' no es nil o 'top' no es igual a "[" {
					Se retorna false
				}
			En el caso de que coincida con "}":
				Si con los parámetros 'top' y 'error' que son los retornos de la función 'Pop' de 'aux', si 'error' no es nil o 'top' no es igual a "{" {
					Se retorna false
				}
		}
	}
	Se retorna el bool de la función 'IsEmpty' de 'aux'
}
*/
func Balanceada(cadena string) bool {
	aux := stack.New[string]()
	for _, runa := range cadena {
		caracter := string(runa)
		switch caracter {
		case "(", "[", "{":
			aux.Push(caracter)
		case ")":
			if top, error := aux.Pop(); error != nil || top != "(" {
				return false
			}
		case "]":
			if top, error := aux.Pop(); error != nil || top != "[" {
				return false
			}
		case "}":
			if top, error := aux.Pop(); error != nil || top != "{" {
				return false
			}
		}
	}
	return aux.IsEmpty()
}

// Escribir una función, tal que, dadas dos colas, construya una cola con el resultado de poner una a continuación de la otra. Por ejemplo: si `q1 = [1,2,3]` (con 1 al frente y 3 al final) y `q2 = [5,7]`, el resultado es `[1,2,3,5,7]` (con 1 al frente y 7 al final). Analizar el orden.
/*
Función que utiliza una variable T cualquiera y que recibe como parámetros las variables 'q1' y 'q2' de tipo *Queue[T] y que retorna un *Queue[T] {
	Se declara la variable 'final' de tipo *Queue[T] inicializada con el retorno de la función 'New[T]' de 'Queue'
	Para cuando 'q1' no esté vacío {
		Se declaran las variables 'elemento' y '_' inicializadas como los retornos de la función 'Dequeue' de 'q1'
		Se llama a la función 'Enqueue' de 'final' pasándole 'elemento' como parámetro
	}
	Para cuando 'q2' no esté vacío {
		Se declaran las variables 'elemento' y '_' inicializadas como los retornos de la función 'Dequeue' de 'q2'
		Se llama a la función 'Enqueue' de 'final' pasándole 'elemento' como parámetro
	}
	Se retorna 'final'
}
*/
func UnirColas[T any](q1, q2 *queue.Queue[T]) *queue.Queue[T] {
	final := queue.New[T]()
	for !q1.IsEmpty() {
		elemento, _ := q1.Dequeue()
		final.Enqueue(elemento)
	}
	for !q2.IsEmpty() {
		elemento, _ := q2.Dequeue()
		final.Enqueue(elemento)
	}
	return final
}