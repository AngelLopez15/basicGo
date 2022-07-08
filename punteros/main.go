package main

import "fmt"

// los punteros en Go son variables que almacenan la direccion en memoria de un valor

func main() {
	fruit := "🐺"
	// Para crear un puntero que apunte a la direccion en memoria.
	// pera declararlo se usa el asterisco
	var puntero *string
	// El & es el operador de direccion para ver la direccion en memoria
	puntero = &fruit
	// podemos cambiar el valor de la variable puntero
	// con el operador de desreferenciacion accedemos a la variable fruta y
	// le asignamos el nuevo valor
	*puntero = "🥵"

	fmt.Printf("Tipo: %T, Valor: %s, Direccion: %v\n", fruit, fruit, &fruit)
	// si al imprimirlo punemos el asterisco el nombre de la variable podemos acceder al valor de la variable
	fmt.Printf("Tipo: %T, Valor: %v, Desreferenciacion: %s", puntero, puntero, *puntero)
}
