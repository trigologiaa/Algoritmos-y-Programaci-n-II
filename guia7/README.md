# Guía 7

## Ejercicios

### Recursividad

1. Escribir una función recursiva que tome un entero `n` y devuelva la suma de
   los primeros `n` números.

2. Escribir una función recursiva que tome un entero `n` y devuelva su
   factorial.

3. Escribir una función recursiva que devuelva la cantidad de unos en la
   representación binaria de un número `n`. Use el hecho de que es igual al
   número de unos en la representación binaria de `n/2`, mas 1 si es impar.

4. Escribir una función recursiva que tome una cadena y devuelva como resultado
   la cadena invertida.

5. Escribir una función recursiva que reciba una cadena y un caracter y devuelva
   la cantidad de veces que aparece el caracter en la cadena.

6. Escribir una función recursiva que calcule el máximo común divisor entre dos
   números enteros.

   Nota: Se puede usar el algoritmo de Euclides:

   `MCD(a, 0) = a`

   `MCD(a, b) = MCD(b, a mod b)`

7. Escribir una función recursiva que tome dos números enteros y calcule la
   multiplicación entre ellos, usando sólo sumas.

8. Escribir una función recursiva que devuelva el cociente y el resto de la
   división entera mediante restas sucesivas.

9. Escribir una función recursiva que tome un array de números enteros y
   devuelva la suma de todos sus elementos.

10. Escribir una función recursiva que dado un número entero decimal devuelva su
    equivalente en binario.

11. Escribir una función recursiva que reciba 2 enteros `n` y `b` y devuelva
    true si `n` es potencia de `b`.

    Por ejemplo: `esPotencia(8, 2)` devuelve `true`.

12. Escribir una función recursiva que calcule Fibonacci de `n`.

### Búsquedas y Ordenamientos Recursivos

13. Escribir una función recursiva para determinar si un elemento está en un
    arreglo utilizando búsqueda binaria. Calcule su eficiencia.

14. Programar la busqueda ternaria recursiva. Donde en lugar de dividir el
    arreglo en dos partes iguales, se divide en tres partes iguales. Calcular el
    orden.

15. Se tiene un arreglo de `len(n) >= 3` elementos en forma de pico, esto es:
    estrictamente creciente hasta una determinada posición `p`, y estrictamente
    decreciente a partir de ella (con `0 < p < n-1`). Por ejemplo, en el arreglo
    `[1, 2, 3, 1, 0, -2]` la posición del pico es `p=2`.

    Se pide implementar un algoritmo de división y conquista de orden _O(log n)_
    que encuentre la posición p del pico.
