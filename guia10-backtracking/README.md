# Guía 10 - Backtracking

En la carpeta `/ejercicios` encontrarás los esqueletos de la implementación para
las siguientes consignas.

## Ejercicios

Usando los conceptos de backtracking resuelva los siguientes ejercicios.

0. En los ejercicios del [laberinto](https://github.com/untref-ayp2/examples/tree/main/backtracking/ratmaze)
   y [sudoku](https://github.com/untref-ayp2/examples/tree/main/backtracking/sudoku)
   investigar la estructura de datos que soporta la solución, e identificar las
   operaciones `EsFactible`, `EsSolución`, `Extender`, `Registrar` y `Borrar`.
   Tomar como ejemplo el código de [cambio de moneda](https://github.com/untref-ayp2/examples/tree/main/backtracking/cambio)
   y [N-reinas](https://github.com/untref-ayp2/examples/tree/main/backtracking/reinas)
   que está comentado.

1. **Problema de los dados**: Se tienen `n` dados de `k` caras cada uno, se
   desea saber la cantidad de formas de obtener una suma de `x` puntos al lanzar
   los `n` dados.

   > Ejemplo: Si tenemos `n=3` dados de `k=6` caras y queremos obtener el valor
   > `x=7` la solución es:
   >
    `[(1, 1, 5), (1, 2, 4), (1, 3, 3), (1, 4, 2), (1, 5, 1), (2, 1, 4), (2, 2, 3), (2, 3, 2), (2, 4, 1), (3, 1, 3), (3, 2, 2), (3, 3, 1), (4, 1, 2), (4, 2, 1), (5, 1, 1)]`
   >
   > En total 15 variantes distintas.

2. **Problema de la mochila 0-1 (Knapsack)**: Este es un problema [NP completos](https://es.wikipedia.org/wiki/NP-completo).
   Consiste en un problema de optimización combinatoria, donde se espera poder
   llenar una "mochila" con un peso limitado, por una cantidad de objetos, cada
   uno con un peso y valor específico, máximizando el valor total almacenado.
   Los objetos no se pueden fraccionar.

3. Modificar el problema anterior para que devuelva ademas la lista de objetos a
   incluir en la mochila.
