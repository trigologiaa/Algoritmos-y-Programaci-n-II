[![Open in Codespaces](https://classroom.github.com/assets/launch-codespace-7f7980b617ed060a017424585567c406b6ee15c891e84e1186181d67ecf80aa0.svg)](https://classroom.github.com/open-in-codespaces?assignment_repo_id=15011991)
# Guía Patrones de Diseño

En la carpeta `/ejercicios` encontrarás los esqueletos de la implementación para
las siguientes consignas.

## Patrones de Diseño

Usando los patrones que se vieron en clase (_Composite_, _Adapter_, _Iterator_)
resuelva los siguientes ejercicios.

1. Por un lado, tenemos a nuestra disposición una clase concreta (externa, que
   no podemos modificar) denominada `FileMp3` que tiene un método para
   reproducir archivos con la siguiente declaración:

   ```go
   func (f *FileMp3) PlayMp3() string
   ```

   Por el otro lado, necesitamos reproducir archivos de ese tipo cumpliendo con
   la siguiente interfaz `ArchivoDeMusica`:

   ```go
   type ArchivoDeMusica interface {
        Reproducir() string
   }
   ```

   ¿Qué patrón de los vistos en clase resuelve el problema? Utilizandolo,
   implementar la clase `ArchivoMp3`, capaz de cumplir con la funcionalidad
   solicitada. Incluya todo el código necesario para poder probar la
   funcionalidad solicitada.

   (La reprodución puede simularse devolviendo un mensaje de texto, tal como
   `“Reproduciendo archivo <extensión>”`)

2. Una fábrica de piezas para automotores compra cuatro productos básicos y con
   ellos construye nuevos productos. Un nuevo producto puede estar compuesto por
   básicos y también por otros productos ya construidos por la fábrica. Los
   cuatro productos básicos que compra son tornillos, tuercas, rulemanes y
   varillas de acero. Se necesita poder calcular el costo de producción de cada
   producto, donde, si se trata de un producto básico el precio es igual al
   precio de compra, y si se trata de un producto compuesto, el costo es la suma
   de todos los productos básicos que lo componen.

   Por ejemplo, supongamos que `producto1` está compuesto por un `productoA`, un
   `tornillo` y una `tuerca`. Entonces:

   ```text
   costo(producto1) = costo(productoA) + costo(tornillo) + costo(tuerca)
   ```

3. Se quiere construir un editor/resolvedor de expresiones matemáticas. Una
   expresión válida estará formada o bien por un número, o bien por la
   suma/resta/división/multiplicación de dos expresiones. Ejemplos de
   expresiones válidas son:

   ```text
   4
   3+8
   14 \* (3+5)
   ```

   Revisar las pruebas para mayor detalle.

4. Se desea recorrer un array, pero visitando solamente las posiciones pares.
   No debe alterarse el array original. La estructura base se encuentra en el
   archivo `utils.go`.

5. Proporcionar el código para recorrer sólamente los elementos pares de un
   array. No debe alterarse el array original. La estructura base se encuentra
   en el archivo `utils.go`.
