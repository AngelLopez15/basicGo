package main

import "fmt"

func main() {
	// Tipos de datos bool, string, numeric

	// tipo bool
	var tipoBooleano = false

	// tipo string
	// Para declarar una cadena de texto debemos de usar fuerza las comillas dobles ""
	// ya que las comillas simples se usan para otras cosas en GO
	var tipoString string = "Esta es una cadena"

	// tipo de datos numericos
	// podemos usar los tipo uint (unsigned integer) y init (integer)
	// uint8   - unsigned 8bit integers ( 0 a 255)
	// uint16  - unsigned 16bit integers ( 0 a 65,535)
	// uint32  - unsigned 32bit integers ( 0 a 4,294,967,295)
	// uint64  - unsigned 64bit integers ( 0 a 18,446,744,073,709,551,615)

	// int8   - 8bit integers ( -128 a 127)
	// int16  - 16bit integers ( -32,768 a 32767 )
	// int32  - 32bit integers ( -2,147,483,648 a 2,147,483,647)
	// int64  - 64bit integers ( -9,223,372,036,854,775,808 a 9,223,372,036,854,775,807)

	// Tenemos los tipo de datos "byte" y "rune" pero estos son tipos de datos aparte si no
	// que sin alias "byte" es alias de uint8 y "rune" es alias de int32 y tambien representa
	// a unicode code point

	var tipoNumericoEntero uint8 = 100

	// para los unicode es que se usan las comillas simples
	var unicode rune = 'a'

	// Datos tipo float para almacenar numeros decimales tenemos los float32 y float64
	var tiposFloat float32 = 25.64

	// Go no permite hacer operaciones entre tipos de datos ya que es un lenguaje fuetement tipado
	// para poder hacer eso podemos hacer uso del casting
	var a uint8 = 100
	var b uint16 = 23000
	// aunque convertimos el tipo de dato momentaneamente para hacer la operacion, el tipo de dato
	// original no se ve modificado
	c := uint16(a) + b

	// Identificador Blanco
	// Permite que exista variables que no se van a usar.
	// Nota: No se puede usar la variable corta (:=), porque no se está declarando (no tiene un nombre).
	var _ string = "Hola Mundo"

	// El metodo "Printf" acepta mas parametros para poder imprimir el valor y el tipo de dato
	// para esto hacemos uso de el signo % y de los verbos T = type y v=variable
	fmt.Printf("Tipo: %T, Valor: %v", tipoBooleano, tipoBooleano)
	fmt.Printf("Tipo: %T, Valor: %v", tipoString, tipoString)
	fmt.Printf("Tipo: %T, Valor: %v", tipoNumericoEntero, tipoNumericoEntero)
	fmt.Printf("Tipo: %T, Valor: %v", unicode, unicode)
	fmt.Printf("Tipo: %T, Valor: %v", tiposFloat, tiposFloat)
	fmt.Printf("Tipo: %T, Valor: %v", c, c)
}
