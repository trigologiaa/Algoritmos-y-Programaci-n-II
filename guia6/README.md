# Guía 6

## Diccionario

### Ejercicio 1

Escriba una función que reciba un texto en castellano y un diccionario, y
devuelva la oración traducida palabra por palabra según el diccionario. Si una
palabra no se encuentra en el diccionario, deberá sustituirla por "error" en la
cadena resultante.

> Para parsear el string pueden usar: `func Split(str, sep string) []string`

Archivo: [`ejercicios.go`](ejercicios.go)

### Ejercicio 2

Escriba una función que reciba una cadena de caracteres y calcule la frecuencia
de aparición de cada caracter.

Por ejemplo:

```go
Frecuencia("hola como estas...")
```

debe devolver

```text
{a: 2, c: 1, e: 1, h: 1, l: 1, m: 1, o: 3, s: 2, t: 1}
```

> Recuerden que al recorrer una cadena deben castear cada valor a string ya que
> GO representa los caracteres como enteros. No se deben incluir los blancos en
> el conjunto. Por lo tanto se recomienda usar
> [`stringUtils`](https://pkg.go.dev/github.com/agrison/go-commons-lang/stringUtils)
> para chequear si un caracter es un blanco (espacio, tabulador, salto de línea,
> etc.)

Archivo: [`ejercicios.go`](ejercicios.go)

### Ejercicio 3

Escriba una función que reciba 2 slices y devuelva una nueva lista que es el
resultado de la interseccion de los 2 slices anteriores, la complejidad del
metodo debe ser _O(n)_

Archivo: [`ejercicios.go`](ejercicios.go)

### Ejercicio 4

Tenemos que ayudar a los docentes a preparar la información solicitada por el
departamento de alumnos.

Cuando toman asistencia anotan los presentes por día de clase, pero ahora desde
alumnos se les pide que envíen la información de asistencia por alumno.

Por ejemplo, si la lista que se recibe es:

```json
{"Mie 10": ["Ana", "Pedro"], "Vie 12":, ["Ana", "Luz"], "Mie 17":, ["Luz", "Pedro"]}
```

el resultado debe ser:

```json
{ "Ana": ["Mie 10", "Vie 12"], "Luz": ["Vie 12", "Mie 17"], "Pedro": ["Mie 10", "Mie 17"] }
```

Archivo: [`ejercicios.go`](ejercicios.go)
