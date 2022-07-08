package main

import "fmt"

func main() {
	// para declarar una arregle en go
	// En Go, los arrays tienen tamaño fijo, no se puede modificar el tamaño después que se ha declarado.
	// Si se agregan mas posiciones el compilador lanzara un error (pero los "slices" permiten trabajar con arreglos de forma dinamica)
	// 1- Se declaran con la palabra reservada var
	// 2- entre corchetes se declaran las posisciones que tendra el array
	// 3- Se indica el tipo de dato que se almacenara
	// par asignar los valores, se usa el nombre de la variable seguido del indice entre corchetes donde se
	// almacenara el valor
	// var food [3]string
	// food[0] = "🍗"
	// food[1] = "🍖"
	// food[2] = "🥩"

	// Otra forma de trabjar con arrays es con los Arrays Literales
	// Es otra manera de declarar y asignar valores a un array.
	// se debe usar el operador de variable corta y despues del tipo de
	// dato que se almacenara se deben de abrir llaves y asignar los valores
	food := [3]string{"🥩", "🍗", "🍖"}

	fmt.Println(food)
}
