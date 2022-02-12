package main

import "fmt"

func main() {
	// Operadores aritmeticos (), *, /, %, +, -
	var a = 4 + (2 * 3)
	fmt.Println(a)
	// Operadores de asignacion: =, +=, -=, *=, /=, %=
	var b = 10
	b += 2 // <- Esto es lo mismo que poner b = b + 2
	fmt.Println(b)

	// Declaraciones post-incremento y post-decremento: ++ , --
	// En go no existen las declaraciones pre-incremento
	// Estos operadores no son una expresion sino una declaracion
	// Una expresion es una parte del codigo que generan un resultado
	// Una declaracion es una instruccion del lenguaje que realiza una accion
	// Una expresion c = c++ + 2 normal en otro lenguaje en Go causa un error
	var c = 3
	var c2 = 3
	c++
	c2--
	fmt.Println(c)
	fmt.Println(c2)

	// Operadores de comparacion: >, <, ==, !=, >=, <=
	// Comparan dos expresiones y retornan un booleano
	fmt.Println(4 > 5)
	fmt.Println(4 < 5)

	// Operadores lógicos &&, ||
	// Permiten agrupar compraciones y retornan un booleano
	var edad = 25
	var edad2 = 8
	fmt.Println("Adulto joven: ", edad >= 18 && edad <= 30)
	fmt.Println("niño o viejo: ", edad2 < 18 || edad2 > 60)

	// Unario: !
	// negacion (invierte) del resultado booleano, true lo pasa a false, false lo pasa a true
	fmt.Println(!(5 == 5))

}
