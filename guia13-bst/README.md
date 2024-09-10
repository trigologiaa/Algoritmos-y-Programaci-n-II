# Guía Binary Search Tree - Árboles Binarios de Búsqueda
## Ejercicios

### Ejercicio 1
Escribir una función que encuentre el segundo elemento mas grande en un BST.
Example:

```bash
    Input: Root of below BST

        10
       /
    5

    Output:  5

    Input: Root of below BST

            10

          /   \

        5      20

                 \ 

                  30 

    Output:  20
```

### Ejercicio 2
Escribir una funcion que dado un BST, encuentre el predecesor inorder de una clave dada. Si la clave no se encuentra en el BST, devuelve el nodo más grande anterior (si lo hay) presente en la BST. 

```bash
          15
        /    \
       /      \
      10       20
     / \      /  \
    /   \    /    \
   8    12  16    25

    El predecesor de 8 no existe. 
    El predecesor de 10 es 8 
    El predecesor de 12 es 10 
    El predecesor de 20 es 16  
```

### Ejercicio 3
Escribir una funcion para comprobar si un árbol binario es BST o no. Devuelve un booleano

### Ejercicio 4
Implementar un `TreeSet[T Ordered]` a partir de un árbol binario de búsqueda

### Ejercicio 5
a. Dada la siguiente lista de números realizar lo que se pide en cada ítem devolviendo el dibujo del árbol resultante:
>	4    19    -7    49    100    0    22    12

b. Construir el árbol binario de búsqueda que resulta de insertar los números en el orden en que aparecen:
- Tomar el árbol anterior e insertar elemento 10.
- Tomar el árbol anterior y eliminar el elemento 49.
- Tomar el árbol anterior e insertar el elemento 1.

### Ejercicio 6
a. Dada la siguiente lista de números realizar lo que se pide en cada ítem devolviendo el dibujo del árbol resultante:

>	23    7    41    26    32    52    11    5    56

b. Construir el árbol binario de búsqueda que resulta de insertar los números en el orden en que aparecen:
- Tomar el árbol anterior y eliminar el elemento 41.
- Tomar el árbol anterior e insertar el elemento 9.
- Tomar el árbol anterior e insertar el elemento 28.
