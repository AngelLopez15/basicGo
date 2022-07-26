package main

import "fmt"

func main() {
	// al estructura de control if es similar a la de todos los lenguajes de programacion
	// con la diferencia de que en GO la condicion no necesita esta dentro de parentesis
	// if condition {
	// codigo
	// } else if condicion2 {
	// 	codigo
	// } else {
	// 	codigo
	// }
	emoji := "❤"
	// emoji	:= "💔"
	// emoji2	:= "👻"

	if emoji == "❤" {
		fmt.Println("Este es un kokoro feliz")
	} else if emoji == "💔" {
		fmt.Println("Este es un kokoro roto")
	} else {
		fmt.Printf("el emoji es este -> %s", emoji)
	}

	// En los if de GO podemos declarar una variable corta y usarla unicamente en el scope del if
	// por tanto podemos ocuparla para hacer comparaciones en todo el if, else if de esta validacion
	if emoji2 := "👻"; emoji2 == "💔" {
		fmt.Println("Este es un kokoro roto")
	} else if emoji2 == "👻" {
		fmt.Println("Este es un fantasmon")
	} else {
		fmt.Printf("el emoji es este -> %s", emoji)
	}
}
