package main

import "fmt"

// Constantes en GO

func main() {
	// para declarar una constante se usa la palabra reservada "const"
	// En las constantes el valor que se le asigna no puede cambiar
	// No se puede usar el operador de variable corta ":=" por que el compilador la cambia a variable
	const pi = 3.1416

	fmt.Println(pi)
}
