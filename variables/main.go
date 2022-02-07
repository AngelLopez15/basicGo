package main

import "fmt"

// usando variables
// En Go no es necesario declarar el tipo de de datos de la variable Go se lo asignara implicitamente
// pero este tipo de dato no podra ser mutado mas adelante en el codigo

func main() {
	// Declarando una variable y asignandole un valor despues
	var wolf string
	wolf = "🐺"

	// En Go cada variable que se declare debe de usarse de lo contrario marcara error

	// declarando varias variables en una sola linea
	var dog, cat = "🐶", "🐱"

	// Si se ve a declarar y asignar en la misma linea no es necesario usar la palabra "var"
	// pero si es necesario usar el shorthand ":="
	mono := "🐵"

	// para reasignar el valor si es del mismo tipo basta con poner el "="
	cat = "gatito"

	fmt.Println(wolf)
	fmt.Println(dog, cat)
	fmt.Println(mono)
}
