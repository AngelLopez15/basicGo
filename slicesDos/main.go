package main

import "fmt"

func main() {
	// los slices en Go tiene tamaΓ±o y capacidad
	// len(): numero de elementos en el slice
	// cap(): numero de elementos del array donde apunta el slice, a partir del indice de
	// donde se creo el slice
	food := [5]string{"π₯©", "π", "π", "π", "π©"}
	beef := food[1:3]

	// con la funcion append podemos ir agregando elementos a un array que fue hecho a partir de un slice
	// y lo agrega en la ultima posicion del array. PERO este tambien modifica al arreglo original, es decir
	// tambien va a insertar el nuevo elemento en al arreglo de donde se seca el slice sustituyendo al anterior valor
	// que estaba en esa posicion.
	beef = append(beef, "π₯", "π€", "π£")

	// Cuando se ha llegado a la capacidad maxima del arreglo original al que estamos haciendo referencia
	// array a partir del slice -> Array[4]{"π", "π", "π", "π©"}
	// hasta aqui tenemos una capacidad de 4 y si vamos agregando elementos con append se sustituiran la pizza y la dona
	// PERO al llegar a ese maximo de capacidad y seguimos agregando elementos lo que va a suceder es que GO creara un nuevo
	// ARRAY con una capacidad al doble es decir ahora tendra 8 y si llena esa tendra 16 y asi sucesivamente.
	// al ser un ARRAT nuevo se pierde la referencia al array original y ya NO SE SUSTITUIRAN LOS ELEMENTOS
	// new Array[8]{"π", "π", "π₯", "π€", "π£"}

	// Se pueden crear array desde cero es decir que no referencien a otro array (array literales)
	// como estamos referenciando a ningun array lo que hace GO es primero crear ese array con los elementos que le estamos
	// pasando
	fruits := []string{"π", "π"}
	fruits = append(fruits, "π")

	// Usando la funcion make para crear slices donde podemos indicar explicitamente tamaΓ±o y capacidad
	// el primer parametro es el tipo de elementos que tendra el array, el segundo el tamaΓ±o y el tercero
	// es la capacidad
	animals := make([]string, 0, 3)
	animals = append(animals, "πΆ", "π±", "π­", "π°")

	// El valor cero de los slices es nil que en otro lenguajes de programacion se conoce como nulo (null)
	// es decir un arreglo que esta inicializado sin valores
	var body []string
	// Si se inicializa como un array literal entonces nil ya es falso
	body2 := []string{}

	fmt.Println("food", food)
	fmt.Println("beef", beef)
	fmt.Println("TAMAΓO:", len(beef))
	fmt.Println("CAPACIDAD:", cap(beef))
	fmt.Println("-----ARRAY LITERALES------")
	fmt.Println("fruits", fruits)
	fmt.Println("TAMAΓO:", len(fruits))
	fmt.Println("CAPACIDAD:", cap(fruits))
	fmt.Println("-----ARRAY CON MAKE------")
	fmt.Println("animals", animals)
	fmt.Println("TAMAΓO:", len(animals))
	fmt.Println("CAPACIDAD:", cap(animals))
	fmt.Println("-----ARRAY nil------")
	fmt.Println("body", body)
	fmt.Println("TAMAΓO:", len(body))
	fmt.Println("CAPACIDAD:", cap(body))
	fmt.Println(body == nil)
	fmt.Println("body2", body2)
	fmt.Println(body2 == nil)
}
